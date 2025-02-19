package email

import (
	"log"
	"regexp"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"

	"github.com/proconlon/redelivery-core/storage"
)

// EmailClient struct to handle email operations
type EmailClient struct {
	Username string
	Password string
	Server   string
}

// Regex patterns for tracking numbers
var trackingPatterns = []string{
	`(\b[1-9A-Z]{10,}\b)`,  // Generic tracking number pattern
	`\b(1Z[0-9A-Z]{16})\b`, // UPS
	`\b(\d{20,22})\b`,      // FedEx / USPS
}

// Extract tracking numbers from subject
func extractTracking(subject string) string {
	for _, pattern := range trackingPatterns {
		re := regexp.MustCompile(pattern)
		if match := re.FindString(subject); match != "" {
			return match
		}
	}
	return ""
}

// Connect and fetch emails
func (ec *EmailClient) FetchEmails() {
	// Connect to IMAP server
	c, err := client.DialTLS(ec.Server, nil)
	if err != nil {
		log.Fatal("Failed to connect to IMAP server:", err)
	}
	defer c.Logout()

	// Login
	if err := c.Login(ec.Username, ec.Password); err != nil {
		log.Fatal("Login failed:", err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal("Failed to select INBOX:", err)
	}
	log.Printf("INBOX contains %d messages\n", mbox.Messages)

	// Fetch last 10 emails
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(uint32(mbox.Messages-9), uint32(mbox.Messages))

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	for msg := range messages {
		log.Printf("From: %s | Subject: %s\n", msg.Envelope.From[0].Address(), msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal("Failed to fetch emails:", err)
	}

	for msg := range messages {
		subject := msg.Envelope.Subject
		tracking := extractTracking(subject)

		if tracking != "" {
			log.Printf("Found tracking number: %s in subject: %s\n", tracking, subject)
		}
	}

	var orders []storage.Order
	for msg := range messages {
		subject := msg.Envelope.Subject
		tracking := extractTracking(subject)

		if tracking != "" {
			merchant := "Unknown"
			if strings.Contains(strings.ToLower(subject), "amazon") {
				merchant = "Amazon"
			}

			order := storage.Order{
				Merchant:   merchant,
				TrackingID: tracking,
				Status:     "Shipped",
			}
			orders = append(orders, order)
			log.Printf("Order detected: %+v\n", order)
		}
	}

	storage.SaveOrders(orders)
}
