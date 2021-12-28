package rediscmd

import (
	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
)

// 获取Hash value
func HGETALL(key string) map[string]string {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HGETALL", key)
	res, err := redis.StringMap(rc.Do("HGETALL", key))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	logger.Info(res)
	return res
}

// 新建Hash 单个field
// 如果 key 不存在，一个新的哈希表被创建并进行 HSET 操作。
// 如果域 field 已经存在于哈希表中，旧值将被覆盖。
func HSET(key, field string, value interface{}) string {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HSET", key, field, value)
	res, err := redis.String(rc.Do("HSET", key, field, value))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	logger.Info(res)
	return res
}


// HMSET key field value [field value ...]
// 新建Hash 多个field
// 同时将多个 field-value (域-值)对设置到哈希表 key 中。
// 此命令会覆盖哈希表中已存在的域。
func HMSET(key string, values []interface{}) error {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, value := range values {
		for k, v := range value.(map[string]interface{}) {
			args = args.Add(k)
			args = args.Add(v)
		}
	}
	logger.Info("执行redis : ", "HMSET", args)
	res, err := rc.Do("HMSET", args...)
	if err != nil {
		logger.Error("GET error", err.Error())
		return err
	}
	logger.Info(res)
	return nil
}

// HSETNX key field value
// 给hash追加field value
// 将哈希表 key 中的域 field 的值设置为 value ，当且仅当域 field 不存在。
func HSETNX(key, field string, value interface{}) error {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HSETNX", key, field, value)
	res, err := rc.Do("HSETNX", key, field, value)
	if err != nil {
		logger.Error("GET error", err.Error())
		return err
	}
	logger.Info(res)
	return nil
}

// HDEL key field [field ...]
// 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。
func HDEL(key string, fields []string) error {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, v := range fields {
		args = args.Add(v)
	}
	logger.Info("执行redis : ", "HDEL", args)
	res, err := rc.Do("HDEL", args)
	if err != nil {
		logger.Error("GET error", err.Error())
		return err
	}
	logger.Info(res)
	return nil
}

// HEXISTS key field
// 查看哈希表 key 中，给定域 field 是否存在。
func HEXISTS(key, fields string) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HEXISTS", key, fields)
	res, err := redis.Int(rc.Do("HEXISTS", key, fields))
	if err != nil {
		logger.Error("GET error", err.Error())
		return false
	}
	if res == 0 {
		return false
	}
	logger.Info(res)
	return true
}

// HGET key field
// 返回哈希表 key 中给定域 field 的值。
func HGET(key, fields string) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HGET", key, fields)
	res, err = redis.String(rc.Do("HGET", key, fields))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)
	return
}

// HINCRBY key field increment
// 为哈希表 key 中的域 field 的值加上增量 increment 。
// 增量也可以为负数，相当于对给定域进行减法操作。
// 如果 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令。
// 如果域 field 不存在，那么在执行命令前，域的值被初始化为 0
func HINCRBY(key, field string, increment int64) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HINCRBY", key, field, increment)
	res, err = redis.Int64(rc.Do("HINCRBY", key, field, increment))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)
	return
}

// HINCRBYFLOAT key field increment
// 为哈希表 key 中的域 field 加上浮点数增量 increment 。
// 如果哈希表中没有域 field ，那么 HINCRBYFLOAT 会先将域 field 的值设为 0 ，然后再执行加法操作。
// 如果键 key 不存在，那么 HINCRBYFLOAT 会先创建一个哈希表，再创建域 field ，最后再执行加法操作。
func HINCRBYFLOAT(key, field string, increment float64) (res float64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HINCRBYFLOAT", key, field, increment)
	res, err = redis.Float64(rc.Do("HINCRBYFLOAT", key, field, increment))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)
	return
}

// HKEYS key
// 返回哈希表 key 中的所有域。
func HKEYS(key string) (res []string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HKEYS", key)
	res, err = redis.Strings(rc.Do("HKEYS", key))
	if err != nil {
		return
	}
	return
}

// HLEN key
// 返回哈希表 key 中域的数量。
func HLEN(key string) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HLEN", key)
	res, err = redis.Int64(rc.Do("HLEN", key))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)
	return
}

// HMGET key field [field ...]
// 返回哈希表 key 中，一个或多个给定域的值。
// 如果给定的域不存在于哈希表，那么返回一个 nil 值。
func HMGET(key string, fields []string) (res []string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, v := range fields {
		args = args.Add(v)
	}
	logger.Info("执行redis : ", "HMGET", args)
	res, err = redis.Strings(rc.Do("HMGET", args))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)
	return
}

// HVALS key
// 返回哈希表 key 中所有域的值。
func HashHVALS(key string) (res []string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "HVALS", key)
	res, err = redis.Strings(rc.Do("HVALS", key))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	return
}

//HSCAN
//搜索value hscan test4 0 match *b*