package subscriber

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	. "nats-subscriber/model"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

var (
	clusterID string  = "test-cluster"
	clientID  string  = "subscriber"
	channelID string  = "foo"
	Cache     []Order = make([]Order, 0, 10000)
)

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

func restoreCache(db *sql.DB) error {
	Cache = make([]Order, 0, 10000)
	rows, err := db.Query("select * from orders;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		o := Order{}
		var Chrt_id string
		err := rows.Scan(&o.Order_uid, &o.Track_number, &o.Entry, 
			&o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, 
			&o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, 
			&o.Delivery.Email, &o.Payment.Transaction, 
			&o.Payment.Request_id, &o.Payment.Currency, 
			&o.Payment.Provider, &o.Payment.Amount, 
			&o.Payment.Payment_dt, &o.Payment.Bank, 
			&o.Payment.Delivery_cost, &o.Payment.Goods_total, 
			&o.Payment.Custom_fee, &Chrt_id, &o.Locale, 
			&o.Internal_signature, &o.Customer_id, &o.Delivery_service, 
			&o.Shardkey, &o.Sm_id, &o.Date_created, &o.Oof_shard)
		if err != nil {
			return err
		}

			o.Items = make([]Item, 0, 5)

			items, err := db.Query("select * from items where items.chrt_id = '"+Chrt_id+"';")
			defer items.Close()
			if err != nil {
				return err
			}

			for items.Next() {
				item := Item{}
				err := items.Scan(&item.Chrt_id, &item.Track_number, 
					&item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, 
					&item.Total_price, &item.Nm_id, &item.Brand, &item.Status)
				if err != nil {
					return err
				}
				o.Items = append(o.Items, item)
			}
		Cache = append(Cache, o)
	}
	return nil
}

func pushToDB(o Order, db *sql.DB) error{
	_, err := db.Exec("call push_order('"+o.Order_uid+"','"+o.Track_number+"','"+o.Entry+"','"+
	o.Delivery.Name+"','"+o.Delivery.Phone+"','"+o.Delivery.Zip+"','"+
	o.Delivery.City+"','"+o.Delivery.Address+"','"+o.Delivery.Region+"','"+
	o.Delivery.Email+"','"+o.Payment.Transaction+"','"+
	o.Payment.Request_id+"','"+o.Payment.Currency+"','"+
	o.Payment.Provider+"',"+fmt.Sprintf("%d",o.Payment.Amount)+","+
	fmt.Sprintf("%d",o.Payment.Payment_dt)+",'"+o.Payment.Bank+"',"+
	fmt.Sprintf("%d",o.Payment.Delivery_cost)+","+fmt.Sprintf("%d", o.Payment.Goods_total)+","+
	fmt.Sprintf("%d",o.Payment.Custom_fee)+","+fmt.Sprintf("%d", o.Items[0].Chrt_id)+",'"+o.Locale+"','"+
	o.Internal_signature+"','"+o.Customer_id+"','"+o.Delivery_service+"','"+
	o.Shardkey+"',"+fmt.Sprintf("%d", o.Sm_id)+",'"+o.Date_created+"','"+o.Oof_shard+"');")
	if err != nil {
		return err
	}

	for i := range o.Items {
		_, err := db.Exec("call push_item("+fmt.Sprintf("%d", o.Items[i].Chrt_id)+",'"+
		o.Items[i].Track_number+"',"+fmt.Sprintf("%d", o.Items[i].Price)+",'"+
		o.Items[i].Rid+"','"+o.Items[i].Name+"',"+fmt.Sprintf("%d", o.Items[i].Sale)+",'"+
		o.Items[i].Size+"',"+fmt.Sprintf("%d", o.Items[i].Total_price)+","+
		fmt.Sprintf("%d", o.Items[i].Nm_id)+",'"+o.Items[i].Brand+"',"+fmt.Sprintf("%d", o.Items[i].Status)+");")
		if err != nil {
			return err
		}
	}
	return nil
}

func Run() {
	// Connect to stan cluster
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println("Cannot connect to cluster")
	}

	// Database connect
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Subscriber cannot conect to database")
	}
	defer db.Close()

	err = restoreCache(db)
	if err != nil {
		fmt.Printf("Cannot restore cache: %d \n", err)
	} else {
		fmt.Println("Succesfully restore cache")
	}

	// Parse and append correct messages to cache
	msgHandler := func(m *stan.Msg) {
		// Parsing message
		order := Order{}
		err := json.Unmarshal(m.Data, &order)
		fmt.Printf("Recieved and parsed message:"+order.Order_uid+"\n")
		if err == nil {
			// Append message to cache
			Cache = append(Cache, order)
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
	sub, err := sc.Subscribe(channelID, msgHandler)
	if err != nil {
		fmt.Println("Cannot subscribe to channel")
	}

	fmt.Printf("Connected to clusterID: [%s] clientID: [%s]\n", clusterID, clientID)

	// Unsubscribe if receiving Ctrl+C interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
	sub.Unsubscribe()
	sc.Close()
}
