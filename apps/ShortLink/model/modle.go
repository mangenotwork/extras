package model

import (
	"github.com/mangenotwork/extras/apps/ShortLink/service"
	"github.com/mangenotwork/extras/common/conn"
	"log"
	"fmt"
	"github.com/garyburd/redigo/redis"
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
}

const (
	ShortLinkInfoKey = "link:%s:i" // 短链接基础数据，hash
	ShortLinkBlockList = "link:%s:b" // 短链接黑名单， set
	ShortLinkWhiteList = "link:%s:w" // 短链接白名单， set
)

func (sl *ShortLink) Save(){
	shortLinkKey := fmt.Sprintf(ShortLinkInfoKey, "/"+service.MustGenerate())
	rc := conn.RedisConn().Get()
	defer rc.Close()
	args := redis.Args{}.Add(shortLinkKey)
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

	log.Println("执行redis : ", "HMSET", args)
	res, err := rc.Do("HMSET", args...)
	if err != nil {
		log.Println("GET error", err.Error())
	}
	log.Println(res)


	//blockListKey := fmt.Sprintf(ShortLinkBlockList, "/"+service.MustGenerate())
	//whiteListKey := fmt.Sprintf(ShortLinkWhiteList, "/"+service.MustGenerate())
}