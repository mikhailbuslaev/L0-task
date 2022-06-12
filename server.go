package main

import (
	"fmt"
	"log"
	"strings"
	"nats-subscriber/publisher"
	"nats-subscriber/subscriber"
	"nats-subscriber/cache"
	"github.com/valyala/fasthttp"
)

func main() {
	c := cache.New()

	sub := subscriber.New()
	sub.ConnectCache(c)
	go sub.Run()

	pub := publisher.New()
	go pub.Run()
	
	reqHandler := func(ctx *fasthttp.RequestCtx) {
		query := strings.Split(string(ctx.Path()), "=")
		if query[0] == "/order" {
			order, err := sub.Cache.Get(query[1])
			if err == nil {
				ctx.Write([]byte(order.Data))
			} else {
				ctx.Write([]byte(`{"message": "server cannot find your order"}`))
			}
		}
	}

	fmt.Println("Starting server...")
	if err := fasthttp.ListenAndServe(":1111", reqHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
