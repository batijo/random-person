package handlers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/app/ssp"
	"github.com/batijo/random-person/database"
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
	var (
		gender = -1
		err    error
	)
	if p != "" {
		gender, err = strconv.Atoi(p)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "parameter must be number",
			})
			return nil
		}
	}
	var name models.Name
	if gender < 0 || gender > 1 {
		name = h.DB.RandomName(-1)
	} else {
		name = h.DB.RandomName(gender)
	}
	err = c.JSON(fiber.Map{
		"name": name.Name,
	})
	return err
}

func (h *Handlers) Surname(c *fiber.Ctx) error {
	p := c.Params("gender")
	var (
		gender = -1
		err    error
	)
	if p != "" {
		gender, err = strconv.Atoi(p)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "parameter must be number",
			})
			return nil
		}
	}
	surname := h.DB.RandomSurname()
	if gender == 1 {
		ms := c.Query("m_status")
		var maritalStatus int
		if ms != "" {
			maritalStatus, err = strconv.Atoi(ms)
			if err != nil || maritalStatus < 0 || maritalStatus > 2 {
				c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "m_status must be number (0,1,2)",
				})
				return nil
			}
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
	err = c.JSON(fiber.Map{
		"surname": surname.Surname,
	})
	return err
}

func (h *Handlers) Person(c *fiber.Ctx) error {
	return nil
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}
