package sp

import (
	"testing"
)

func Test_RemoveSuffix(t *testing.T) {
	tests := []struct {
		val string
		exp string
	}{
		{"Kazlauskas", "Kazlausk"},
		{"Stankevičius", "Stankevič"},
		{"Butkus", "Butk"},
		{"Adomaitis", "Adomait"},
		{"Kairys", "Kair"},
		{"Juška", "Jušk"},
		{"Dirsė", "Dirs"},
		{"Stundžia", "Stundž"},
		{"Džiūve", "Džiūv"},
		{"Šlymuo", "Šlym"},
		{"Talat", "Talat"},
		{"Kupcias", "Kupc"},
		{"Šaltais", "Šalt"},
		{"Tarlo", "Tarl"},
		{"Taulai", "Taul"},
		{"Kalnau", "Kaln"},
	}
	for _, d := range tests {
		res := RemoveSuffix(d.val)
		if res != d.exp {
			t.Errorf(
				"DATA: %v EXPECTED: %v, GOT: %v",
				d.val,
				d.exp,
				res,
			)
		}
	}
}
