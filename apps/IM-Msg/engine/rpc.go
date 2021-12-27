package engine

import "log"

func StartRPC(){
	go func() {
		log.Println("StartRPC")
	}()
}