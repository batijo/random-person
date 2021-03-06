package models

type Name struct {
	ID              uint   `json:"-" gorm:"primaryKey"`
	Gender          uint   `json:"gender" gorm:"type:smallint"`
	NormativeStatus string `json:"normative_status"`
	Origin          string `json:"origin,omitempty"`
	Note            string `json:"note,omitempty"`
	*NameOnly
}

type NameOnly struct {
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
