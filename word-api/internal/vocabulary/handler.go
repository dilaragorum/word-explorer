package vocabulary

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, vocabulary Vocabulary) error
	Filter(ctx context.Context, args SearchArgs) (FilterResult, error)
}

type handler struct {
	service Service
}

func NewHandler(f *fiber.App, service Service) {
	h := handler{service: service}
	f.Post("/api/v1/words", h.Create)
	f.Get("/api/v1/words", h.Filter)
}

func (h *handler) Create(c *fiber.Ctx) error {
	v := new(Vocabulary)
	if err := c.BodyParser(v); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.service.Create(c.Context(), *v); err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) Filter(c *fiber.Ctx) error {
	var args SearchArgs
	var err error

	args.Size, _ = strconv.Atoi(c.Query("size"))
	args.Page, _ = strconv.Atoi(c.Query("page"))
	args.SearchWord = c.Query("search_word")

	args.SetDefaults()

	vocabularies, err := h.service.Filter(c.Context(), args)
	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(vocabularies)
}
