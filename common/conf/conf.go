package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

var Arg Configs

type Configs struct {
	App *App `yaml:"app"`
	HttpServer *HttpServer `yaml:"httpServer"`
	GrpcServer *GrpcServer `yaml:"grpcServer"`
	GrpcClient *GrpcClient `yaml:"grpcClient"`
	TcpServer *TcpServer `yaml:"tcpServer"`
	TcpClient *TcpClient `yaml:"tcpClient"`
	UdpServer *UdpServer `yaml:"udpServer"`
	UdpClient *UdpClient `yaml:"udpClient"`
	Redis *Redis `yaml:"redis"`
	Mysql []*Mysql `yaml:"mysql"`
	MqType string `yaml:"mqType"`
	Nsq *Nsq `yaml:"nsq"`
	Rabbit *Rabbit `yaml:"rabbit"`
	Kafka *Kafka `yaml:"kafka"`
	Mongo *Mongo `yaml:"mongo"`
	TTF string `yaml:"ttf"`
	Cluster *Cluster `yaml:"cluster"`
	LogCentre *LogCentre `yaml:"logCentre"`
	Jwt *Jwt `yaml:"jwt"`
	Minio *Minio `yaml:"minio"`
}

// App app相关基础信息
type App struct {
	Name string `yaml:"name"`
	RunType string `yaml:"runType"`
}

// HttpServer http服务
type HttpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

// GrpcServer grpc服务
type GrpcServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
	Log bool `yaml:"log"`
}

// GrpcClient grpc客户端
type GrpcClient struct {
	Prod string `yaml:"prod"`
}

// TcpServer tcp服务
type TcpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

// TcpClient tcp客户端
type TcpClient struct {
	Prod string `yaml:"prod"`
}

// UdpServer udp服务
type UdpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

// UdpClient udp客户端
type UdpClient struct {
	Prod string `yaml:"prod"`
}

// Redis redis配置
type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB int `yaml:"db"`
	Password string `yaml:"password"`
	MaxIdle int `yaml:"maxidle"`
	MaxActive int `yaml:"maxactive"`
}

// Mysql mysql配置
type Mysql struct {
	DBName string `yaml:"dbname"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
}

// MqType 消息队列类型
type MqType struct {

}

// Nsq 消息队列nsq配置
type Nsq struct {
	Producer string `yaml:"producer"`
	Consumer string `yaml:"consumer"`
}

// Rabbit 消息队列rabbit配置
type Rabbit struct {
	Addr string `yaml:"addr"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

// Kafka 消息队列kafka配置
type Kafka struct {
	Addr string `yaml:"addr"`
}

// Mongo mongo配置
type Mongo struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

// Cluster 集群使用 主要用于 ServiceTable
type Cluster struct {
	Open bool `yaml:"open"`
	MyAddr string `yaml:"myAddr"`
	InitCluster string `yaml:"initCluster"`
}

// LogCentre 日志服务收集日志配置
type LogCentre struct {
	Host string `yaml:"host"`
	Port int `yaml:"prod"`
}

// Jwt jwt配置
type Jwt struct {
	Secret string `yaml:"secret"`
	Expire int `yaml:"expire"`
}

// Minio 对象存储 minio配置
type Minio struct {
	Host string `yaml:"host"`
	Access string `yaml:"access"`
	Secret string `yaml:"secret"`
}

// 读取yaml文件
// 获取配置
func InitConf(){
	confFileName := "app.yaml"
	workPath, _ := os.Getwd()
	if os.Getenv("RUNMODE") != "" {
		confFileName = os.Getenv("RUNMODE") + ".yaml"
	}
	appConfigPath := filepath.Join(workPath, "configs", confFileName)
	if !fileExists(appConfigPath) {
		panic("【启动失败】 未找到配置文件!")
	}
	log.Println("[启动]读取配置文件:", appConfigPath)
	//读取yaml文件到缓存中
	config, err := ioutil.ReadFile(appConfigPath)
	if err != nil {
		panic("【启动失败】"+err.Error())
	}
	err =yaml.Unmarshal(config, &Arg)
	if err!=nil{
		panic("【启动失败】"+err.Error())
	}

	if Arg.Jwt == nil {
		Arg.Jwt = &Jwt{}
	}
	if Arg.Jwt.Secret == "" {
		Arg.Jwt.Secret = "mange-extras"
	}
	if Arg.Jwt.Expire == 0 {
		Arg.Jwt.Expire = 3600 // 默认一小时
	}

	b,_ := json.Marshal(Arg)
	log.Println("[conf arg] ", string(b))
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}