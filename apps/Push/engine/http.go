package engine

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/net/netutil"

	"github.com/mangenotwork/extras/apps/Push/handler"
	"github.com/mangenotwork/extras/common/conf"
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
		log.Panic("Listen Err : %v", err)
		return
	}
	defer l.Close()

	// 开启最高连接数， 注意: linux/uinx有效果， win无效
	var rLimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("rLimit.Cur = ", rLimit.Cur)
	log.Println("rLimit.Max = ", rLimit.Max)
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Starting block word http server -> ", conf.Arg.HttpServer.Prod)
	// 对连接数的保护， 设置为最高连接数是 本机的最高连接数
	// https://github.com/golang/net/blob/master/netutil/listen.go
	l = netutil.LimitListener(l, int(rLimit.Max)*10)
	err = server.Serve(l)
	if err != nil {
		log.Panic("ListenAndServe err : ", err)
	}
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	mux.Handle("/ws", m(http.HandlerFunc(handler.Ws)))

	// 登记, 下发一个uuid以便确认设备
	mux.Handle("/register", m(http.HandlerFunc(handler.Register)))
	return mux
}