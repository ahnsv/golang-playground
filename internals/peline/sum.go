package peline

import "log"

func SumStage(sumChannel <-chan Payload) {
	functionName := TraceFunction()
	log.Println("SumStage start")

	totalPoint := 0

	for {
		msg := <-sumChannel
		totalPoint += msg.Point
		log.Printf("%s: totalPoint is now [%v]", functionName, totalPoint)
	}
}
