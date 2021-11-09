package rediscmd

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
)

//获取ZSet value 返回集合 有序集成员的列表。
func ZRANGEALL(key string) []interface{} {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZRANGE", key, 0, -1, "WITHSCORES")
	res, err := redis.Values(rc.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil {
		log.Println("GET error", err.Error())
	}
	log.Println(res)
	return res
}

// ZRANGE key start stop [WITHSCORES]
// 返回有序集 key 中，指定区间内的成员。
// 其中成员的位置按 score 值递增(从小到大)来排序。
func ZRANGE(key string, start, stop int64) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZRANGE", key, start, stop, "WITHSCORES")
	res, err = redis.Values(rc.Do("ZRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		log.Println("GET error", err.Error())
	}
	log.Println(res)
	return
}

// ZREVRANGE key start stop [WITHSCORES]
// 返回有序集 key 中，指定区间内的成员。
// 其中成员的位置按 score 值递减(从大到小)来排列。
// 具有相同 score 值的成员按字典序的逆序(reverse lexicographical order)排列。
func ZREVRANGE(key string, start, stop int64) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREVRANGE", key, start, stop, "WITHSCORES")
	res, err = redis.Values(rc.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		log.Println("GET error", err.Error())
	}
	log.Println(res)
	return
}

//新创建ZSet 将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
func ZADD(key string, values []interface{}) error {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, value := range values {
		log.Println(value)
		for k, v := range value.(map[string]interface{}) {
			args = args.Add(v)
			args = args.Add(k)
		}
	}
	log.Println("执行redis : ", "ZADD", args)
	res, err := rc.Do("ZADD", args...)
	if err != nil {
		log.Println("GET error", err.Error())
		return err
	}
	log.Println(res)
	return nil
}

// ZCARD key
// 返回有序集 key 的基数。
func ZCARD(key string) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZCARD", key)
	res, err = redis.Int64(rc.Do("ZCARD", key))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZCOUNT key min max
// 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量。
func ZCOUNT(key string, min, max int64) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZCOUNT", key, min, max)
	res, err = redis.Int64(rc.Do("ZCOUNT", key, min, max))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZINCRBY key increment member
// 为有序集 key 的成员 member 的 score 值加上增量 increment 。
// 可以通过传递一个负数值 increment ，让 score 减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
// 当 key 不存在，或 member 不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
func ZINCRBY(key, member string, increment int64) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZINCRBY", key, increment, member)
	res, err = redis.String(rc.Do("ZINCRBY", key, increment, member))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
// 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
// 具有相同 score 值的成员按字典序(lexicographical order)来排列(该属性是有序集提供的，不需要额外的计算)。
// 可选的 LIMIT 参数指定返回结果的数量及区间(就像SQL中的 SELECT LIMIT offset, count )，注意当 offset 很大时，
// 定位 offset 的操作可能需要遍历整个有序集，此过程最坏复杂度为 O(N) 时间。
// 可选的 WITHSCORES 参数决定结果集是单单返回有序集的成员，还是将有序集成员及其 score 值一起返回。
// 区间及无限
// min 和 max 可以是 -inf 和 +inf ，这样一来，你就可以在不知道有序集的最低和最高 score 值的情况下，使用 ZRANGEBYSCORE 这类命令。
// 默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
func ZRANGEBYSCORE(key string, min, max, offset, count int64) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZRANGEBYSCORE", key, min, max, offset, count)
	res, err = redis.Values(rc.Do("ZRANGEBYSCORE", key, min, max, offset, count))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

func ZRANGEBYSCOREALL(key string) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZRANGEBYSCORE", key, "-inf", "+inf")
	res, err = redis.Values(rc.Do("ZRANGEBYSCORE", key, "-inf", "+inf"))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
// 返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员。有序集成员按 score 值递减(从大到小)的次序排列。
// 具有相同 score 值的成员按字典序的逆序(reverse lexicographical order )排列。
func ZREVRANGEBYSCORE(key string, min, max, offset, count int64) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREVRANGEBYSCORE", key, min, max, offset, count)
	res, err = redis.Values(rc.Do("ZREVRANGEBYSCORE", key, min, max, offset, count))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

func ZREVRANGEBYSCOREALL(key string) (res []interface{}, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREVRANGEBYSCORE", key, "-inf", "+inf")
	res, err = redis.Values(rc.Do("ZREVRANGEBYSCORE", key, "-inf", "+inf"))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZRANK key member
// 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从小到大)顺序排列。
// 排名以 0 为底，也就是说， score 值最小的成员排名为 0 。
func ZRANK(key string, member interface{}) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZRANK", key, member)
	res, err = redis.Int64(rc.Do("ZRANK", key, member))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZREM key member [member ...]
// 移除有序集 key 中的一个或多个成员，不存在的成员将被忽略。
func ZREM(key string, member []interface{}) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(key)
	for _, v := range member {
		args = args.Add(v)
	}
	log.Println("执行redis : ", "ZREM", args)
	res, err := rc.Do("ZREM", args...)
	if err != nil {
		log.Println("GET error", err.Error())
		return err
	}
	log.Println(res)
	return nil
}

// ZREMRANGEBYRANK key start stop
// 移除有序集 key 中，指定排名(rank)区间内的所有成员。
// 区间分别以下标参数 start 和 stop 指出，包含 start 和 stop 在内。
func ZREMRANGEBYRANK(key string, start, stop int64) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREMRANGEBYRANK", key, start, stop)
	res, err := redis.Int64(rc.Do("ZREMRANGEBYRANK", key, start, stop))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZREMRANGEBYSCORE key min max
// 移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
func ZREMRANGEBYSCORE(key string, min, max int64) (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREMRANGEBYSCORE", key, min, max)
	res, err := rc.Do("ZREMRANGEBYSCORE", key, min, max)
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZREVRANK key member
// 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递减(从大到小)排序。
// 排名以 0 为底，也就是说， score 值最大的成员排名为 0 。
// 使用 ZRANK 命令可以获得成员按 score 值递增(从小到大)排列的排名。
func ZREVRANK(key string, member interface{}) (res int64, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZREVRANK", key, member)
	res, err = redis.Int64(rc.Do("ZREVRANK", key, member))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZSCORE key member
// 返回有序集 key 中，成员 member 的 score
func ZSCORE(key string, member interface{}) (res string, err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	log.Println("执行redis : ", "ZSCORE", key, member)
	res, err = redis.String(rc.Do("ZSCORE", key, member))
	if err != nil {
		log.Println("GET error", err.Error())
		return
	}
	log.Println(res)
	return
}

// ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
// 计算给定的一个或多个有序集的并集，其中给定 key 的数量必须以 numkeys 参数指定，并将该并集(结果集)储存到 destination 。
func ZUNIONSTORE() {}

// ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
// 计算给定的一个或多个有序集的交集，其中给定 key 的数量必须以 numkeys 参数指定，并将该交集(结果集)储存到 destination 。
// 默认情况下，结果集中某个成员的 score 值是所有给定集下该成员 score 值之和.
func ZINTERSTORE() {}

//搜索值  ZSCAN key cursor [MATCH pattern] [COUNT count]
