package models

type Name struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

func (n *Name) TableName() string {
	return "Names"
}

func RemoveDuplicateNames(names []Name) []Name {
	var newNames []Name
	for _, n := range names {
		if !containsName(newNames, n) {
			newNames = append(newNames, n)
		}
	}
	return newNames
}

func containsName(arr []Name, s Name) bool {
	for _, a := range arr {
		if a.Name == s.Name {
			return true
		}
	}
	return false
}
