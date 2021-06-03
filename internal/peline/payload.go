package peline

import (
	"sync"
	"time"
)

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

type PelineInSchema struct {
	Username *string `json:"username"`
	Point    *int    `json:"point"`
	Group    *string `json:"group"`
}

type PelineId struct {
	mu           sync.Mutex
	SerialNumber int
}

func (p *PelineId) Increment() {
	p.mu.Lock()
	p.SerialNumber += 1
	p.mu.Unlock()
}
