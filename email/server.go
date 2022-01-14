package email

import (
	"log"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/sad-pixel/mailbin/repository"
)

type EmailSettings struct {
	Port              string
	Host              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxMessageBytes   int
	MaxRecipients     int
	AllowInsecureAuth bool
}

func ListenAndServe(settings *EmailSettings, repo *repository.EmailRepository) error {
	be := &Backend{repo}
	be.StartStats()

	s := smtp.NewServer(be)
	s.Addr = settings.Port
	s.Domain = settings.Host
	s.ReadTimeout = settings.ReadTimeout
	s.WriteTimeout = settings.WriteTimeout
	s.MaxMessageBytes = settings.MaxMessageBytes
	s.MaxRecipients = settings.MaxRecipients
	s.AllowInsecureAuth = settings.AllowInsecureAuth

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
