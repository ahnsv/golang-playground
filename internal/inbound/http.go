package inbound

import (
	"log"
	"net/http"
)

type InboundHttp struct {
}

func (i *InboundHttp) Start() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (i *InboundHttp) AddHandler(route string, handlerFunc http.HandlerFunc) {
	http.Handle(route, handlerFunc)
}
