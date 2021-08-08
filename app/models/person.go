package models

import "time"

type Person struct {
	*NameOnly
	*Surname
	Gender    bool
	BirthDate time.Time `json:"birth_date,omitempty"`
	Email     string    `json:"email,omitempty"`
}
