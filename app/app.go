package app

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetUpServer() *fiber.App {
	srv := fiber.New(fiber.Config{
		GETOnly:               true,
		ServerHeader:          "github.com/batijo/random-person",
		DisableStartupMessage: true,
	})
	maxConnsPerIP, err := strconv.Atoi(os.Getenv("RP_MAX_CONNS_PER_IP"))
	if err != nil {
		log.Fatal("error: ", err)
	}
	srv.Server().MaxConnsPerIP = maxConnsPerIP
	return srv
}

func SetUpMidleware(srv *fiber.App) {
	// Midleware to limit multiple requests
	srv.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == os.Getenv("RP_IP")
		},
		Max:        5,
		Expiration: 5 * time.Second,
	}))
	// Status Not Found midleware
	srv.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			map[string]interface{}{"message": "page not found"},
		)
	})
}
