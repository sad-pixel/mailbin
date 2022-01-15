package main

import (
	"log"
	"time"

	"github.com/sad-pixel/mailbin/email"
	"github.com/sad-pixel/mailbin/repository"
	"github.com/sad-pixel/mailbin/web"
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
	go email.ListenAndServe(emailSettings, emailRepository)

	webSettings := &web.WebSettings{
		Port: ":1026",
	}

	webHandlers, err := web.NewHandlers(emailRepository)
	if err != nil {
		log.Fatalln("could not setup web handlers: " + err.Error())
	}

	router := web.NewRouter()
	web.SetupRoutes(router, webHandlers)
	if err := web.ListenAndServe(webSettings, router); err != nil {
		log.Fatalln("could not start web server: " + err.Error())
	}
}
