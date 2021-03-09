package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/alex-airbnb/account_api/adapter"
	"github.com/alex-airbnb/account_api/client"
	"github.com/alex-airbnb/account_api/resthandler"
	"github.com/alex-airbnb/account_api/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	go func() {
		if err := client.SetUpPostgres(); err != nil {
			log.Panic(err)
		}
	}()

	app := resthandler.NewRESTHandler(
		usecase.NewAccountREST(
			adapter.NewPostgres(client.Postgres),
		),
	)

	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Panic(err)
		}
	}()

	shutdownChannnel := make(chan os.Signal)

	signal.Notify(shutdownChannnel, os.Interrupt)
	signal.Notify(shutdownChannnel, os.Kill)

	shutdownNotification := <-shutdownChannnel

	log.Println("Received Shutdown Notification", shutdownNotification)

	app.Shutdown()
}
