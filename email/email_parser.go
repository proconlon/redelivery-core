package email

import (
	"log"
	"os"

	"github.com/mnako/letters"
)

func main() {
	rawEmail, err := os.Open("email.eml")
	if err != nil {
		log.Fatal("error while reading email from file: %w", err)
		return
	}

	defer func() {
		if err := rawEmail.Close(); err != nil {
			log.Fatal("error while closing rawEmail: %w", err)
			return
		}
	}()

	email, err := letters.ParseEmail(rawEmail)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("text:", email.Text)
}
