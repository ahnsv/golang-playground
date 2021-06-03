package peline

import (
	"log"
	"sync"
)

type TotalPoint struct {
	mu  sync.Mutex
	sum int
}

func (t *TotalPoint) addPoint(point int) {
	t.mu.Lock()
	t.sum += point
	t.mu.Unlock()
}

func SumStage(sumChannel <-chan Payload, safePoint *TotalPoint) {
	functionName := TraceFunction()
	log.Println("SumStage start")

	for {
		msg := <-sumChannel

		totalPoint := safePoint
		totalPoint.addPoint(msg.Point)

		log.Printf("%s: totalPoint is now [%v]", functionName, totalPoint.sum)
	}
}
