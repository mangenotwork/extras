package engine

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		// Set
		// Command : SetAdd key value1,value2,
		// 集合添加数据
		mux.Router("/set/add", handler.SetAdd)
		// Command : SetAddExpire key value timeUnix
		// 集合添加数据并指定过期时间
		mux.Router("/set/addExpire", handler.SetAddExpire)
		// Command : SetValueExpire key value timeUnix
		// 指定集合数据过期时间
		mux.Router("/set/valueExpire", handler.SetValueExpire)
		// Command : SetGet key
		// 获取集合所有数据
		mux.Router("/set/get", handler.SetGet)
		// Command : SetDel key
		// 删除指定集合
		mux.Router("/set/del", handler.SetDel)
		// Command : SetDelValue key value
		// 删除指定集合的元素
		mux.Router("/set/value", handler.SetDelValue)

		// key 是否存在
		// Command : KeyHas key
		mux.Router("/key/has", handler.KeyHas)
		// key 模糊查询
		// Command : KeyLike key
		mux.Router("/key/like", handler.KeyLike)
		// key 列表
		// Command : KeyAll
		mux.Router("/key/all", handler.KeyAll)

		// KV
		// 增,改
		// Command : KVAdd key value
		mux.Router("/kv/add", handler.KVAdd)
		// Command : KVAddExpire key value expire
		mux.Router("/kv/addExpire", handler.KVAddExpire)
		// Command : KVExpire key expire
		mux.Router("/kv/expire", handler.KVExpire)
		// 删
		// Command : KVDel key
		mux.Router("/kv/del", handler.KVDel)
		// 查
		// Command : KVGet key
		mux.Router("/kv/get", handler.KVGet)

		mux.Run()

	}()
}
