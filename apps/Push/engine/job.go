package engine

import "log"

func StartJobServer(){
	go func() {
		log.Println("StartJobServer")
	}()
}