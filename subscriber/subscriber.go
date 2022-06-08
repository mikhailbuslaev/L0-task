package subscriber

import (
	"log"
	"fmt"
	"os"
	"os/signal"
	. "nats-subscriber/order"
	stan "github.com/nats-io/stan.go"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	clusterID string = "test-cluster"
	clientID  string = "subscriber"
	channelID string = "foo"
)

type Subscriber struct {
	Cache []Order
}

func connectToDB() (*sql.DB, error){

	connectionString := "host=localhost port=8888 user=postgres "+
	"password=postgres dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s Subscriber) Run() {
	// Connect to stan cluster
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println("Cannot connect to cluster")
	}

	// Cache define
	s.Cache = make([]Order, 0, 10000)
	
	// Database connect
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Subscriber cannot conect to database")
	}

	defer db.Close()
	// Parse and append correct messages to cache
	msgHandler := func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		order := Order{}
		err := order.Unmarshall(m.Data)
		if err == nil {
			s.Cache = append(s.Cache, order)
			// Push order to db
			//err = recordToDB(order)
		} else {
			fmt.Printf("Invalid message: %d\n", err)
		}
	}
	// Subscribe to channel
	sub, err := sc.Subscribe(channelID, msgHandler)
	if err != nil {
		fmt.Println("Cannot subscribe to channel")
	}

	fmt.Printf("Connected to clusterID: [%s] clientID: [%s]\n", clusterID, clientID)

	// Unsubscribe if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<- signalChan
	fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
	sub.Unsubscribe()
	sc.Close()
}
