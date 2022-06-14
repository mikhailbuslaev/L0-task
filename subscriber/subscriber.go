package subscriber

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"nats-subscriber/cache"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

type Subscriber struct {
	Name        string
	Channel     string
	Cluster     string
	RestoreFile string
	Cache       *cache.Cache
	DbConfig    DbConfig
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func New() *Subscriber {
	s := &Subscriber{}
	s.Name = "subscriber"
	s.Channel = "foo"
	s.Cluster = "test-cluster"
	s.RestoreFile = "restore.csv"
	s.DbConfig.Host = "localhost"
	s.DbConfig.Port = "5432"
	s.DbConfig.User = "postgres"
	s.DbConfig.Password = "postgres"
	s.DbConfig.DbName = "orders_test"
	return s
}

func (s *Subscriber) ConnectCache(c *cache.Cache) {
	s.Cache = c
}

func (s *Subscriber) getDBconnString() string {
	return "host=" + s.DbConfig.Host + " port=" + s.DbConfig.Port +
		" user=" + s.DbConfig.User + " password=" + s.DbConfig.Password + " dbname=" + s.DbConfig.DbName + " sslmode=disable"
}

func (s *Subscriber) checkDB() bool {
	db, err := sql.Open("postgres", s.getDBconnString())
	if err != nil {
		return false
	}
	err = db.Ping()
	return err == nil
}

func (s *Subscriber) connectToDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", s.getDBconnString())
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
	if !s.checkDB() {
		return errors.New("cannot restore cache: bad connection with database")
	}

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
		s.Cache.Add(&o)
	}
	return nil
}

func (s *Subscriber) pushToFile(o *cache.Order) error {
	file, err := os.OpenFile(s.RestoreFile, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	w.Comma = ';'
	err = w.Write([]string{o.Id, o.Data})
	if err != nil {
		return err
	}
	return nil
}

func (s *Subscriber) restoreFromFile() error {
	s.Cache = cache.New()
	file, err := os.OpenFile(s.RestoreFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		s.Cache.Add(&cache.Order{Id: record[0], Data: record[1]})
	}
	os.Truncate(s.RestoreFile, 0)
	return nil
}

func (s *Subscriber) pushToDB(o cache.Order, db *sql.DB) error {
	if !s.checkDB() {
		return errors.New("cannot push to db: bad connection with database")
	}
	_, err := db.Exec("call push_order('" + o.Id + "', '" + o.Data + "');")
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
	db, err := s.connectToDB()

	if err != nil {
		fmt.Println("Subscriber cannot connect to database")
	}
	defer db.Close()

	// Cache restore
	err = s.restoreCache(db)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully restore postgres cache")
	}
	// Restore csv cache
	err = s.restoreFromFile()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully restore csv cache")
	}

	// Parse and append correct messages to cache
	msgHandler := func(m *stan.Msg) {
		// Parsing message
		order := cache.Order{}
		err := json.Unmarshal(m.Data, &order)
		order.Data = string(m.Data)
		fmt.Printf("Recieved message:" + order.Id + "\n")
		if err == nil {
			// Append message to cache
			s.Cache.Add(&order)
			// Push order to db
			err = s.pushToDB(order, db)
			if err != nil {
				fmt.Printf("Cannot push to db: %d\n", err)
				// Temporary saving
				_ = s.pushToFile(&order)
			}
		} else {
			fmt.Printf("Invalid message: %d\n", err)
		}
	}
	// Subscribe to channel
	sub, err := sc.QueueSubscribe(s.Channel, "bar", msgHandler, stan.DurableName("dur"))
	if err != nil {
		fmt.Println("Cannot subscribe to channel")
	}

	fmt.Printf("Connected to clusterID: [%s] clientID: [%s]\n", s.Cluster, s.Name)

	// Unsubscribe if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Received an interrupt, unsubscribing and closing connection...")
	sub.Unsubscribe()
	sc.Close()
}
