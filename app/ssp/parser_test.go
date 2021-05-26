package ssp

import (
	"testing"
)

func Test_removeSuffix(t *testing.T) {
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
		res := removeSuffix(d.val)
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

func Test_Feminize(t *testing.T) {
	tests := []struct {
		val           string
		maritalStatus uint
		exp           string
	}{
		{"Kazlauskas", 0, "Kazlauskaitė"},
		{"Stankevičius", 0, "Stankevičiūtė"},
		{"Butkus", 0, "Butkutė"},
		{"Adomaitis", 0, "Adomaitytė"},
		{"Kairys", 0, "Kairytė"},
		{"Juška", 0, "Juškaitė"},
		{"Dirsė", 0, "Dirsaitė"},
		{"Stundžia", 0, "Stundžiūtė"},
		{"Džiūve", 0, "Džiūvytė"},
		{"Šlymuo", 0, "Šlymytė"},
		{"Talat", 0, "Talatytė"},
		{"Kupcias", 0, "Kupcaitė"},
		{"Šaltais", 0, "Šaltaitė"},
		{"Tarlo", 0, "Tarlaitė"},
		{"Taulai", 0, "Taulytė"},
		{"Kalnau", 0, "Kalnaitė"},
		{"Kazlauskas", 1, "Kazlauskienė"},
		{"Stankevičius", 1, "Stankevičienė"},
		{"Butkus", 1, "Butkienė"},
		{"Adomaitis", 1, "Adomaitienė"},
		{"Kairys", 1, "Kairienė"},
		{"Juška", 1, "Juškienė"},
		{"Dirsė", 1, "Dirsienė"},
		{"Stundžia", 1, "Stundžienė"},
		{"Džiūve", 1, "Džiūvienė"},
		{"Šlymuo", 1, "Šlymienė"},
		{"Talat", 1, "Talatienė"},
		{"Kupcias", 1, "Kupcienė"},
		{"Šaltais", 1, "Šaltienė"},
		{"Tarlo", 1, "Tarlienė"},
		{"Taulai", 1, "Taulienė"},
		{"Kalnau", 1, "Kalnienė"},
		{"Kazlauskas", 2, "Kazlauskė"},
		{"Stankevičius", 2, "Stankevičė"},
		{"Butkus", 2, "Butkė"},
		{"Adomaitis", 2, "Adomaitė"},
		{"Kairys", 2, "Kairė"},
		{"Juška", 2, "Juškė"},
		{"Dirsė", 2, "Dirsė"},
		{"Stundžia", 2, "Stundžė"},
		{"Džiūve", 2, "Džiūvė"},
		{"Šlymuo", 2, "Šlymė"},
		{"Talat", 2, "Talatė"},
		{"Kupcias", 2, "Kupcė"},
		{"Šaltais", 2, "Šaltė"},
		{"Tarlo", 2, "Tarlė"},
		{"Taulai", 2, "Taulė"},
		{"Kalnau", 2, "Kalnė"},
	}
	for _, d := range tests {
		res := Feminize(d.val, d.maritalStatus)
		if res != d.exp {
			t.Errorf(
				"DATA: val: %v maritalStatus: %v EXPECTED: %v, GOT: %v",
				d.val,
				d.maritalStatus,
				d.exp,
				res,
			)
		}
	}
}
