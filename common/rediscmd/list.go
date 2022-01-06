package rediscmd

import (
	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
)

// 获取List value
func LRANGEAll(key string) []interface{} {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LRANGE", key, 0, -1)
	res, err := redis.Values(rc.Do("LRANGE", key, 0, -1))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	return res
}

// LRANGE key start stop
// 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。
func LRANGE(key string, start, stop int64) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LRANGE", key, start, stop)
	res, err = redis.Values(rc.Do("LRANGE", key, start, stop))
	return
}

// 新创建list 将一个或多个值 value 插入到列表 key 的表头
func LPUSH(key string, values []interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, value := range values {
		args = args.Add(value)
	}
	logger.Info("执行redis : ", "LPUSH", args)
	_, err = rc.Do("LPUSH", args...)
	return
}

// RPUSH key value [value ...]
// 将一个或多个值 value 插入到列表 key 的表尾(最右边)。
// 如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表尾：比如对一个空列表 mylist 执行
// RPUSH mylist a b c ，得出的结果列表为 a b c ，等同于执行命令 RPUSH mylist a 、 RPUSH mylist b 、 RPUSH mylist c 。
// 新创建List  将一个或多个值 value 插入到列表 key 的表尾(最右边)。
func RPUSH(key string, values []interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, value := range values {
		args = args.Add(value)
	}
	logger.Info("执行redis : ", "RPUSH", args)
	_, err = rc.Do("RPUSH", args)
	return
}

// BLPOP key [key ...] timeout
// BLPOP 是列表的阻塞式(blocking)弹出原语。
// 它是 LPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BLPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
func BLPOP() {}

// BRPOP key [key ...] timeout
// BRPOP 是列表的阻塞式(blocking)弹出原语。
// 它是 RPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BRPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
func BRPOP() {}

// BRPOPLPUSH source destination timeout
// BRPOPLPUSH 是 RPOPLPUSH 的阻塞版本，当给定列表 source 不为空时， BRPOPLPUSH 的表现和 RPOPLPUSH 一样。
// 当列表 source 为空时， BRPOPLPUSH 命令将阻塞连接，直到等待超时，或有另一个客户端对 source 执行 LPUSH 或 RPUSH 命令为止。
func BRPOPLPUSH() {}

// LINDEX key index
// 返回列表 key 中，下标为 index 的元素。
func LINDEX(key string, index int64) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LINDEX", key, index)
	res, err = redis.String(rc.Do("LINDEX", key, index))
	return
}

// LINSERT key BEFORE|AFTER pivot value
// 将值 value 插入到列表 key 当中，位于值 pivot 之前或之后。
// 当 pivot 不存在于列表 key 时，不执行任何操作。
// 当 key 不存在时， key 被视为空列表，不执行任何操作。
// 如果 key 不是列表类型，返回一个错误。
// direction : 方向 bool true:BEFORE(前)    false: AFTER(后)
func LINSERT(direction bool, key, pivot, value string) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	directionStr := "AFTER"
	if direction {
		directionStr = "BEFORE"
	}
	logger.Info("执行redis : ", "LINSERT", key, directionStr, pivot, value)
	_, err = rc.Do("LINSERT", key, directionStr, pivot, value)
	return
}

// LLEN key
// 返回列表 key 的长度。
// 如果 key 不存在，则 key 被解释为一个空列表，返回 0 .
func LLEN(key string) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LLEN", key)
	res, err = redis.Int64(rc.Do("LLEN", key))
	return
}

// LPOP key
// 移除并返回列表 key 的头元素。
func LPOP(key string) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LPOP", key)
	res, err = redis.String(rc.Do("LPOP", key))
	return
}

// LPUSHX key value
// 将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。
// 和 LPUSH 命令相反，当 key 不存在时， LPUSHX 命令什么也不做。
func LPUSHX(key string, value interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LPUSHX", key, value)
	_, err = rc.Do("LPUSHX", key, value)
	return
}

// LREM key count value
// 根据参数 count 的值，移除列表中与参数 value 相等的元素。
// count 的值可以是以下几种：
// count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
// count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
// count = 0 : 移除表中所有与 value 相等的值。
func LREM(key string, count int64, value interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LREM", key, count, value)
	_, err = rc.Do("LREM", key, count, value)
	return
}

// LSET key index value
// 将列表 key 下标为 index 的元素的值设置为 value 。
// 当 index 参数超出范围，或对一个空列表( key 不存在)进行 LSET 时，返回一个错误。
func LSET(key string, index int64, value interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LSET", key, index, value)
	_, err = rc.Do("LSET", key, index, value)
	return
}

// LTRIM key start stop
// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
// 举个例子，执行命令 LTRIM list 0 2 ，表示只保留列表 list 的前三个元素，其余元素全部删除。
func LTRIM(key string, start, stop int64) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "LTRIM", key, start, stop)
	_, err = rc.Do("LTRIM", key, start, stop)
	return
}

// RPOP key
// 移除并返回列表 key 的尾元素。
func RPOP(key string) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "RPOP", key)
	res, err = redis.String(rc.Do("RPOP", key))
	return
}

// RPOPLPUSH source destination
// 命令 RPOPLPUSH 在一个原子时间内，执行以下两个动作：
// 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。
// 将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。
// 举个例子，你有两个列表 source 和 destination ， source 列表有元素 a, b, c ， destination
// 列表有元素 x, y, z ，执行 RPOPLPUSH source destination 之后， source 列表包含元素 a, b ，
// destination 列表包含元素 c, x, y, z ，并且元素 c 会被返回给客户端。
// 如果 source 不存在，值 nil 被返回，并且不执行其他动作。
// 如果 source 和 destination 相同，则列表中的表尾元素被移动到表头，并返回该元素，可以把这种特殊情况视作列表的旋转(rotation)操作。
func RPOPLPUSH(key, destination string) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "RPOPLPUSH", key, destination)
	res, err = redis.String(rc.Do("RPOPLPUSH", key, destination))
	return
}

// RPUSHX key value
// 将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表。
func RPUSHX(key string, value interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "RPUSHX", key, value)
	_, err = rc.Do("RPUSHX", key, value)
	return
}
