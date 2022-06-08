package order

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"errors"
)

type Delivery struct {
	Name	 string `json:"name"`
	Phone	 string	`json:"phone"`
	Zip		 string	`json:"zip"`
	City 	 string `json:"city"`
	Address	 string `json:"address"`
	Region	 string `json:"region"`
	Email	 string `json:"email"`
}

type Payment struct {
	Transaction			string 	`json:"transaction"`
	Request_id	 		string	`json:"request_id"`
	Currency			string	`json:"currency"`
	Provider			string 	`json:"provider"`
	Amount				float64 `json:"amount"`
	Payment_dt			int 	`json:"payment_dt"`
	Bank		 		string 	`json:"bank"`
	Delivery_cost		float64 `json:"delivery_cost"`
	Goods_total		 	int 	`json:"goods_total"`
	Custom_fee		 	float64 `json:"custom_fee"`
}

type Item struct {
	Chrt_id			int 	`json:"chrt_id"`
	Track_number	string	`json:"track_number"`
	Price			float64 `json:"price"`
	Rid				string	`json:"rid"`
	Name			string	`json:"name"`
	Sale			float64	`json:"sale"`
	Size		 	string	`json:"size"`
	Total_price		float64	`json:"total_price"`
	Nm_id		 	int		`json:"nm_id"`
	Brand		 	string	`json:"brand"`
	Status		 	int   	`json:"brand"`
}

type Order struct{
	Order_uid 			string 		`json:"order_uid"`
	Track_number 		string 		`json:"track_number"`
	Entry 				string 		`json:"entry"`
	Delivery 			Delivery	`json:"delivery"`
	Payment 			Payment		`json:"payment"`
	Items 				[]Item 		`json:"items"`
	Locale 				string 		`json:"locale"`
	Internal_signature 	string 		`json:"internal_signature"`
	Customer_id 		string 		`json:"customer_id"`
	Delivery_service 	string 		`json:"delivery_service"`
	Shardkey 			string 		`json:"shardkey"`
	Sm_id 				int 		`json:"sm_id"`
	Date_created 		string 		`json:"date_created"`
	Oof_shard 			string 		`json:"oof_shard"`
}

func (o Order) Unmarshall(buf []byte) error {
	err := json.Unmarshal(buf, &o)
	if err != nil {
		return err
	}
	if o.Order_uid == "" {
		return errors.New("This is not order")
	}
	return nil
}

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o Order) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	switch data := v.(type) {
	case string:
		return json.Unmarshal([]byte(data), &o)
	case []byte:
		return json.Unmarshal(data, &o)
	default:
		return fmt.Errorf("cannot scan type %t into Map", v)
	}
}