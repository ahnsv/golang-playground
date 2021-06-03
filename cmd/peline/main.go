package main

import (
	"time"

	"github.com/ahnsv/golang-playground/internals/peline"
	"github.com/ahnsv/golang-playground/internals/picker"
)

func main() {
	chanBufSize := 3

	eventChannel := make(chan peline.Payload, chanBufSize)
	uniqueCountChannel := make(chan peline.Pair, chanBufSize)
	sumChannel := make(chan peline.Payload, chanBufSize)

	userGroup := []string{"humphrey", "hardy", "thomas", "jeff", "knox", "flash", "grab"}

	go func() {
		for i := 1; i <= 100; i++ {
			currentTime := time.Now()
			currentTimeInInt64 := currentTime.UnixNano()

			eventChannel <- peline.Payload{
				Id:        i,
				Username:  picker.PickOne(userGroup, currentTimeInInt64),
				Point:     picker.GenerateRandomNumber(10, currentTimeInInt64),
				Group:     "DE",
				CreatedAt: currentTime,
			}
		}
	}()

	totalPoint := &peline.TotalPoint{}

	for i := 1; i <= 5; i++ {
		go peline.RawStage(eventChannel, uniqueCountChannel, sumChannel)
		go peline.UniqueCountStage(uniqueCountChannel)
		go peline.SumStage(sumChannel, totalPoint)
	}

	time.Sleep(time.Second * 2)
}
