package engine

import "log"

func StartMQ(){
	go func() {
		log.Println("StartMQ")
	}()
}