package models

type Surname struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Surname string `json:"surname" gorm:"uniqueIndex"`
}

func (n *Surname) TableName() string {
	return "Surnames"
}

func RemoveDuplicateSurnames(surnames []Surname) []Surname {
	var newSurnames []Surname
	for _, s := range surnames {
		if !containsSurname(newSurnames, s) {
			newSurnames = append(newSurnames, s)
		}
	}
	return newSurnames
}

func containsSurname(arr []Surname, s Surname) bool {
	for _, a := range arr {
		if a.Surname == s.Surname {
			return true
		}
	}
	return false
}
