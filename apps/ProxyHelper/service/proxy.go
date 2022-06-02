package service

import (
	"bytes"
	"fmt"
	"github.com/mangenotwork/extras/common/conf"
	"io"
	"log"
	"net"
	"net/url"
	"strings"

	//"github.com/mangenotwork/extras/common/conf"
	gt "github.com/mangenotwork/gathertool"
)

func Run(){
	//l, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Arg.HttpServer.Prod))
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//for {
	//	client, err := l.Accept()
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	go handleClientRequest(client)
	//}

	ipt := &gt.Intercept{
		Ip : "0.0.0.0:"+conf.Arg.HttpServer.Prod,
		HttpPackageFunc: func(pack *gt.HttpPackage){
			// 查看 ContentType
			log.Println("ContentType = ", pack.ContentType)
			//// 获取数据包数据为 http,json等 TXT格式的数据
			//log.Println("Txt = ", pack.Html())
			//// 获取数据包为图片，将图片转为 base64
			//log.Println("img base64 = ", pack.Img2Base64())
			//// 获取数据包为图片，存储图片
			//log.Println(pack.SaveImage(""))
		},
	}
	// 启动服务
	ipt.RunServer()
}

func handleClientRequest(client net.Conn){
	if client == nil {
		return
	}
	defer func() {
		_=client.Close()
	}()

	var b [1024*100]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	_,_=fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	log.Println("请求: ", host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		_,_ = fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		_,_ = server.Write(b[:n])
	}


	go func() {
		_,_ = io.Copy(server, client)
	}()

	_,_ = io.Copy(client, server)


}