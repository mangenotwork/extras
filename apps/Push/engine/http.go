package engine

import (
	"net"
	"net/http"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/net/netutil"

	"github.com/mangenotwork/extras/apps/Push/handler"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/middleware"

)

func StartHttpServer(){
	go func() {
		HttpServer()
	}()
}

func HttpServer(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := &http.Server{
		Addr:         ":"+conf.Arg.HttpServer.Prod,
		ReadTimeout:  4*time.Second,
		WriteTimeout: 4*time.Second,
		IdleTimeout:  4*time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:      Router(),
		// tls.Config 有个属性 Certificates []Certificate
		// Certificate 里有属性 Certificate PrivateKey 分别保存 certFile keyFile 证书的内容
	}

	// 如果在高频高并发的场景下, 有很多请求是可以复用的时候
	// 最好开启 keep-alives 减少三次握手 tcp 销毁连接时有个 timewait 时间
	server.SetKeepAlivesEnabled(true)
	l, err := net.Listen("tcp", server.Addr)
	if err != nil {
		logger.Panic("Listen Err : %v", err)
		return
	}
	defer l.Close()

	// 开启最高连接数， 注意: linux/uinx有效果， win无效
	var rLimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("rLimit.Cur = ", rLimit.Cur)
	logger.Info("rLimit.Max = ", rLimit.Max)
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("Starting block word http server -> ", conf.Arg.HttpServer.Prod)
	// 对连接数的保护， 设置为最高连接数是 本机的最高连接数
	// https://github.com/golang/net/blob/master/netutil/listen.go
	l = netutil.LimitListener(l, int(rLimit.Max)*10)
	err = server.Serve(l)
	if err != nil {
		logger.Panic("ListenAndServe err : ", err)
	}
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	mux.Handle("/ws", m(http.HandlerFunc(handler.Ws)))

	// [post] 登记, 下发一个随机uuid可以作为设备id,以便确认设备
	mux.Handle("/register", m(http.HandlerFunc(handler.GetDeviceId)))

	// [post] 创建 Topic
	mux.Handle("/topic/create", m(http.HandlerFunc(handler.TopicCreate)))

	// [post] 发布
	mux.Handle("/topic/publish", m(http.HandlerFunc(handler.Publish)))

	// [post] 设备订阅, 支持批量
	mux.Handle("/topic/sub", m(http.HandlerFunc(handler.Subscription)))

	// [post] 设备取消订阅, 支持批量
	mux.Handle("/topic/cancel", m(http.HandlerFunc(handler.TopicCancel)))

	// [get] 查询设备订阅的topic
	mux.Handle("/device/view/topic", m(http.HandlerFunc(handler.DeviceViewTopic)))

	// [get] 查询topic被哪些设备订阅
	mux.Handle("/topic/all/device", m(http.HandlerFunc(handler.TopicAllDevice)))

	// [get] 查询topic是否被指定device订阅
	mux.Handle("/topic/check/device", m(http.HandlerFunc(handler.TopicCheckDevice)))

	// [get] 强制指定topic下全部设备断开接收推送
	mux.Handle("/topic/disconnection/all", m(http.HandlerFunc(handler.TopicDisconnection)))

	// [get] 获取推送数据记录
	mux.Handle("/topic/log", m(http.HandlerFunc(handler.TopicLog)))


	return mux
}