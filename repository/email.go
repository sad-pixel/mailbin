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
	Id          int
	Email       parsemail.Email
	Attachments []MessageAttachment
	isRead      bool
}

// NewMessage creates a new Message from a parsemail.Email
func NewMessage(email parsemail.Email) Message {
	return Message{
		Email:  email,
		isRead: false,
	}
}

// IsRead returns if the email has been read by the user
func (m *Message) IsRead() bool {
	return m.isRead
}

// MarkRead marks the message as read
func (m *Message) MarkRead() {
	m.isRead = true
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
	mu       sync.RWMutex
	Messages []Message
}

// Store parses and stores an email from a io.Reader into memory
func (e *EmailRepository) Store(r io.Reader) error {

	mail, err := parsemail.Parse(r)
	if err != nil {
		return fmt.Errorf("could not parse email: %v", err)
	}

	message := NewMessage(mail)

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

	message.Id = len(e.Messages) + 1
	e.Messages = append(e.Messages, message)
	return nil
}

// GetAll returns all emails in the repository
func (e *EmailRepository) GetAll() []Message {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.Messages
}

// GetOne returns one email by index from the repository
func (e *EmailRepository) GetOne(index int) (Message, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if index > len(e.Messages) || index < 0 {
		return Message{}, fmt.Errorf("email index out of bounds")
	}

	return e.Messages[index], nil
}

// MarkRead marks an email in the repository as read.
func (e *EmailRepository) MarkRead(index int) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if index > len(e.Messages) || index < 0 {
		return fmt.Errorf("email index out of bounds")
	}

	e.Messages[index].MarkRead()
	return nil
}
