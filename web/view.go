package web

import (
	"net/mail"
	"strings"
	"time"

	"github.com/sad-pixel/mailbin/repository"
)

type DisplayEmail struct {
	Id      int
	Subject string
	From    string
	To      string
	Date    string
}

func FormatDate(t *time.Time) string {
	return t.Format("Jan 2, 15:04")
}

func FormatAddresses(addrs []*mail.Address) string {
	var adds []string
	for _, a := range addrs {
		s := a.String()
		adds = append(adds, s)
	}
	return strings.Join(adds, ", ")
}

func NewDisplayEmail(mail repository.Message) DisplayEmail {
	d := DisplayEmail{
		Id:      mail.Id,
		Subject: mail.Email.Subject,
		From:    FormatAddresses(mail.Email.From),
		To:      FormatAddresses(mail.Email.To),
		Date:    FormatDate(&mail.Email.Date),
	}
	return d
}

func ToDisplayEmails(mails []repository.Message) (ds []DisplayEmail) {
	for _, m := range mails {
		d := NewDisplayEmail(m)
		ds = append(ds, d)
	}

	return ds
}
