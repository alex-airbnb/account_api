package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	postgresDB, err := client.SetUpPostgres()

	if err != nil {
		log.Panic(err)
	}

	app := resthandler.NewRESTHandler(
		usecase.NewAccountREST(
			adapter.NewPostgres(postgresDB),
		),
	)

	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Panic(err)
		}
	}()

	shutdownChannnel := make(chan os.Signal)

	signal.Notify(shutdownChannnel, os.Interrupt, os.Kill, syscall.SIGTERM)

	shutdownNotification := <-shutdownChannnel

	log.Println("Received Shutdown Notification", shutdownNotification)
}
