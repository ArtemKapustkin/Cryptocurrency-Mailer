package handler

import (
	"bitcoin-exchange-rate/internal/model"
	"bitcoin-exchange-rate/internal/repository"
	"bitcoin-exchange-rate/internal/service"
	"bitcoin-exchange-rate/pkg/parser"
	"errors"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SubscriberRepository interface {
	Create(subscriber *model.Subscriber) error
}

type MailerHandler struct {
	mailerService        *service.MailerService
	binanceCryptoParser  *parser.BinanceCryptoParser
	subscriberRepository SubscriberRepository
	validator            *validator.Validate

	exchangeRateBaseCurrency  string
	exchangeRateQuoteCurrency string
}

func NewMailerHandler(
	mailerService *service.MailerService,
	binanceCryptoParser *parser.BinanceCryptoParser,
	subscriberRepository SubscriberRepository,
	validator *validator.Validate,
) *MailerHandler {
	return &MailerHandler{
		mailerService:             mailerService,
		binanceCryptoParser:       binanceCryptoParser,
		subscriberRepository:      subscriberRepository,
		validator:                 validator,
		exchangeRateBaseCurrency:  os.Getenv("BASE_CURRENCY"),
		exchangeRateQuoteCurrency: os.Getenv("QUOTE_CURRENCY"),
	}
}

func (h *MailerHandler) SendExchangeRate(c *fiber.Ctx) error {
	value, err := h.binanceCryptoParser.GetExchangeRate(h.exchangeRateBaseCurrency, h.exchangeRateQuoteCurrency)
	if err != nil {
		return c.SendStatus(500)
	}

	err = h.mailerService.SendValueToAllEmails(strconv.FormatFloat(value, 'f', 2, 64))
	if err != nil {
		return c.SendStatus(400)
	}

	return c.SendStatus(200)
}

func (h *MailerHandler) Subscribe(c *fiber.Ctx) error {
	var payload subscribeDTO

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(500)
	}

	if h.validator.Struct(&payload) != nil {
		return c.SendStatus(400)
	}

	err := h.subscriberRepository.Create(model.NewSubscriber(payload.Email))
	if err != nil {
		if errors.Is(err, repository.ErrEmailAlreadyExist) {
			return c.SendStatus(409)
		}

		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}

type subscribeDTO struct {
	Email string `validate:"required,email"`
}
