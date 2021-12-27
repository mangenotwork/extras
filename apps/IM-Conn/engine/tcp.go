package engine

import "log"

func StartTCP(){
	go func() {
		log.Println("StartTCP")
	}()
}