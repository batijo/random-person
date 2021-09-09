package main

import (
	"log"
	"os"

	"github.com/batijo/random-person/app/email"
	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/server"
)

const (
	configFolder   = "config/"
	namesFile      = "names.json"
	surnamesFile   = "surnames.json"
	configEnv      = ".env"
	emailDomains   = "email_domains.json"
	emailTemplates = "email_templates.json"
)

func main() {
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
	err = email.LoadDomains(configFolder + emailDomains)
	if err != nil {
		log.Println("warning: error loading email domains data")
	}
	err = email.LoadTemplates(configFolder + emailTemplates)
	if err != nil {
		log.Println("warning: error loading email templates data")
	}
	srv := server.New(&db)
	db.InsertData(configFolder, namesFile, surnamesFile)
	log.Println("Server is running on: ", os.Getenv("RP_IP")+":"+os.Getenv("RP_PORT"))
	log.Fatal(srv.App.Listen(":" + os.Getenv("RP_PORT")))
}
