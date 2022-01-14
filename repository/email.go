package repository

import (
	"fmt"
	"io"
	"sync"

	"github.com/DusanKasan/parsemail"
)

type Message struct {
	Email       parsemail.Email
	Attachments []MessageAttachment
}

type MessageAttachment struct {
	Attachment parsemail.Attachment
	File       []byte
}

type EmailRepository struct {
	mu       sync.Mutex
	Messages []Message
}

func (e *EmailRepository) Store(mail parsemail.Email) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	message := Message{
		Email: mail,
	}

	for _, attachment := range message.Email.Attachments {
		a := MessageAttachment{
			Attachment: attachment,
		}

		attachmentBytes, err := io.ReadAll(attachment.Data)
		if err != nil {
			return err
		}

		a.File = attachmentBytes
		message.Attachments = append(message.Attachments, a)
	}

	e.Messages = append(e.Messages, message)
	return nil
}

func (e *EmailRepository) GetAll() *[]Message {
	e.mu.Lock()
	defer e.mu.Unlock()

	return &e.Messages
}

func (e *EmailRepository) GetOne(index int) (*Message, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if index > len(e.Messages) || index < 0 {
		return nil, fmt.Errorf("email index out of bounds")
	}

	return &e.Messages[index], nil
}
