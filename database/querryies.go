package database

import (
	"github.com/batijo/random-person/app/models"
)

// gender must be 0 - male or 1 - female else it will return random name
func (db *Database) RandomName(gender int) models.Name {
	return db.randomName("", gender)
}

func (db *Database) RandomNameNormativeStatus(normativeStatus string, gender int) models.Name {
	return db.randomName(normativeStatus, gender)
}

func (db *Database) RandomSurname() models.Surname {
	var surname models.Surname
	db.Limit(1).Order("RANDOM()").Find(&surname)
	return surname
}

func (db *Database) randomName(normativeStatus string, gender int) models.Name {
	var name models.Name
	if normativeStatus != "" {
		if gender >= 0 && gender < 2 {
			db.Limit(1).Order("RANDOM()").Where("gender = ?", gender).Where("normative_status = ?", normativeStatus).Find(&name)
		} else {
			db.Limit(1).Order("RANDOM()").Where("normative_status = ?", normativeStatus).Find(&name)
		}
	} else {
		if gender >= 0 && gender < 2 {
			db.Limit(1).Order("RANDOM()").Where("gender = ?", gender).Find(&name)
		} else {
			db.Limit(1).Order("RANDOM()").Find(&name)
		}
	}
	return name
}
