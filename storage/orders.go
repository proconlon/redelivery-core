package storage

import (
	"encoding/json"
	"log"
	"os"
)

type Order struct {
	Merchant     string `json:"merchant"`
	TrackingID   string `json:"tracking_id"`
	Status       string `json:"status"`
	DeliveryDate string `json:"delivery_date"`
}

type OrderDB struct {
	Orders []Order `json:"orders"`
}

var dbFile = "orders.json"

// Save orders to JSON
func SaveOrders(orders []Order) {
	file, err := os.Create(dbFile)
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer file.Close()

	jsonData, _ := json.MarshalIndent(OrderDB{Orders: orders}, "", "  ")
	file.Write(jsonData)
	log.Println("Orders saved successfully.")
}

// Load orders from JSON
func LoadOrders() []Order {
	file, err := os.Open(dbFile)
	if err != nil {
		log.Println("No existing orders found, starting fresh.")
		return []Order{}
	}
	defer file.Close()

	var orderDB OrderDB
	json.NewDecoder(file).Decode(&orderDB)
	return orderDB.Orders
}
