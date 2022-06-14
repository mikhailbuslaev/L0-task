package publisher

import (
	"time"
	"math/rand"
	"fmt"
	"os"
	"os/signal"
	stan "github.com/nats-io/stan.go"
)

type Publisher struct {
	Name string
	Channel string
	Cluster string
	Timeout time.Duration
}

func New() *Publisher {
	p := &Publisher{}
	p.Name = "publisher"
	p.Channel = "foo"
	p.Cluster = "test-cluster"
	p.Timeout = 20
	return p
}

func randomizeId() int{
	return rand.Int() % 100000000
}

func (p *Publisher) Publish(msg []byte) {
	// Connecting to stan cluster
	sc, err := stan.Connect(p.Cluster, p.Name)
	if err != nil {
		fmt.Println("Cannot connect to invalid cluster")
		return
	}
	defer sc.Close()
	// Publish message
	sc.Publish(p.Channel, msg)
	fmt.Printf("Published message to channel: '%s' \n", p.Channel)
}

func (p *Publisher) Run() {
	var i int = 0
	go func(){
		for {
		p.Publish([]byte(`{
			"order_uid": "`+fmt.Sprintf("%d", randomizeId())+`",
			"track_number": "WBILMTESTTRACK",
			"entry": "WBIL",
			"delivery": {
			  "name": "Test Testov",
			  "phone": "+9720000000",
			  "zip": "2639809",
			  "city": "Kiryat Mozkin",
			  "address": "Ploshad Mira 15",
			  "region": "Kraiot",
			  "email": "test@gmail.com"
			},
			"payment": {
			  "transaction": "b563feb7b2b84b6test",
			  "request_id": "",
			  "currency": "USD",
			  "provider": "wbpay",
			  "amount": 1817,
			  "payment_dt": 1637907727,
			  "bank": "alpha",
			  "delivery_cost": 1500,
			  "goods_total": 317,
			  "custom_fee": 0
			},
			"items": [
			  {
				"chrt_id": 9934930,
				"track_number": "WBILMTESTTRACK",
				"price": 453,
				"rid": "ab4219087a764ae0btest",
				"name": "Mascaras",
				"sale": 30,
				"size": "0",
				"total_price": 317,
				"nm_id": 2389212,
				"brand": "Vivienne Sabo",
				"status": 202
			  }
			],
			"locale": "en",
			"internal_signature": "",
			"customer_id": "test",
			"delivery_service": "meest",
			"shardkey": "9",
			"sm_id": 99,
			"date_created": "2021-11-26T06:22:19Z",
			"oof_shard": "1"
		  }`))
		  i++
		  time.Sleep(p.Timeout*time.Second)
	}
}()
		// Unsubscribe if receiving Ctrl+C interrupt
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
	
		<-signalChan
		fmt.Println("Received an interrupt, end publishing...")
}
