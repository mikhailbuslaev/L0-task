package subscriber

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"nats-subscriber/cache"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

var (
	Cache []Order = make([]Order, 0, 10000)
)

type Subscriber struct {
	Name string
	Channel string
	Cluster string
	Cache *cache.Cache
}

func New() *Subscriber {
	s := &Subscriber{}
	s.Name = "subscriber"
	s.Channel = "foo"
	s.Cluster = "test-cluster"
	return s
}

func (s *Subscriber) ConnectCache(c *cache.Cache) {
	s.Cache = c
}

func connectToDB() (*sql.DB, error) {

	connectionString := "host=localhost port=5432 user=postgres " +
		"password=postgres dbname=orders_test sslmode=disable"

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

func (s *Subscriber) restoreCache(db *sql.DB) error {
	s.Cache = cache.New()
	rows, err := db.Query("select * from orders;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		o := cache.Order{}
		err := rows.Scan(&o.Id, &o.Data)
		if err != nil {
			return err
		}
		s.Cache.Add(o)
	}
	return nil
}

func pushToDB(o Order, db *sql.DB) error{
	_, err := db.Exec("call push_order('"+o.Id+"', '"+o.Data+"');")
	if err != nil {
		return err
	}
	return nil
}

func (s *Subscriber) Run() {
	// Connect to stan cluster
	sc, err := stan.Connect(s.Cluster, s.Name)
	if err != nil {
		fmt.Println("Cannot connect to cluster")
	}

	// Database connect
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Subscriber cannot conect to database")
	}
	defer db.Close()

	s.restoreCache(db)
	if err != nil {
		fmt.Printf("Cannot restore cache: %d \n", err)
	} else {
		fmt.Println("Succesfully restore cache")
	}

	// Parse and append correct messages to cache
	msgHandler := func(m *stan.Msg) {
		// Parsing message
		order := cache.Order{}
		err := json.Unmarshal(m.Data, &order.Id)
		order.Data = m.Data
		fmt.Printf("Recieved and parsed message:"+order.Id+"\n")
		if err == nil {
			// Append message to cache
			s.Cache.Add(order)
			// Push order to db
			err = pushToDB(order, db)
			if err != nil {
				fmt.Printf("Cannot push to db: %d  \n", err)
			}
		} else {
			fmt.Printf("Invalid message: %d\n", err)
		}
	}
	// Subscribe to channel
	sub, err := sc.Subscribe(s.Channel, msgHandler)
	if err != nil {
		fmt.Println("Cannot subscribe to channel")
	}

	fmt.Printf("Connected to clusterID: [%s] clientID: [%s]\n", s.Cluster, s.Name)

	// Unsubscribe if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
	sub.Unsubscribe()
	sc.Close()
}
