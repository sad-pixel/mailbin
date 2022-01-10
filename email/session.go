package email

import (
	"io"
	"log"

	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-smtp"
)

// A Session is returned after EHLO.
type Session struct {
	Repository *EmailRepository
}

func (s *Session) StoreEmail(r io.Reader) error {
	email, err := parsemail.Parse(r) // returns Email struct and error
	if err != nil {
		return err
	}

	// log.Println(email.Subject)
	// log.Println(email.From)
	log.Println(email.To)
	if err := s.Repository.Store(email); err != nil {
		return err
	}
	return nil
}

func (s *Session) Data(r io.Reader) error {
	s.StoreEmail(r)
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	return nil
}

func (s *Session) Rcpt(to string) error {
	return nil
}
