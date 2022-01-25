package engine

import (
	"github.com/mangenotwork/extras/apps/Push/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		mux.RouterFunc("/ws", handler.Ws)

		// [post] 登记, 下发一个随机uuid可以作为设备id,以便确认设备
		mux.Router("/register", handler.GetDeviceId)

		// [post] 创建 Topic
		mux.Router("/topic/create", handler.TopicCreate)

		// [post] 发布
		mux.Router("/topic/publish", handler.Publish)

		// [post] 设备订阅, 支持批量
		mux.Router("/topic/sub", handler.Subscription)

		// [post] 设备取消订阅, 支持批量
		mux.Router("/topic/cancel", handler.TopicCancel)

		// [get] 查询设备订阅的topic
		mux.Router("/device/view/topic", handler.DeviceViewTopic)

		// [get] 查询topic被哪些设备订阅
		mux.Router("/topic/all/device", handler.TopicAllDevice)

		// [get] 查询topic是否被指定device订阅
		mux.Router("/topic/check/device", handler.TopicCheckDevice)

		// [get] 强制指定topic下全部设备断开接收推送
		mux.Router("/topic/disconnection/all", handler.TopicDisconnection)

		// [get] 获取推送数据记录
		mux.Router("/topic/log", handler.TopicLog)


		mux.Run()

	}()
}
