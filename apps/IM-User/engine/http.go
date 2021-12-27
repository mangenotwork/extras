package engine

import "log"

func StartHTTP(){
	go func() {
		log.Println("StartHTTP")
	}()
}