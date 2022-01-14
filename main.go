package main

import (
	"time"

	"github.com/sad-pixel/mailbin/email"
	"github.com/sad-pixel/mailbin/repository"
)

func main() {
	emailRepository := &repository.EmailRepository{}
	emailSettings := &email.Settings{
		Port:              ":1025",
		Host:              "localhost",
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxMessageBytes:   1024 * 1024,
		MaxRecipients:     50,
		AllowInsecureAuth: true,
	}
	email.ListenAndServe(emailSettings, emailRepository)
}
