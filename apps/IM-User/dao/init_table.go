package dao

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-User/global"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
)

/*
	初始化table, 创建table
 */


var userBaseTable = `CREATE TABLE ` + global.UserBaseTableName + ` (
	id int(11) NOT NULL AUTO_INCREMENT,
	table_id int(11),
	uid varchar(20),
	uname varchar(20),
	account varchar(20),
	password varchar(20),
	sex int(11) DEFAULT 0,
	birthday int(11) DEFAULT 0,
	site varchar(20),
	mail varchar(30),
	phone varchar(11),
	avatar varchar(50),
	updated int(11),
	created int(11),
	deleted int(11),
	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`


func InitUserBaseTable(){
	db := conn.GetMysqlDriver("imtest")
	err := db.Conn()
	if err != nil {
		logger.Error(err)
		return
	}
	for i:=0; i<global.MaxUserBaseTable; i++ {
		_, err := db.DB.Exec(fmt.Sprintf(userBaseTable, i))
		if err != nil {
			logger.Error(err)
		}
	}
}

