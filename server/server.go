package server

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/batijo/random-person/app/handlers"
	"github.com/batijo/random-person/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Obj struct {
	*fiber.App

	db *database.Database
}

func (srv *Obj) registerApiRoutes(r fiber.Router) {
	h := handlers.New(srv.db)
	r.Get("/name/:gender?", h.Name)
	r.Get("/surname/:gender?", h.Surname)
	r.Get("/person/:gender?", h.Person)
	r.Get("/version", h.Version)
}

func (srv *Obj) registerWebRoutes() {
	h := handlers.New(srv.db)
	srv.App.Get("/", h.Api)
}

func New(db *database.Database) *Obj {
	srv := Obj{
		App: fiber.New(fiber.Config{
			GETOnly:               true,
			ServerHeader:          "random-person",
			DisableStartupMessage: true,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(500).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		}),
		db: db,
	}
	maxConnsPerIP, err := strconv.Atoi(os.Getenv("RP_MAX_CONNS_PER_IP"))
	if err != nil {
		log.Fatal("error: ", err)
	}
	srv.Server().MaxConnsPerIP = maxConnsPerIP
	srv.Static("/", "./public")
	srv.registerMiddleware()
	srv.registerWebRoutes()
	api := srv.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Redirect("api/v0/")
	})
	v0 := api.Group("/v0", func(c *fiber.Ctx) error {
		c.Set("Version", os.Getenv("RP_VERSION"))
		return c.Next()
	})
	srv.registerApiRoutes(v0)
	srv.statusNotFoundMiddleware()
	return &srv
}

func (srv *Obj) registerMiddleware() {
	// Panic recover middleware
	srv.Use(recover.New())
	// Cache favicon to prevent frequent disk access
	srv.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))
	// Middleware to limit multiple requests
	maxRequests, err := strconv.Atoi(os.Getenv("RP_MAX_REQUESTS"))
	if err != nil {
		log.Panic(err)
	}
	requestExp, err := strconv.Atoi(os.Getenv("RP_MAX_REQUEST_EXP"))
	if err != nil {
		log.Panic(err)
	}
	srv.Use(limiter.New(limiter.Config{
		Max:        maxRequests,
		Expiration: time.Duration(requestExp) * time.Second,
	}))
}

func (srv *Obj) statusNotFoundMiddleware() {
	// Status Not Found middleware
	srv.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			map[string]interface{}{"message": "page not found"},
		)
	})
}
