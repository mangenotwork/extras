package grpc

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/mangenotwork/extras/common/logger"
	"go.etcd.io/etcd/client/v3"
)

// 采用单列模式
var (
	etcdConn *EtcdCli
	once sync.Once
	randEr = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type EtcdCli struct {
	cli *clientv3.Client
	EtcdAddr []string
	ttl int // 申请租约的时间,单位秒, ttl秒后就会自动移除
	connTimeOut int // 连接etcd的timeout
}

func NewEtcdCli(etcdAddr []string) *EtcdCli {
	once.Do(
		func() {
			etcdConn = &EtcdCli{
				EtcdAddr: etcdAddr,
				ttl: 3,
				connTimeOut: 5,
			}
		},
	)
	return etcdConn
}

// SetTTl 设置ttl
func (etcdCil *EtcdCli) SetTTl(ttl int) *EtcdCli {
	etcdCil.ttl = ttl
	return etcdCil
}

// SetConnTimeOut 设置连接etcd的timeout
func (etcdCil *EtcdCli) SetConnTimeOut(timeOut int) *EtcdCli {
	etcdCil.connTimeOut = timeOut
	return etcdCil
}

// Conn 连接ETCD
func (etcdCil *EtcdCli) Conn() (*EtcdCli, error) {
	err := etcdCil.clientv3New()
	return etcdCil, err
}

func (etcdCil *EtcdCli) clientv3New() (err error) {
	if len(etcdCil.EtcdAddr) < 1 {
		err = fmt.Errorf("etcd addr is null")
		return
	}
	if etcdCil.cli == nil {
		etcdCil.cli, err = clientv3.New(clientv3.Config{
			Endpoints:   etcdCil.EtcdAddr,
			DialTimeout: time.Duration(etcdCil.connTimeOut) * time.Second,
		})
	}
	return
}

// Register 注册,并创建租约
func (etcdCil *EtcdCli) Register(key, value string) error {
	if err := etcdCil.clientv3New(); err != nil {
		return  err
	}
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			getResp, err := etcdCil.cli.Get(context.Background(), key)
			if err != nil {
				logger.Error("[ETCD] Register err : %s", err)
			} else if getResp.Count == 0 {
				err = etcdCil.withAlive(key, value)
				if err != nil {
					logger.Error("[ETCD] keep alive err :%s", err)
				}
			}
			<-ticker.C
		}
	}()
	return nil
}

// withAlive 创建租约
func (etcdCil *EtcdCli) withAlive(key, value string) error {
	leaseResp, err := etcdCil.cli.Grant(context.Background(), int64(etcdCil.ttl))
	if err != nil {
		return err
	}
	_, err = etcdCil.cli.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		logger.Error("[ETCD] put etcd error:%s", err)
		return err
	}

	ch, err := etcdCil.cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		logger.Error("[ETCD] keep alive error:%s", err)
		return err
	}

	// 清空 keep alive 返回的channel
	go func() {
		for {
			<-ch
		}
	}()

	return nil
}

// UnRegister 解除注册
func (etcdCil *EtcdCli) UnRegister(key string) error {
	if err := etcdCil.clientv3New(); err != nil {
		return  err
	}
	etcdCil.cli.Delete(context.Background(), key)
	return nil
}

// Get 获取
func (etcdCil *EtcdCli) Get(key string) (*clientv3.GetResponse, error) {
	if err := etcdCil.clientv3New(); err != nil {
		return  nil, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(etcdCil.connTimeOut)*time.Second)
	defer cancelFunc()
	return etcdCil.cli.Get(ctx, key, clientv3.WithPrefix())
}

// GetMinKey 轮询获取key, 前置条件: value是用于计数的
// key 是模糊key, 如 /A/
func (etcdCil *EtcdCli) GetMinKey(key string) (string, error) {

	response, err := etcdCil.Get(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if len(response.Kvs) < 1 {
		return "", fmt.Errorf("[ETCD] 在注册表没有找到服务")
	}

	// 采用轮询的思想, 每次拿最小被连数的那一个
	tmp := response.Kvs[0].Key
	tmpValue := response.Kvs[0].Value
	for i:=1; i<len(response.Kvs); i++ {
		if byte2int(tmpValue) > byte2int(response.Kvs[i].Value) {
			tmp = response.Kvs[i].Key
			tmpValue = response.Kvs[i].Value
		}
	}
	return string(tmp), nil
}

// GetMinKeyCallBack 与 EtcdCli.GetMinKey 配合使用客户端连接成功后调用
func (etcdCil *EtcdCli) GetMinKeyCallBack(key string) error {
	if err := etcdCil.clientv3New(); err != nil {
		return err
	}

	// 获取
	getResp, err := etcdCil.cli.Get(context.Background(), key)
	if err != nil {
		return err
	}

	if getResp.Count > 0 {
		v := getResp.Kvs[0].Value
		vInt := byte2int(v)+1
		_, err = etcdCil.cli.Put(context.Background(),key, fmt.Sprintf("%d",vInt), clientv3.WithPrevKV())
		if err != nil {
			logger.Error("[ETCD] put etcd error:%s", err)
			return err
		}
	}
	return nil
}

// GetAllKey 模糊key查询对应的所有key
func (etcdCil *EtcdCli) GetAllKey(key string) ([]string, error) {
	keys := make([]string, 0)

	response, err := etcdCil.Get(key)
	if err != nil {
		fmt.Println(err)
		return keys, err
	}

	if len(response.Kvs) < 1 {
		return keys, fmt.Errorf("[ETCD] 在注册表没有找到服务")
	}

	for _, ev := range response.Kvs {
		grpcIPList := strings.Split(string(ev.Key), "/")
		grpcIP := grpcIPList[len(grpcIPList)-1]
		keys = append(keys, grpcIP)
	}
	return keys, nil
}

//GetRandKey 模糊key查询随机反回一个key
func (etcdCil *EtcdCli) GetRandKey(key string) (string, error) {
	keys, err := etcdCil.GetAllKey(key)
	if err != nil {
		return "",err
	}
	l := len(keys)
	if l < 1 {
		return "", fmt.Errorf("[ETCD] 在注册表没有找到服务")
	}
	if l == 1 {
		return keys[0], nil
	}
	r := randEr.Intn(l)
	return keys[r], nil
}

// GetHash 通过传入h把value进行hash,然后后返回
// h : 用于计算hash的值
func (etcdCil *EtcdCli) GetHash(key, h string) {

}

// Delete
func (etcdCil *EtcdCli) Delete(key string) error {
	if err := etcdCil.clientv3New(); err != nil {
		return err
	}
	_, err := etcdCil.cli.Delete(context.TODO(), key, clientv3.WithPrevKV())
	return err
}

// Delete All
func (etcdCil *EtcdCli) DeleteAll(key string) error {
	if err := etcdCil.clientv3New(); err != nil {
		return err
	}
	_, err := etcdCil.cli.Delete(context.TODO(), key, clientv3.WithPrefix())
	return err
}

