package main

import (
	"log"
	"os"

	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/server"
	"github.com/joho/godotenv"
)

const (
	configFolder = "config/"
	namesFile    = "names.json"
	surnamesFile = "surnames.json"
	configEnv    = ".env"
)

func main() {
	if err := godotenv.Load(configFolder + configEnv); err != nil {
		log.Fatal("cannot load .env file. error: ", err)
	}
	db, err := database.Connect(database.Config{
		Host:    os.Getenv("RP_DB_HOST"),
		Name:    os.Getenv("RP_DB_NAME"),
		User:    os.Getenv("RP_DB_USER"),
		Pasword: os.Getenv("RP_DB_PASSWORD"),
		Port:    os.Getenv("RP_DB_PORT"),
	})
	if err != nil {
		log.Panic(err)
	}
	srv := server.New(&db)
	db.InsertData(configFolder, namesFile, surnamesFile)
	log.Println("Server is running on: ", os.Getenv("RP_IP")+":"+os.Getenv("RP_PORT"))
	log.Fatal(srv.App.Listen(":" + os.Getenv("RP_PORT")))
}
