package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sad-pixel/mailbin/email"
	"github.com/sad-pixel/mailbin/repository"
	"github.com/sad-pixel/mailbin/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("could not read config files, will use default configuration instead")
		os.Setenv("MAILBIN_SMTP_PORT", "1025")
		os.Setenv("MAILBIN_HTTP_PORT", "1026")
	}

	emailRepository := &repository.EmailRepository{}
	emailSettings := &email.Settings{
		Port:              ":" + os.Getenv("MAILBIN_SMTP_PORT"),
		Host:              "localhost",
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxMessageBytes:   1024 * 1024,
		MaxRecipients:     50,
		AllowInsecureAuth: true,
	}
	go email.ListenAndServe(emailSettings, emailRepository)

	webSettings := &web.WebSettings{
		Port: ":" + os.Getenv("MAILBIN_HTTP_PORT"),
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
