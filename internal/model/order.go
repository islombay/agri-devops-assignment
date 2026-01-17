package model

import "time"

type Order struct {
	ID         string    `json:"id"`
	FarmerID   string    `json:"farmer_id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	OrderDate  time.Time `json:"order_date"`
	Status     string    `json:"status"`
}
