package vocabulary

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, vocabulary *Vocabulary) error
	Filter(ctx context.Context, args SearchArgs) (*[]Vocabulary, error)
}

type handler struct {
	service Service
}

func NewHandler(f *fiber.App, service Service) {
	h := handler{service: service}
	f.Post("/api/v1/words", h.Create)
	f.Get("/api/v1/words", h.Get)
}

func (h *handler) Create(c *fiber.Ctx) error {
	v := new(Vocabulary)
	if err := c.BodyParser(v); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.service.Create(c.Context(), v); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something wrong")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) Get(c *fiber.Ctx) error {
	// TODO:Burada nasıl alırım from ve size'ı? Get metodu body almamalı-Parametre olarak mı almak mantıklı? --> Bunu query parametresi olarak al.
	sargs := SearchArgs{}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid size value")
	}

	sargs.Size = size

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid size value")
	}
	sargs.Page = page

	sargs.SubWord = c.Query("sub_word")

	vocabularies, err := h.service.Filter(c.Context(), sargs)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(vocabularies)
}
