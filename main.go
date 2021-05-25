package main

import (
	"log"
	"os"

	"github.com/batijo/random-person/app"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("cannot load .env file. error: ", err)
	}
}

func main() {
	srv := app.SetUpServer()
	api := srv.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Redirect("api/v0/")
	})
	v0 := api.Group("/v0")
	v0.Get("", func(c *fiber.Ctx) error {
		c.AcceptsLanguages("lt", "en")
		c.SendString("Hello")
		return nil
	})
	app.SetUpMidleware(srv)
	log.Println("Server is running...")
	//log.Fatal(srv.ListenTLS(os.Getenv("RP_IP") + ":" + os.Getenv("RP_PORT"), os.Getenv("RP_CERT_FILE"), os.Getenv("RP_KEY_FILE")))
	log.Fatal(srv.Listen(os.Getenv("RP_IP") + ":" + os.Getenv("RP_PORT")))
}
