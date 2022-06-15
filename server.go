package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"nats-subscriber/cache"
	"nats-subscriber/publisher"
	"nats-subscriber/subscriber"

	"github.com/valyala/fasthttp"
)

func main() {
	c := cache.New()

	sub := subscriber.New()
	sub.ConnectCache(c)
	go sub.Run()

	pub := publisher.New()
	go pub.Run()

	getOrderHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Method()) != "POST" {
			ctx.Write([]byte(`{"message": "'/order' accept only POST method"}`))
			return
		}
		order, err := sub.Cache.Get(string(ctx.Request.Header.Peek("id")))
		if err == nil {
			ctx.Write([]byte(order.Data))
		} else {
			ctx.Write([]byte(`{"message": "server cannot find your order"}`))
		}
	}

	staticHandler := func(ctx *fasthttp.RequestCtx) {
		fasthttp.ServeFile(ctx, "static/index.html")
	}

	reqHandler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		switch path {
		case "/order":
			getOrderHandler(ctx)
		default:
			staticHandler(ctx)
		}
	}

	go func(){
		fmt.Println("Starting server...")
		if err := fasthttp.ListenAndServe(":1111", reqHandler); err != nil {
			log.Fatalf("error in ListenAndServe: %v", err)
		}
	}()

	// Stop serving if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
		
	<-signalChan
	fmt.Println("Received an interrupt, end serving...")
	time.Sleep(3*time.Second)
}
