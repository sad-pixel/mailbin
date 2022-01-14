package email

import (
	"fmt"
	"io"

	"github.com/emersion/go-smtp"
	"github.com/sad-pixel/mailbin/repository"
)

// A Session is returned after EHLO.
type Session struct {
	Repository *repository.EmailRepository
}

// Data is called after a new email is received
func (s *Session) Data(r io.Reader) error {
	if err := s.Repository.Store(r); err != nil {
		return fmt.Errorf("could not store email: %v", err)
	}
	return nil
}

// Discard currently processed message. Does nothing.
func (s *Session) Reset() {}

// Free all resources associated with session.
func (s *Session) Logout() error {
	return nil
}

// Set return path for currently processed message.
func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	return nil
}

// Add a recipient to the currently processed message
func (s *Session) Rcpt(to string) error {
	return nil
}
