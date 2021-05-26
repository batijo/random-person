package handlers

import (
	"github.com/batijo/random-person/database"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *database.Database
}

func (h *Handlers) Api(c *fiber.Ctx) error {
	c.AcceptsLanguages("lt", "en")
	c.SendString(h.DB.RandomName(0).Name)
	return nil
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}
