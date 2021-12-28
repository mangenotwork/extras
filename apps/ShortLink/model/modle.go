package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

type ShortLink struct {
	Short string
	Url string
	Expiration int64 // 过期时间
	IsPrivacy bool // 是否隐私
	Password string // 只有当IsPrivacy=true使用
	Creation int64 // 创建时间
	View int64 // 请求次数
	OpenBlockList bool // 是否启用黑名单，启用后黑名单不能访问
	OpenWhiteList bool // 是否启用白名单，启用后只有白名单才能访问
	BlockList []string
	WhiteList []string
}

const (
	ShortLinkInfoKey = "link:%s:i" // 短链接基础数据，hash
	ShortLinkBlockListKey = "link:%s:b" // 短链接黑名单， set
	ShortLinkWhiteListKey = "link:%s:w" // 短链接白名单， set
)

func (sl *ShortLink) Save() (err error) {
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(fmt.Sprintf(ShortLinkInfoKey, sl.Short))
	// 不使用反射， 反射效率低
	args = args.Add("Short").Add(sl.Short)
	args = args.Add("Url").Add(sl.Url)
	args = args.Add("Expiration").Add(sl.Expiration)
	args = args.Add("IsPrivacy").Add(sl.IsPrivacy)
	args = args.Add("Password").Add(sl.Password)
	args = args.Add("Creation").Add(sl.Creation)
	args = args.Add("View").Add(sl.View)
	args = args.Add("OpenBlockList").Add(sl.OpenBlockList)
	args = args.Add("OpenWhiteList").Add(sl.OpenWhiteList)

	logger.Info("执行redis : ", "HMSET", args)
	res, err := rc.Do("HMSET", args...)
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	logger.Info(res)

	if sl.OpenBlockList {
		args := redis.Args{}.Add(fmt.Sprintf(ShortLinkBlockListKey, sl.Short))
		for _, value := range sl.BlockList {
			args = args.Add(value)
		}
		logger.Info("执行redis : ", "SADD", args)
		res, err := rc.Do("SADD", args...)
		if err != nil {
			logger.Error("GET error", err.Error())
		}
		logger.Info(res)
	}

	//whiteListKey := fmt.Sprintf(ShortLinkWhiteList, "/"+service.MustGenerate())
	if sl.OpenWhiteList {
		argsBlock := redis.Args{}.Add(fmt.Sprintf(ShortLinkWhiteListKey, sl.Short))
		for _, value := range sl.WhiteList {
			argsBlock = argsBlock.Add(value)
		}
		logger.Info("执行redis : ", "SADD", argsBlock)
		res, err := rc.Do("SADD", argsBlock...)
		if err != nil {
			logger.Error("GET error", err.Error())
		}
		logger.Info(res)
	}

	return
}

func (sl *ShortLink) Get(key string) error {
	rc := conn.RedisConn().Get()
	defer rc.Close()

	logger.Info("执行redis : ", "HGETALL", fmt.Sprintf(ShortLinkInfoKey, key))
	res, err := redis.StringMap(rc.Do("HGETALL", fmt.Sprintf(ShortLinkInfoKey, key)))
	if err != nil {
		logger.Error("GET error", err.Error())
		return err
	}
	logger.Info(res)
	sl.Short = key
	sl.Url = res["Url"]
	sl.Expiration = utils.Str2Int64(res["Expiration"])
	sl.IsPrivacy = utils.Str2Bool(res["IsPrivacy"])
	sl.Password = res["Password"]
	sl.Creation = utils.Str2Int64(res["Creation"])
	sl.View = utils.Str2Int64(res["View"])
	sl.OpenBlockList = utils.Str2Bool(res["OpenBlockList"])
	sl.OpenWhiteList = utils.Str2Bool(res["OpenWhiteList"])
	return nil
}

func (sl *ShortLink) GetUrl(key string) (url string,err error){
	err = sl.Get(key)
	url = sl.Url
	return
}

func (sl *ShortLink) IsWhiteList(ip string) (resBool bool) {
	resBool = true
	if !sl.OpenWhiteList {
		return
	}
	rc := conn.RedisConn().Get()
	defer rc.Close()
	res, err := redis.Int64(rc.Do("SISMEMBER", fmt.Sprintf(ShortLinkWhiteListKey, sl.Short), ip))
	if err != nil || res != 1 {
		resBool = false
		return
	}
	return
}

func (sl *ShortLink) IsBlockList(ip string) (resBool bool) {
	resBool = false
	if !sl.OpenBlockList {
		return
	}
	rc := conn.RedisConn().Get()
	defer rc.Close()
	res, err := redis.Int64(rc.Do("SISMEMBER", fmt.Sprintf(ShortLinkBlockListKey, sl.Short), ip))
	if err != nil {
		logger.Error("GET error", err.Error())
		return
	}
	if res == 1 {
		resBool = true
		return
	}
	return
}