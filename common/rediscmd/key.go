package rediscmd

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
)

// 获取所有的key
func GetALLKeys(matchValue string) (ksyList map[string]int) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	//初始化拆分值
	matchSplit := matchValue
	//matchvalue :匹配值，没有则匹配所有 *
	if matchValue == "" {
		matchValue = "*"
	} else {
		matchValue = fmt.Sprintf("*%s*", matchValue)
	}
	//cursor :初始游标为0
	cursor := "0"
	ksyList = make(map[string]int)
	ksyList, cursor = addGetKey(ksyList, matchValue, matchSplit, rc, cursor)
	//当游标等于0的时候停止获取key
	//线性获取，一直循环获取key,直到游标为0
	if cursor != "0" {
		for {
			ksyList, cursor = addGetKey(ksyList, matchValue, matchSplit, rc, cursor)
			if cursor == "0" {
				break
			}
		}
	}
	logger.Info("ksyList= ", ksyList)
	return
}

// addGetKey 内部方法
// 针对分组的key进行分组合并处理
func addGetKey(ksyList map[string]int, matchValue string, matchSplit string, conn redis.Conn, cursor string) (map[string]int, string) {
	//count_number :一次10000
	countNumber := "10000"
	logger.Info("执行redis : ", "scan", cursor, "MATCH", matchValue, "COUNT", countNumber)
	res, err := redis.Values(conn.Do("scan", cursor, "MATCH", matchValue, "COUNT", countNumber))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	//获取	matchvalue 含有多少:
	cfNumber := strings.Count(matchValue, ":")
	//获取新的游标
	newCursor := string(res[0].([]byte))
	allKey := res[1]
	allKeyData := allKey.([]interface{})
	for _, v := range allKeyData {
		keyData := string(v.([]byte))
		//没有:的key 则不集合
		if strings.Count(keyData, ":") == cfNumber || keyData == matchValue {
			ksyList[keyData] = 0
			continue
		}
		//有:需要集合
		keyDataNew, _ := fenGeYingHaoOne(keyData, matchSplit)
		ksyList[keyDataNew] = ksyList[keyDataNew] + 1
	}
	return ksyList, newCursor
}

// 对查询出来的key进行拆分，集合，分组处理
func fenGeYingHaoOne(str string, matchSplit string) (string, int) {
	likeKey := ""
	if matchSplit != "" {
		likeKey = fmt.Sprintf("%s", matchSplit)
	}
	str = strings.Replace(str, likeKey, "", 1)
	fg := strings.Split(str, ":")
	if len(fg) > 0 {
		return fmt.Sprintf("%s%s", likeKey, fg[0]), len(fg)
	}
	return "", len(fg)
}

func SearchKeys(matchValue string) (ksyList map[string]int) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	ksyList = make(map[string]int)
	//matchValue :匹配值，没有则返回空
	if matchValue == "" {
		return
	} else {
		matchValue = fmt.Sprintf("*%s*", matchValue)
	}
	//cursor :初始游标为0
	cursor := "0"
	ksyList = make(map[string]int)
	ksyList, cursor = addSearchKey(ksyList, matchValue, rc, cursor)
	//当游标等于0的时候停止获取key
	//线性获取，一直循环获取key,直到游标为0
	if cursor != "0" {
		for {
			ksyList, cursor = addSearchKey(ksyList, matchValue, rc, cursor)
			if cursor == "0" {
				break
			}
		}
	}
	logger.Info("ksyList= ", ksyList)
	return
}

// addGetKey 内部方法获取key
func addSearchKey(ksyList map[string]int, matchValue string, conn redis.Conn, cursor string) (map[string]int, string) {
	//count_number :一次10000
	countNumber := "10000"
	logger.Info("执行redis : ", "scan", cursor, "MATCH", matchValue, "COUNT", countNumber)
	res, err := redis.Values(conn.Do("scan", cursor, "MATCH", matchValue, "COUNT", countNumber))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	//获取新的游标
	newCursor := string(res[0].([]byte))
	allKey := res[1]
	allKeyData := allKey.([]interface{})
	for _, v := range allKeyData {
		keyData := string(v.([]byte))
		ksyList[keyData] = 0
	}
	return ksyList, newCursor
}

// 获取所有key name
// 返回切片
func GetAllKeyName() ([]interface{}, int) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	allKey := make([]interface{}, 0)
	keyData, cursor := getAllKey(rc, "0")
	allKey = append(allKey, keyData...)
	if cursor != "0" {
		for {
			keyData, cursor = getAllKey(rc, cursor)
			allKey = append(allKey, keyData...)
			if cursor == "0" {
				break
			}
		}
	}
	return allKey, len(allKey)
}

func getAllKey(conn redis.Conn, cursor string) ([]interface{}, string) {
	countNumber := "10000"
	logger.Info("执行redis : ", "scan", cursor, "MATCH", "*", "COUNT", countNumber)
	res, err := redis.Values(conn.Do("scan", cursor, "MATCH", "*", "COUNT", countNumber))
	if err != nil {
		logger.Error("GET error", err.Error())
	}
	return res[1].([]interface{}), string(res[0].([]byte))
}

// 获取key的信息
func GetKeyInfo(key string) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
}

// GetKeyType 获取key的类型
func GetKeyType(key string) string {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "TYPE", key)
	res, err := redis.String(rc.Do("TYPE", key))
	if err != nil {
		return ""
	}
	return res
}

// GetKeyTTL 获取key的过期时间
func GetKeyTTL(key string) int64 {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("执行redis : ", "TTL", key)
	res, err := redis.Int64(rc.Do("TTL", key))
	if err != nil {
		logger.Error("GET error", err.Error())
		return 0
	}
	return res
}

// EXISTSKey 检查给定 key 是否存在。
func EXISTS(key string) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("[Execute redis command]: ", "EXISTS", key)
	data, err := redis.String(rc.Do("DUMP", key))
	if err != nil || data == "0" {
		logger.Error("GET error", err.Error())
		return false
	}
	return true
}

// 修改key名称
func RenameKey(key, newName string) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("[Execute redis command]: ", "RENAME", key, newName)
	_, err := rc.Do("RENAME", key, newName)
	if err != nil {
		logger.Error("GET error", err.Error())
		return false
	}
	return true
}

// 更新key ttl
func UpdateKeyTTL(key string, ttlValue int64) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("[Execute redis command]: ", "EXPIRE", key, ttlValue)
	_, err := rc.Do("EXPIRE", key, ttlValue)
	if err != nil {
		logger.Error("GET error", err.Error())
		return false
	}
	return true
}

// 指定key多久过期 接收的是unix时间戳
func EXPIREATKey(key string, date int64) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("[Execute redis command]: ", "EXPIREAT", key, date)
	_, err := rc.Do("EXPIREAT", key, date)
	if err != nil {
		logger.Info("GET error", err.Error())
		return false
	}
	return true
}

// 删除key
func DELKey(key string) bool {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	logger.Info("[Execute redis command]: ", "DEL", key)
	_, err := rc.Do("DEL", key)
	if err != nil {
		logger.Error("GET error", err.Error())
		return false
	}
	return true
}

