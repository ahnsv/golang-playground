package peline

import "fmt"

func RawStage(eventChannel <-chan Payload) {
	fmt.Println("Raw Stage Start")
	for {
		msg := <-eventChannel
		fmt.Printf("%v\n", msg)
	}
}
