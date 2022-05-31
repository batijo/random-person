package handlers

import (
	"math/rand"
	"os"
	"time"

	"github.com/batijo/random-person/app/age"
	"github.com/batijo/random-person/app/email"
	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/utils"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *database.Database
}

func (h *Handlers) Api(c *fiber.Ctx) error {
	err := c.JSON(fiber.Map{
		"app":     "github.com/batijo/random-person",
		"api":     "https://github.com/batijo/random-person/tree/release#api-usage",
		"version": os.Getenv("RP_VERSION"),
	})
	c.Set("content-type", "application/json; charset=utf-8")
	return err
}

func (h *Handlers) Name(c *fiber.Ctx) error {
	name := h.DB.RandomNameNormativeStatus(os.Getenv("RP_DEF_NORMATIVE_STAT"), getGender(c.Params("gender")))
	err := c.JSON(fiber.Map{
		"name": name.Name,
	})
	c.Set("content-type", "application/json; charset=utf-8")
	return err
}

func (h *Handlers) Surname(c *fiber.Ctx) error {
	q := new(surnConf)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	surname := q.randomSurname(h.DB, getGender(c.Params("gender")))
	err := c.JSON(fiber.Map{
		"surname": surname.Surname,
	})
	c.Set("content-type", "application/json; charset=utf-8")
	return err
}

func (h *Handlers) Person(c *fiber.Ctx) error {
	gender := getGender(c.Params("gender"))
	if !utils.StringContainsInt("0 1", gender) {
		rand.Seed(time.Now().UnixNano())
		gender = rand.Intn(2)
	}
	var person models.Person
	person.NameOnly = &models.NameOnly{Name: h.DB.RandomNameNormativeStatus(os.Getenv("RP_DEF_NORMATIVE_STAT"), gender).Name}
	q := new(surnConf)
	if err := c.QueryParser(q); err != nil {
		return err
	}
	person.SurnameOnly = &models.SurnameOnly{Surname: q.randomSurname(h.DB, gender).Surname}
	age.Random(&person)
	email.Random(&person)
	err := c.JSON(person)
	c.Set("content-type", "application/json; charset=utf-8")
	return err
}

func (h *Handlers) Version(c *fiber.Ctx) error {
	err := c.JSON(os.Getenv("RP_VERSION"))
	c.Set("content-type", "application/json; charset=utf-8")
	return err
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}
