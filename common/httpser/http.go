package httpser

import (
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"golang.org/x/net/netutil"
)

func HttpServer(router *http.ServeMux){
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := &http.Server{
		Addr:         ":"+conf.Arg.HttpServer.Prod,
		ReadTimeout:  4*time.Second,
		WriteTimeout: 4*time.Second,
		IdleTimeout:  4*time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:      router,
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

	logger.Info("Starting http server port -> ", conf.Arg.HttpServer.Prod)
	// 对连接数的保护， 设置为最高连接数是 本机的最高连接数
	// https://github.com/golang/net/blob/master/netutil/listen.go
	l = netutil.LimitListener(l, int(rLimit.Max)*10)
	err = server.Serve(l)
	if err != nil {
		logger.Panic("ListenAndServe err : ", err)
	}
}

type Engine struct {
	mux  *http.ServeMux
	base func (next http.Handler) http.Handler
}

func NewEngine() *Engine {
	engine := &Engine{
		mux: http.NewServeMux(),
		base : Base,
	}
	engine.mux.Handle("/", engine.base(http.HandlerFunc(Index)))
	engine.mux.Handle("/hello", engine.base(http.HandlerFunc(Hello)))
	engine.mux.Handle("/health", engine.base(http.HandlerFunc(Health)))

	return engine
}

func SimpleEngine() *Engine {
	return &Engine{
		mux:  http.NewServeMux(),
		base: Base,
	}
}

func (engine *Engine) GetMux() *http.ServeMux {
	return engine.mux
}

func (engine *Engine) Router(path string, f func(w http.ResponseWriter, r *http.Request)) {
	engine.mux.Handle(path, engine.base(http.HandlerFunc(f)))
}

func (engine *Engine) RouterFunc(path string, f func(w http.ResponseWriter, r *http.Request)) {
	engine.mux.HandleFunc(path, f)
}

func (engine *Engine) Run() {
		runtime.GOMAXPROCS(runtime.NumCPU())
		server := &http.Server{
			Addr:         ":"+conf.Arg.HttpServer.Prod,
			ReadTimeout:  4*time.Second,
			WriteTimeout: 4*time.Second,
			IdleTimeout:  4*time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:      engine.mux,
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

		logger.Info("Starting http server port -> ", conf.Arg.HttpServer.Prod)
		// 对连接数的保护， 设置为最高连接数是 本机的最高连接数
		// https://github.com/golang/net/blob/master/netutil/listen.go
		l = netutil.LimitListener(l, int(rLimit.Max)*10)
		err = server.Serve(l)
		if err != nil {
			logger.Panic("ListenAndServe err : ", err)
		}
}

func (engine *Engine) OpenPprof() {
	engine.mux.Handle("/debug/pprof/", engine.base(http.HandlerFunc(pprof.Index)))
	engine.mux.Handle("/debug/pprof/cmdline", engine.base(http.HandlerFunc(pprof.Cmdline)))
	engine.mux.Handle("/debug/pprof/profile", engine.base(http.HandlerFunc(pprof.Profile)))
	engine.mux.Handle("/debug/pprof/symbol", engine.base(http.HandlerFunc(pprof.Symbol)))
	engine.mux.Handle("/debug/pprof/trace", engine.base(http.HandlerFunc(pprof.Trace)))
}


var Path, _ = os.Getwd()

func Index(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path,"/img"){
		log.Println("is img")
		file := Path + "/img" +r.URL.Path[len("/str"):]
		log.Println(file)
		f,err := os.Open(file)
		defer f.Close()
		if err != nil && os.IsNotExist(err){
			Out404(w)
			return
		}
		http.ServeFile(w,r,file)
		return
	}

	if strings.HasPrefix(r.URL.Path,"/static"){
		log.Println("is static")
		file := Path + "/static" +r.URL.Path[len("/str"):]
		log.Println(file)
		f,err := os.Open(file)
		defer f.Close()
		if err != nil && os.IsNotExist(err){
			Out404(w)
			return
		}
		http.ServeFile(w,r,file)
		return
	}

	if strings.HasPrefix(r.URL.Path,"/views"){
		log.Println("is views")
		file := Path + "/views" +r.URL.Path[len("/str"):]
		log.Println(file)
		f,err := os.Open(file)
		defer f.Close()
		if err != nil && os.IsNotExist(err){
			Out404(w)
			return
		}
		http.ServeFile(w,r,file)
		return
	}

	Hello(w,r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_,_= w.Write([]byte("Hello l'm mange.\n"+utils.Logo))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_,_= w.Write([]byte("true"))
}