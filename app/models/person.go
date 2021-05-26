package models

type Person struct {
	*NameOnly
	*Surname
	Email string `json:"email,omitempty"`
}
