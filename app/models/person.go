package models

import "time"

type Person struct {
	*NameOnly
	*SurnameOnly
	Gender    bool      `json:"gender,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Email     string    `json:"email,omitempty"`
}
