package main

import (
	"log"
	"os"

	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("cannot load .env file. error: ", err)
	}
	srv := server.New(&database.Database{})
	log.Println("Server is running...")
	//log.Fatal(srv.ListenTLS(os.Getenv("RP_IP") + ":" + os.Getenv("RP_PORT"), os.Getenv("RP_CERT_FILE"), os.Getenv("RP_KEY_FILE")))
	log.Fatal(srv.App.Listen(os.Getenv("RP_IP") + ":" + os.Getenv("RP_PORT")))
}
