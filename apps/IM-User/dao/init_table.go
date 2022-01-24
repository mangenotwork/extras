package dao

import "github.com/mangenotwork/extras/apps/IM-User/global"

/*
	初始化table, 创建table
 */

var userBaseTable = `CREATE TABLE ` + global.UserBaseTableName + ` (
	id int(11) NOT NULL AUTO_INCREMENT,
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


