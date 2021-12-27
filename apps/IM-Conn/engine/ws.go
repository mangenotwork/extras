package engine

import "log"

func StartWS(){
	go func() {
		log.Println("StartWS")
	}()
}