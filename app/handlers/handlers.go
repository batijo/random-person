package handlers

import (
	"math/rand"
	"os"
	"time"

	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/utils"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *database.Database
}

func (h *Handlers) Api(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"message": "github.com/batijo/random-person",
		"version": os.Getenv("RP_VERSION"),
	})
	return nil
}

func (h *Handlers) Name(c *fiber.Ctx) error {
	p := c.Params("gender")
	name := h.DB.RandomName(getGender(p))
	return c.JSON(fiber.Map{
		"name": name.Name,
	})
}

func (h *Handlers) Surname(c *fiber.Ctx) error {
	q := new(surnConf)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	p := c.Params("gender")
	surname := q.randomSurname(h.DB, getGender(p))
	return c.JSON(fiber.Map{
		"surname": surname.Surname,
	})
}

func (h *Handlers) Person(c *fiber.Ctx) error {
	p := c.Params("gender")
	gender := getGender(p)
	if !utils.StringContainsInt("0 1", gender) {
		rand.Seed(time.Now().UnixNano())
		gender = rand.Intn(2)
	}
	name := h.DB.RandomName(gender)
	q := new(surnConf)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	surname := q.randomSurname(h.DB, gender)
	return c.JSON(fiber.Map{
		"name":    name.Name,
		"surname": surname.Surname,
	})
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}
