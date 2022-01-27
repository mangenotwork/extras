package dao

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-Msg/global"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
)

/*
	初始化table, 创建table
*/


var userBaseTable = `CREATE TABLE ` + global.MsgTableName + ` (
	id int(11) NOT NULL AUTO_INCREMENT,
	send_type int(11),
	from_uid varchar(20),
	from_name varchar(20),
	from_avatar varchar(50),
	to_uid varchar(20),
	to_name varchar(20),
	to_avatar varchar(50),
	to_group_id int(11),
	to_group_name varchar(20),
	to_group_avatar varchar(50),
	content_type int(11),
	content text,
	send_time int(11),
	device varchar(20),
	source varchar(20),
	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`


func InitUserBaseTable(){
	db := conn.GetMysqlDriver("imtest")
	err := db.Conn()
	if err != nil {
		logger.Error(err)
		return
	}
	for i:=0; i<global.MaxMsgTable; i++ {
		_, err := db.DB.Exec(fmt.Sprintf(userBaseTable, i))
		if err != nil {
			logger.Error(err)
		}
	}
}