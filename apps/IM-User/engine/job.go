package engine

import "log"

func StartJob(){
	go func() {
		log.Println("StartJob")
	}()
}