package model

type ShortLink struct {
	Url string
	Aging int64 // 时效，单位秒
	Deadline int64 // 截止日期， 单位时间戳, 只有当Aging为0时才用
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

