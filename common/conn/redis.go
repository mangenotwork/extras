package conn

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conf"
)

var _redisConn *redis.Pool
var _redisOnce sync.Once

func RedisConn() *redis.Pool {
	_redisOnce.Do(func() {
		_redisConn = newRedisClient()
	})
	return _redisConn
}

func newRedisClient() *redis.Pool {
	redisConf := conf.Arg.Redis
	p := &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,
		MaxActive:   redisConf.MaxActive,
		IdleTimeout: 30 * time.Second,
		Wait:        true,
		Dial: func() (conn redis.Conn, err error) {
			return setDialog(redisConf)
		},
	}
	rc := p.Get()
	defer rc.Close()
	_, err := rc.Do("PING")
	if err != nil {
		panic(fmt.Sprintf("[RDS] redis 初始化失败 %v", err))
		return nil
	}
	return p
}

func setDialog(rdc *conf.Redis) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", rdc.Host, rdc.Port))
	if err != nil {
		return nil, err
	}

	if conn == nil {
		return nil, errors.New("连接redis错误")
	}

	if len(rdc.Password) != 0 {
		if _, err := conn.Do("AUTH", rdc.Password); err != nil {
			conn.Close()
		}
	}
	if _, err := conn.Do("SELECT", rdc.DB); err != nil {
		conn.Close()
	}

	return conn, nil
}
