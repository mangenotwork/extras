package conn

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
)

/*
	github.com/go-sql-driver/mysql + database/sql
	应用场景是 执行固定业务sql用
 */

var mysqlDriverDBs map[string]*MysqlDriver

func InitMysqlDriver() {
	mysqlDriverDBs = make(map[string]*MysqlDriver, len(conf.Arg.Mysql))
	for _, v := range conf.Arg.Mysql {
		m, err := NewMysqlDriver(v.DBName, v.User, v.Password, v.Host)
		if err != nil {
			logger.Panic(err)
			continue
		}
		mysqlDriverDBs[v.DBName] = m
	}
}

type MysqlDriver struct {
	host string
	user string
	password string
	dataBase string
	maxOpenConn int
	maxIdleConn int
	DB *sql.DB
	log bool
	once *sync.Once
}

// CloseLog 关闭日志
func (m *MysqlDriver) CloseLog(){
	m.log = false
}

// SetMaxOpenConn
func (m *MysqlDriver) SetMaxOpenConn(number int) {
	m.maxOpenConn = number
}

// SetMaxIdleConn
func (m *MysqlDriver) SetMaxIdleConn(number int) {
	m.maxIdleConn = number
}

// Conn 连接mysql
func (m *MysqlDriver) Conn() (err error){
	m.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s",
		m.user, m.password, "tcp", m.host, m.dataBase))
	if err != nil {
		if m.log{
			log.Println("[Sql] Conn Fail : " + err.Error())
		}
		return err
	}
	m.DB.SetConnMaxLifetime(time.Hour)  //最大连接周期，超过时间的连接就close
	if m.maxOpenConn < 1{
		m.maxOpenConn = 10
	}
	if m.maxIdleConn < 1{
		m.maxIdleConn = 5
	}
	m.DB.SetMaxOpenConns(m.maxOpenConn)//设置最大连接数
	m.DB.SetMaxIdleConns(m.maxIdleConn) //设置闲置连接数
	return
}

func NewMysqlDriver(database, user, password, host string) (*MysqlDriver, error) {
	if len(host) < 1 {
		return nil, fmt.Errorf("host is null")
	}
	m := &MysqlDriver{
		host : host,
		user : user,
		password : password,
		dataBase : database,
		log: true,
		maxOpenConn: 10,
		maxIdleConn: 10,
		once: &sync.Once{},
	}
	return m, nil
}

func GetMysqlDriver(dbName string) *MysqlDriver {
	m, ok :=  mysqlDriverDBs[dbName]
	if !ok {
		logger.Panic("[DB] 未init")
	}
	err := m.Conn()
	if err != nil {
		logger.Error(err)
	}
	return m
}

