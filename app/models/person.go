package models

import "time"

type Person struct {
	*NameOnly
	*SurnameOnly
	Gender          uint      `json:"gender,omitempty"`
	BirthDate       time.Time `json:"-"`
	BirthDateString string    `json:"birth_date,omitempty"`
	Email           string    `json:"email,omitempty"`
}

func (p *Person) StringifyBirthDate() {
	p.BirthDateString = p.BirthDate.Format("2006-01-02")
}
