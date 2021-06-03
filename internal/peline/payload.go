package peline

import "time"

type Payload struct {
	Id        int
	Username  string
	Point     int
	Group     string
	CreatedAt time.Time
}

type Pair struct {
	Id       int
	Username string
}
