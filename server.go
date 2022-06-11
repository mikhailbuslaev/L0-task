package main

import (
	"fmt"
	"log"
	"nats-subscriber/subscriber"
	"github.com/valyala/fasthttp"
)

func main() {
	go subscriber.Run()
	reqHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/order" {
			fmt.Println(subscriber.Cache)
		}
	}

	fmt.Println("Starting server")
	if err := fasthttp.ListenAndServe(":1111", reqHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
