package web

import (
	"github.com/sad-pixel/mailbin/repository"
)

type Email struct {
	Id      int
	Subject string
	From    string
	To      string
	Time    string
}

func GetOneEmail(r *repository.EmailRepository, index int) (Email, error) {
	m, err := r.GetOne(index)
	if err != nil {
		return Email{}, err
	}

	e := Email{
		Id:      index,
		Subject: m.Email.Subject,
		Time:    m.Email.Date.String(),
	}
	return e, nil
}
