package main

import (
	"fmt"
	"log"
	"nats-subscriber/subscriber"
	"nats-subscriber/cache"
	"github.com/valyala/fasthttp"
)

func main() {
	sub := subscriber.New()
	c := cache.New()
	sub.ConnectCache(c)

	go sub.Run()
	
	reqHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/order" {
			fmt.Println(sub.Cache)
		}
	}

	fmt.Println("Starting server")
	if err := fasthttp.ListenAndServe(":1111", reqHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
