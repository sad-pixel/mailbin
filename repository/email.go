package repository

import (
	"fmt"
	"io"
	"sync"

	"github.com/DusanKasan/parsemail"
)

// Message represents an e-mail message along with all data
// and attachments
type Message struct {
	Email       parsemail.Email
	Attachments []MessageAttachment
}

// MessageAttachment represents an attachment, along with the file data
// that is attached to a Message
type MessageAttachment struct {
	Attachment parsemail.Attachment
	File       []byte
}

// EmailRepository provides mechanisms for storing and retrieving
// email messages
type EmailRepository struct {
	mu       sync.Mutex
	Messages []Message
}

// Store parses and stores an email from a io.Reader into memory
func (e *EmailRepository) Store(r io.Reader) error {

	mail, err := parsemail.Parse(r)
	if err != nil {
		return fmt.Errorf("could not parse email: %v", err)
	}

	message := Message{
		Email: mail,
	}

	for _, attachment := range message.Email.Attachments {
		a := MessageAttachment{
			Attachment: attachment,
		}

		attachmentBytes, err := io.ReadAll(attachment.Data)
		if err != nil {
			return fmt.Errorf("could not read attachment bytes: %v", err)
		}

		a.File = attachmentBytes
		message.Attachments = append(message.Attachments, a)
	}

	// acquire lock only after other processing is done
	e.mu.Lock()
	defer e.mu.Unlock()

	e.Messages = append(e.Messages, message)
	return nil
}

// GetAll returns all emails in the repository
func (e *EmailRepository) GetAll() *[]Message {
	e.mu.Lock()
	defer e.mu.Unlock()

	return &e.Messages
}

// GetOne returns one email by index from the repository
func (e *EmailRepository) GetOne(index int) (*Message, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if index > len(e.Messages) || index < 0 {
		return nil, fmt.Errorf("email index out of bounds")
	}

	return &e.Messages[index], nil
}
