package web

import (
	"net/mail"
	"time"

	"github.com/sad-pixel/mailbin/repository"
)

// ApiDisplayEmail holds the representation of an
// email message as returned by the API
type ApiDisplayEmail struct {
	Id          int                    `json:"id"`
	Subject     string                 `json:"subject"`
	From        []string               `json:"from"`
	To          []string               `json:"to"`
	Date        time.Time              `json:"date"`
	Message     string                 `json:"message"`
	Attachments []ApiDisplayAttachment `json:"attachments"`
}

// ApiDisplayAttachment holds the representation of an attachment
// as returned by the API
type ApiDisplayAttachment struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Length      int    `json:"length"`
}

// AddressesToStringArray converts an array of mail.Address into an
// an array of strings with format "Firstname Lastname <email@address.com>"
func AddressesToStringArray(addrs []*mail.Address) (as []string) {
	for _, a := range addrs {
		as = append(as, a.String())
	}
	return as
}

// NewApiDisplayAttachment creates a ApiDisplayAttachment from a repository.MessageAttachment,
// formatting all the fields as required
func NewApiDisplayAttachment(a *repository.MessageAttachment) ApiDisplayAttachment {
	return ApiDisplayAttachment{
		FileName:    a.Attachment.Filename,
		ContentType: a.Attachment.ContentType,
		Length:      len(a.File),
	}
}

// ToApiDisplayAttachments creates an array of ApiDisplayAttachment from an array
// of repository.MessageAttachment
func ToApiDisplayAttachments(attachments []*repository.MessageAttachment) (ada []ApiDisplayAttachment) {
	for _, ma := range attachments {
		ada = append(ada, NewApiDisplayAttachment(ma))
	}

	return ada
}

// NewApiDisplayEmail creates a ApiDisplayEmail from a repository.Message, formatting
// all the fields as required
func NewApiDisplayEmail(mail repository.Message) ApiDisplayEmail {
	return ApiDisplayEmail{
		Id:          mail.Id,
		Subject:     mail.Email.Subject,
		From:        AddressesToStringArray(mail.Email.From),
		To:          AddressesToStringArray(mail.Email.To),
		Date:        mail.Email.Date,
		Attachments: ToApiDisplayAttachments(mail.Attachments),
	}
}
