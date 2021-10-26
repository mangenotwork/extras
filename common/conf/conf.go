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
}

type App struct {
	Name string `yaml:"name"`
	RunType string `yaml:"runType"`
}

type HttpServer struct {
	Prod string `yaml:"prod"`
}

type GrpcServer struct {
	Prod string `yaml:"prod"`
}

type GrpcClient struct {
	Prod string `yaml:"prod"`
}

type TcpServer struct {
	Prod string `yaml:"prod"`
}

type TcpClient struct {
	Prod string `yaml:"prod"`
}

type UdpServer struct {
	Prod string `yaml:"prod"`
}

type UdpClient struct {
	Prod string `yaml:"prod"`
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