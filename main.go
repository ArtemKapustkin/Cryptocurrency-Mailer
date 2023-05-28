package main

import (
	"bitcoin-exchange-rate/internal/handler"
	"bitcoin-exchange-rate/internal/repository"
	"bitcoin-exchange-rate/internal/service"
	"bitcoin-exchange-rate/pkg/mailer"
	"bitcoin-exchange-rate/pkg/parser"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	cryptoParser := parser.NewBinanceCryptoParser()

	subscriberReposity := repository.NewSubscriberFileRepository(os.Getenv("EMAILSFILEPATH"))
	mailer := mailer.NewMailer("smtp.gmail.com", "587")

	mailerService := service.NewMailerService(subscriberReposity, mailer)

	rateHandler := handler.NewRateHandler(cryptoParser)
	app.Get("/rate", rateHandler.GetExchangeRate)

	mailerHandler := handler.NewMailerHandler(mailerService, cryptoParser, subscriberReposity, validator.New())
	app.Post("/sendEmails", mailerHandler.SendExchangeRate)

	app.Post("/subscribe", mailerHandler.Subscribe)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
