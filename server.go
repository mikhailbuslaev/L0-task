package main

import (
	"log"
	"fmt"
	"nats-subscriber/subscriber"
	"github.com/valyala/fasthttp"
)

func main() {
	s := subscriber.Subscriber{}
	go s.Run()
	fmt.Println("Starting subscriber")
	reqHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/get_order" {
			fmt.Fprintf(ctx,"%v", s.Cache)
		}
	}
	fmt.Println("Starting server")
	if err := fasthttp.ListenAndServe(":1111", reqHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}