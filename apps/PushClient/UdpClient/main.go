package main

import (
	"log"
	"net"
	"time"
)

func main() {
	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 1244}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	go func() {
		data := make([]byte, 1024)
		for{
			n, remoteAddr, err := conn.ReadFromUDP(data)
			if err != nil {
				log.Printf("error during read: %s", err)
			}
			log.Printf("<%s> %s\n", remoteAddr, data[:n])
		}
	}()

	go func() {
		time.Sleep(1*time.Second)
		_,_=conn.Write([]byte(`
{
	"cmd":"Auth",
	"data":{
		"device":"123"
	}
}
`))

		//fmt.Printf("<%s>\n", conn.RemoteAddr())
	}()

	//for {
	//	time.Sleep(1*time.Second)
	//	_,_=conn.Write([]byte(`xt`))
	//}

	select {}

}
