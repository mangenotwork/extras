package engine

import "log"

func StartUDP(){
	go func() {
		log.Println("StartUDP")
	}()
}