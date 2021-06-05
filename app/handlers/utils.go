package handlers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/app/ssp"
	"github.com/batijo/random-person/database"
	"github.com/batijo/random-person/utils"
)

type Handlers struct {
	DB *database.Database
}

type surnConf struct {
	MaritalStatus string `query:"m_status"`
}

func New(db *database.Database) *Handlers {
	return &Handlers{
		DB: db,
	}
}

func (s *surnConf) randomSurname(db *database.Database, gender int) models.Surname {
	surname := db.RandomSurname()
	if gender == 1 {
		maritalStatus, err := strconv.Atoi(s.MaritalStatus)
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
	return surname
}

func getGender(query string) int {
	if query == "" {
		return -1
	}
	gender, err := strconv.Atoi(query)
	if err != nil {
		switch query {
		case "male":
			return 0
		case "female":
			return 1
		default:
			return -1
		}
	}
	return gender
}
