package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahnsv/golang-playground/internal/inbound"
	"github.com/ahnsv/golang-playground/internal/peline"
)

func main() {
	chanBufSize := 3

	eventChannel := make(chan peline.Payload, chanBufSize)
	uniqueCountChannel := make(chan peline.Pair, chanBufSize)
	sumChannel := make(chan peline.Payload, chanBufSize)

	httpServer := inbound.InboundHttp{}

	// id := &peline.PelineId{}

	normalNumber := 1
	httpServer.AddHandler(
		"/api/v1/send", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {

				var inSchema *peline.PelineInSchema

				err := json.NewDecoder(r.Body).Decode(&inSchema)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				currentTime := time.Now()

				eventChannel <- peline.Payload{
					Id:        normalNumber,
					Username:  *inSchema.Username,
					Point:     *inSchema.Point,
					Group:     *inSchema.Group,
					CreatedAt: currentTime,
				}
				normalNumber++

				w.WriteHeader(http.StatusAccepted)
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
