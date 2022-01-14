package email

import (
	"log"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/sad-pixel/mailbin/repository"
)

// The Backend implements SMTP server methods.
type Backend struct {
	Repository *repository.EmailRepository
}

// Starts a goroutine that prints statistics every 15 seconds
func (bkd *Backend) StartStats() {
	go func(rp *repository.EmailRepository) {
		for {
			log.Println("--- Stats ---")
			log.Println("Emails held: ", len(rp.Messages))
			attachSize := 0
			attachCount := 0
			for _, v := range rp.Messages {
				for _, b := range v.Attachments {
					attachCount++
					attachSize += len(b.File)
				}
			}
			log.Println("Attachments held: ", attachCount)
			log.Println("Total attachments size: ", attachSize)
			time.Sleep(15 * time.Second)
		}
	}(bkd.Repository)
}

// Creates a new session on SMTP Login. Can authenticate user in the future if needed.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &Session{bkd.Repository}, nil
}

// Creates a new session for SMTP Anonymous Login
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{bkd.Repository}, nil
}
