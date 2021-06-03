package peline

import "log"

func RawStage(eventChannel <-chan Payload, uniqueCountChannel chan<- Pair, sumChannel chan<- Payload) {
	log.Println("Raw Stage Start")
	functionName := TraceFunction()
	for {
		msg := <-eventChannel
		log.Printf("%s: Message Received - %v\n", functionName, msg)
		uniqueCountChannel <- Pair{Id: msg.Id, Username: msg.Username}
		sumChannel <- msg
	}
}
