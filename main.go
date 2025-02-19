package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/proconlon/redelivery-core/email"
	"github.com/proconlon/redelivery-core/storage"
)

// API to fetch orders
func ordersHandler(w http.ResponseWriter, r *http.Request) {
	orders := storage.LoadOrders()
	json.NewEncoder(w).Encode(orders)
}

func main() {
	log.Println("Starting Re:Delivery...")

	// Load email credentials from dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	emailClient := email.EmailClient{
		Username: os.Getenv("EMAIL_USER"),
		Password: os.Getenv("EMAIL_PASS"), // Use app password for Gmail
		Server:   "imap.gmail.com:993",
	}

	// Fetch emails
	emailClient.FetchEmails()

	http.HandleFunc("/orders", ordersHandler)

	log.Println("Starting Re:Delivery API on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
