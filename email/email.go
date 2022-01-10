package email

import (
	"io"

	"github.com/DusanKasan/parsemail"
)

type EmailMessage struct {
	Email       parsemail.Email
	Attachments []EmailAttachment
}

type EmailAttachment struct {
	Attachment parsemail.Attachment
	File       []byte
}

type EmailRepository struct {
	Messages []EmailMessage
}

func (e *EmailRepository) Store(mail parsemail.Email) error {
	message := EmailMessage{
		Email: mail,
	}

	for _, attachment := range message.Email.Attachments {
		a := EmailAttachment{
			Attachment: attachment,
		}

		attachmentBytes, err := io.ReadAll(attachment.Data)
		if err != nil {
			return err
		}
		a.File = attachmentBytes
		// log.Println("stored ", len(attachmentBytes), " bytes attachment")
		message.Attachments = append(message.Attachments, a)
	}

	e.Messages = append(e.Messages, message)
	return nil
}
