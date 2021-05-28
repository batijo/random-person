package handlers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/batijo/random-person/app/ssp"
	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/utils"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *database.Database
}

func (h *Handlers) Api(c *fiber.Ctx) error {
	c.AcceptsLanguages("lt", "en")
	c.SendString("Hello")
	return nil
}

func (h *Handlers) Name(c *fiber.Ctx) error {
	p := c.Params("gender")
	name := h.DB.RandomName(getGender(p))
	err := c.JSON(fiber.Map{
		"name": name.Name,
	})
	return err
}

func (h *Handlers) Surname(c *fiber.Ctx) error {
	p := c.Params("gender")
	gender := getGender(p)
	surname := h.DB.RandomSurname()
	if gender == 1 {
		ms := c.Query("m_status")
		maritalStatus, err := strconv.Atoi(ms)
		if err == nil && utils.StringContainsInt("0 1 2", maritalStatus) {
			surname.Surname = ssp.Feminize(surname.Surname, uint(maritalStatus))
		} else {
			rand.Seed(time.Now().UnixNano())
			maritalStatus = rand.Intn(3)
			surname.Surname = ssp.Feminize(surname.Surname, uint(maritalStatus))
		}
	} else if gender != 0 {
		rand.Seed(time.Now().UnixNano())
		gender := rand.Intn(2)
		if gender != 0 {
			surname.Surname = ssp.Feminize(surname.Surname, uint(rand.Intn(3)))
		}
	}
	return c.JSON(fiber.Map{
		"surname": surname.Surname,
	})
}

func (h *Handlers) Person(c *fiber.Ctx) error {
	return nil
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}
