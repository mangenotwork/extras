package grpc

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/mangenotwork/extras/common/utils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
)

var (
	etcdAddrNilErr = fmt.Errorf("[ETCD参数错误] etcd Addr is null; ex: 192.168.0.1:2379,192.168.0.2:2379")
	notServiceErr = fmt.Errorf("[发现服务失败] 没有发现服务")
)

// NewClient return a new rpc client
// NewClient("192.168.0.1:1234")
func newClient(server string) (client *grpc.ClientConn, err error) {
	client, err = grpc.Dial(
		server,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(unaryInterceptorClient),
		grpc.WithStreamInterceptor(streamInterceptorClient))
	if err != nil {
		err = fmt.Errorf("[RPC] get client error: %v", err)
		return
	}
	return
}

// unaryInterceptor 中间件打印日志
func unaryInterceptorClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			break
		}
	}
	_, file, line, _ := runtime.Caller(3)
	md, _ := metadata.FromOutgoingContext(ctx)
	clientName := getValue(md, "clientname")
	serviceName := getValue(md, "servicename")
	requestId := getValue(md, "requestid")
	startTime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		fmt.Printf("[GRPC ERROR] %s->%s(%v) | id:%s | %s | %s:%d| err = %v",
			clientName,
			serviceName,
			cc.Target(),
			requestId,
			method,
			file, line,
			err)
	} else {
		log.Printf("[GRPC] %v | %s->%s(%v) | id:%s | %s |  %s:%d",
			time.Now().Sub(startTime),
			clientName,
			serviceName,
			cc.Target(),
			requestId,
			method,
			file, line)
	}

	return err
}

// wrappedStream
type wrappedStreamClient struct {
	grpc.ClientStream
}

// RecvMsg
func (w *wrappedStreamClient) RecvMsg(m interface{}) error {
	log.Println("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

// SendMsg
func (w *wrappedStreamClient) SendMsg(m interface{}) error {
	log.Println("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

// newWrappedStreamClient
func newWrappedStreamClient(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStreamClient{s}
}

// streamInterceptor is an example stream interceptor.
func streamInterceptorClient(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			break
		}
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStreamClient(s), nil
}

// setCtx
func setCtx(serviceName, myName, funcName string, grpcConn *grpc.ClientConn) context.Context {
	if grpcConn == nil {
		return nil
	}
	kv := []string {
		"RequestId", utils.IDStr(),
		"ClientName", myName,
		"ServiceName", serviceName,
	}
	return metadata.NewOutgoingContext(context.Background(), metadata.Pairs(kv...))
}


// discovery 发现服务
type discovery struct {
	serverAddr string
	etcdAddr []string
	clientName string
	serviceName string
	reqFuncName string
	isLog bool
	times int
	retry int // 重试次数
	retryTime time.Duration
}

// DiscoveryArg 创建发现服务对象参数
type ClientArg struct {
	ServiceAddr  string
	EtcdAddr    string
	ClientName  string
	ServiceName string
	ReqFuncName string
	OpenLog     bool
}

// NewDiscovery 创建发现服务
func NewClient(dis ClientArg) (*discovery, error) {
	etcdAddr := strings.Split(dis.EtcdAddr, ",")
	if len(etcdAddr) < 1 {
		return nil, etcdAddrNilErr
	}
	if len(dis.ServiceName) < 1 {
		return nil, fmt.Errorf("[Grpc] 参数err: ServiceName is null.")
	}

	return &discovery{
		serverAddr: dis.ServiceAddr,
		etcdAddr: etcdAddr,
		clientName: dis.ClientName,
		serviceName: dis.ServiceName,
		reqFuncName: dis.ReqFuncName,
		isLog: dis.OpenLog,
		times: 0,
		retry: 20,
		retryTime: 50*time.Millisecond,
	}, nil
}

// Conn
func (c *discovery) Conn() (client *grpc.ClientConn, ctx context.Context, err error) {
	client, err = newClient(c.serverAddr)
	ctx = setCtx(c.serviceName, c.clientName, c.reqFuncName, client)
	return
}


// Min  发现服务获取grpc连接; 负载均衡 - 最小连接数法;
func (c *discovery) Min() (client *grpc.ClientConn, ctx context.Context, err error) {
	// 避免一直重试
	if c.times > c.retry {
		err = notServiceErr
	}
	NewEtcdCli(c.etcdAddr)
	grpcIPKey,_ := etcdConn.GetMinKey(serverNameKey(c.serviceName))
	grpcIPList := strings.Split(grpcIPKey, "/")
	if len(grpcIPList) < 1 {
		time.Sleep(c.retryTime)  // 没有获取到服务地址, 可能是服务还在启动中, 等待50ms从新获取
		c.times++
		return c.Min()
	}
	grpcIP := grpcIPList[len(grpcIPList)-1]
	client, err = newClient(grpcIP)
	if err != nil || client == nil {
		time.Sleep(c.retryTime)  // 连不上可能是服务还在启动中, 等待50ms从新获取
		c.times++
		return c.Min()
	}

	// 使用GetMinKey方式需要执行GetMinKeyCallBack
	_=etcdConn.GetMinKeyCallBack(grpcIPKey)
	ctx = setCtx(c.serviceName, c.clientName, c.reqFuncName, client)
	return
}

// Rand  发现服务获取grpc连接; 负载均衡 - 随机法;
func (c *discovery) Rand() (client *grpc.ClientConn, ctx context.Context, err error) {
	if c.times > c.retry {
		err = notServiceErr
	}
	NewEtcdCli(c.etcdAddr)
	serviceNameKey := serverNameKey(c.serviceName)
	grpcIP, _ := etcdConn.GetRandKey(serviceNameKey)
	log.Println("grpcIP = ", grpcIP)
	client, err = newClient(grpcIP)
	if err != nil || client == nil {
		time.Sleep(c.retryTime) // 连不上可能是服务还在启动中, 等待50ms从新获取
		c.times++
		return c.Rand()
	}
	ctx = setCtx(c.serviceName, c.clientName, c.reqFuncName, client)
	return
}

// TODO  Discovery.Poll     发现服务获取grpc连接; 负载均衡 - 轮询法;

// TODO  Discovery.HashIP   发现服务获取grpc连接; 负载均衡 - 源地址哈希法;

// TODO  Discovery.Hash     发现服务获取grpc连接; 负载均衡 - 一致性哈希法;

// TODO  Discovery.WRR      发现服务获取grpc连接; 负载均衡 - 加权轮询法;

// TODO  Discovery.RandWH   发现服务获取grpc连接; 负载均衡 - 加权随机法;

// TODO  Discovery.Fastest  发现服务获取grpc连接; 负载均衡 - 最快响应速度法;