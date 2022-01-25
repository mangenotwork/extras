package httpser

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/mangenotwork/extras/common/logger"
	"golang.org/x/time/rate"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

func (lrw *ResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Base Http基础中间件,日志
func Base(next http.Handler) http.Handler {
	return BaseFunc(next)
}

// Base Http基础中间件,日志
func BaseFunc(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		/*
			// 中间件 上下文传递值
			data := map[string]interface{}{
			   "1": "one",
			   "2": "two",
			}
			ctx := context.WithValue(r.Context(), "data", data)
			r.WithContext(ctx)

			// 下文读值
			data := r.Context().Value("data").(ContextValue)["2"]
			fmt.Println(data) // 会打印 two
		*/

		start := time.Now().UnixNano()
		ip := GetIP(r)

		newW := NewResponseWriter(w)

		next.ServeHTTP(newW, r)

		logStr := fmt.Sprintf("%s#%s#%s#%d#%f", ip, r.Method, r.URL.String(), newW.StatusCode, float64(time.Now().UnixNano()-start)/100000)
		logger.Http(logStr, true)
	}
}

// ReqLimit 基础中间件 IP限流, IP黑白名单
func ReqLimit(ipv *IpVisitor, nextHeader http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := GetIP(r)
		if ipv.IsBlackList(ip) {
			_,_= w.Write([]byte("已经拉入黑名单，禁止访问！"))
			return
		}
		if !ipv.IsWhiteList(ip) {
			limiter := ipv.GetVisitor(ip)
			if limiter.AllowN(time.Now(), 1) == false {
				logger.Info("ip限流")
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return
			}
		}
		nextHeader.ServeHTTP(w, r)
		logger.Info("[%s] %s %s %v", ip, r.Method, r.URL.Path, time.Since(start))
	})
}


// GetIP 获取ip
// - X-Real-IP：只包含客户端机器的一个IP，如果为空，某些代理服务器（如Nginx）会填充此header。
// - X-Forwarded-For：一系列的IP地址列表，以,分隔，每个经过的代理服务器都会添加一个IP。
// - RemoteAddr：包含客户端的真实IP地址。 这是Web服务器从其接收连接并将响应发送到的实际物理IP地址。 但是，如果客户端通过代理连接，它将提供代理的IP地址。
//
// RemoteAddr是最可靠的，但是如果客户端位于代理之后或使用负载平衡器或反向代理服务器时，它将永远不会提供正确的IP地址，因此顺序是先是X-REAL-IP，
// 然后是X-FORWARDED-FOR，然后是 RemoteAddr。 请注意，恶意用户可以创建伪造的X-REAL-IP和X-FORWARDED-FOR标头。
func GetIP(r *http.Request) (ip string) {
	for _, ip := range strings.Split(r.Header.Get("X-Forward-For"), ",") {
		if net.ParseIP(ip) != nil {
			return ip
		}
	}
	if ip = r.Header.Get("X-Real-IP"); net.ParseIP(ip) != nil {
		return ip
	}
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if net.ParseIP(ip) != nil {
			return ip
		}
	}
	return "0.0.0.0"
}

type IpVisitor struct {
	ips map[string]*visitor
	mtx sync.Mutex
	BlackList map[string]struct{}
	WhiteList map[string]struct{}
}

func NewIpVisitor() *IpVisitor {
	return &IpVisitor{
		ips : make(map[string]*visitor),
	}
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// 启动一个协成  10分钟查一下ip限流数据，看看有没有超过1小时删除记录，有就删除
// 主要目的的为了释放内存空间
func (ipv *IpVisitor) CleanupVisitors() {
	go func() {
		timer1 := time.NewTicker(10 * time.Millisecond)
		select {
		case <-timer1.C:
			for ip, v := range ipv.ips {
				if time.Now().Sub(v.lastSeen) > 1*time.Hour {
					ipv.mtx.Lock()
					delete(ipv.ips, ip)
					ipv.mtx.Unlock()
				}
			}
		}
	}()
}

func (ipv *IpVisitor) GetVisitor(ip string) *rate.Limiter {
	ipv.mtx.Lock()
	defer  ipv.mtx.Unlock()
	v, exists := ipv.ips[ip]
	if !exists {
		return ipv.AddVisitor(ip)
	}
	// 更新时间
	v.lastSeen = time.Now()
	return v.limiter
}

func (ipv *IpVisitor) AddVisitor(ip string) *rate.Limiter {
	r := rate.Every(10 * time.Second) // 每*s往桶中放一个Token
	// 第一个参数是r Limit 代表每秒可以向Token桶中产生多少token
	// 第二个参数是b int b代表Token桶的容量大小
	limiter := rate.NewLimiter(r, 10)
	ipv.ips[ip] = &visitor{limiter, time.Now()}
	return limiter
}

func (ipv *IpVisitor) AddWhiteList(ip string) {
	ipv.WhiteList[ip] = struct{}{}
}

func (ipv *IpVisitor) IsWhiteList(ip string) (ok bool) {
	_,ok = ipv.WhiteList[ip]
	return
}

func (ipv *IpVisitor) DelWhiteList(ip string) {
	delete(ipv.WhiteList, ip)
}

func (ipv *IpVisitor) AddBlackList(ip string) {
	ipv.BlackList[ip] = struct{}{}
}

func (ipv *IpVisitor) IsBlackList(ip string) (ok bool) {
	_,ok = ipv.BlackList[ip]
	return
}

func (ipv *IpVisitor) DelBlackList(ip string) {
	delete(ipv.BlackList, ip)
}
