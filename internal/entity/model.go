package entity

import (
	"encoding/json"
	"time"
)

func (order Order) String() string {
	str, _ := json.MarshalIndent(order, "", " ")
	return string(str)
}

type Order struct {
	OrderUID          string    `fake:"{regex:[a-z]{5}}"       json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `fakesize:"2"                  json:"items"`
	Locale            string    `fake:"{languageabbreviation}" json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          uint      `fake:"{uint8}"                json:"shardkey"`
	SmId              uint      `fake:"{uint32}"               json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          uint      `fake:"{uint8}"                json:"oof_shard"`
}

type Delivery struct {
	Name    string `fake:"{name}"    json:"name"`
	Phone   string `fake:"{phone}"   json:"phone"`
	Zip     uint   `fake:"{zip}"     json:"zip"`
	City    string `fake:"{city}"    json:"city"`
	Address string `fake:"{address}" json:"address"`
	Region  string `fake:"{state}"   json:"region"`
	Email   string `fake:"{email}"   json:"email"`
}

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestID    uint    `json:"request_id"`
	Currency     string  `fake:"{currencyshort}"            json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float64 `fake:"{price:0.0001,99999.0}"     json:"amount"`
	PaymentDT    uint    `fake:"{second}"                   json:"payment_dt"`
	Bank         string  `fake:"{randomstring:[Alfa,Sber]}" json:"bank"`
	DeliveryCost float64 `fake:"{price:0.0001,99999}"       json:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total"`
	CustomFee    float64 `fake:"{price:0.0001,99999}"       json:"custom_fee"`
}

type Item struct {
	ChrtID      uint    `fake:"{uint64}"               json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       float64 `fake:"{price:0.0001,99999.0}" json:"price"`
	RID         string  `json:"rid"`
	Name        string  `fake:"{productname}"          json:"name"`
	Sale        uint    `fake:"{uint16}"               json:"sale"`
	Size        uint    `fake:"{uint8}"                json:"size"`
	TotalPrice  float64 `fake:"{price:0.0001,99999.0}" json:"total_price"`
	NmID        uint    `fake:"{uint64}"               json:"nm_id"`
	Brand       string  `fake:"{company}"              json:"brand"`
	Status      uint    `fake:"{httpstatuscode}"       json:"status"`
}
