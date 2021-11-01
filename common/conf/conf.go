package conf

import (
	"encoding/json"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	MqType string `yaml:"mqType"`
	Nsq *Nsq `yaml:"nsq"`
	Rabbit *Rabbit `yaml:"rabbit"`
	Kafka *Kafka `yaml:"kafka"`
}

type App struct {
	Name string `yaml:"name"`
	RunType string `yaml:"runType"`
}

type HttpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

type GrpcServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
	Log bool `yaml:"log"`
}

type GrpcClient struct {
	Prod string `yaml:"prod"`
}

type TcpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

type TcpClient struct {
	Prod string `yaml:"prod"`
}

type UdpServer struct {
	Open bool `yaml:"open"`
	Prod string `yaml:"prod"`
}

type UdpClient struct {
	Prod string `yaml:"prod"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB int `yaml:"db"`
	Password string `yaml:"password"`
	MaxIdle int `yaml:"maxidle"`
	MaxActive int `yaml:"maxactive"`
}

type MqType struct {

}

type Nsq struct {
	Producer string `yaml:"producer"`
	Consumer string `yaml:"consumer"`
}

type Rabbit struct {
	Addr string `yaml:"addr"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

type Kafka struct {
	Addr string `yaml:"addr"`
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