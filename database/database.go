package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/utils"
	"github.com/jackc/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

type Config struct {
	Host    string
	Name    string
	User    string
	Pasword string
	Port    string
}

func Connect(conf Config) (Database, error) {
	var (
		err error
		db  *gorm.DB
	)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Pasword,
	)
	/* Attempting to connect to database several times
	   because it needs couple of seconds to set up when using docker-compose
	   and may refuse connections */
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return Database{}, err
	}
	db.AutoMigrate(
		&models.Name{},
		&models.Surname{},
	)
	return Database{db}, nil
}

func (db *Database) InsertData(folder, namesFile, surnamesFile string) {
	names := []models.Name{}
	err := utils.LoadData(folder+namesFile, &names)
	if err != nil {
		log.Println(err)
	}
	names = models.RemoveDuplicateNames(names)
	var perr *pgconn.PgError
	if resp := db.Create(&names); resp.Error != nil {
		// Checks if error is duplicate key violation
		if errors.As(resp.Error, &perr) && perr.Code == "23505" {
			log.Printf(
				"%[1]s contains records that are already in the database. You may want to remove %[1]s\n",
				namesFile,
			)
		} else {
			log.Println(resp.Error)
		}
	}
	surnames := []models.Surname{}
	err = utils.LoadData(folder+surnamesFile, &surnames)
	if err != nil {
		log.Println(err)
	}
	surnames = models.RemoveDuplicateSurnames(surnames)
	if resp := db.Create(&surnames); resp.Error != nil {
		if errors.As(resp.Error, &perr) && perr.Code == "23505" {
			log.Printf(
				"%[1]s contains records that are already in the database. You may want to remove %[1]s\n",
				surnamesFile,
			)
		} else {
			log.Println(resp.Error)
		}
	}
}

// gender must be 0 - male or 1 - female else it will return random name
func (db *Database) RandomName(gender int) models.Name {
	var name models.Name
	if gender > 0 || gender < 2 {
		db.Limit(1).Order("RANDOM()").Where("gender = ?", gender).Find(&name)
	}
	db.Limit(1).Order("RANDOM()").Find(&name)
	return name
}

func (db *Database) RandomSurname() models.Surname {
	var surname models.Surname
	db.Limit(1).Order("RANDOM()").Find(&surname)
	return surname
}
