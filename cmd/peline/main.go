package main

import (
	"net/http"
	"time"

	"github.com/ahnsv/golang-playground/internal/inbound"
	"github.com/ahnsv/golang-playground/internal/peline"
	"github.com/ahnsv/golang-playground/internal/picker"
)

func main() {
	chanBufSize := 3

	eventChannel := make(chan peline.Payload, chanBufSize)
	uniqueCountChannel := make(chan peline.Pair, chanBufSize)
	sumChannel := make(chan peline.Payload, chanBufSize)

	httpServer := inbound.InboundHttp{}

	userGroup := []string{"humphrey", "hardy", "thomas", "jeff", "knox", "flash", "grab"}
	httpServer.AddHandler(
		"/api/v1/send", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				go func() {
					currentTime := time.Now()
					currentTimeInInt64 := currentTime.UnixNano()

					eventChannel <- peline.Payload{
						Id:        picker.GenerateRandomNumber(1000, currentTimeInInt64),
						Username:  picker.PickOne(userGroup, currentTimeInInt64),
						Point:     picker.GenerateRandomNumber(10, currentTimeInInt64),
						Group:     "DE",
						CreatedAt: currentTime,
					}
				}()
			}
		})

	totalPoint := &peline.TotalPoint{}

	for i := 1; i <= 5; i++ {
		go peline.RawStage(eventChannel, uniqueCountChannel, sumChannel)
		go peline.UniqueCountStage(uniqueCountChannel)
		go peline.SumStage(sumChannel, totalPoint)
	}
	httpServer.Start()

}
