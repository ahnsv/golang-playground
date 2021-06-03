package peline

import "log"

func UniqueCountStage(uniqueChannel <-chan Pair) {
	functionName := TraceFunction()
	log.Println("UniqueCountStage start")
	count := 0
	userSet := make(map[string]bool)

	for {
		msg := <-uniqueChannel
		if userSet[msg.Username] == true {
			continue
		}
		userSet[msg.Username] = true
		count += 1
		log.Printf("%s: Count is updated to [%d]", functionName, count)
	}
}
