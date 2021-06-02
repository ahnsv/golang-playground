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

	userGroup := []string{"humphrey", "hardy", "thomas", "jeff", "knox", "flash", "grab"}

	go func() {
		for i := 1; i <= 1000000; i++ {
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

	go peline.RawStage(eventChannel, uniqueCountChannel)
	go peline.UniqueCountStage(uniqueCountChannel)

	time.Sleep(time.Second * 2)
}
