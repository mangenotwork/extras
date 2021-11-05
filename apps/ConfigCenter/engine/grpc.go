package engine

import "log"

func StartRpcServer(){
	go func() {
		log.Println("StartRpcServer...")
	}()
}