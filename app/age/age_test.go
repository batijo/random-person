package age

import (
	"testing"
	"time"
)

func Test_GetAgeAt(t *testing.T) {
	tests := []struct {
		birthDate time.Time
		now       time.Time
		exp       int
	}{
		{GetDate(2000, 4, 6), GetDate(2010, 4, 6), 10},
		{GetDate(2001, 4, 6), GetDate(2009, 4, 6), 8},
		{GetDate(2004, 4, 6), GetDate(2005, 4, 4), 0},
	}
	for _, d := range tests {
		res := GetAgeAt(d.birthDate, d.now)
		if res != d.exp {
			t.Errorf(
				"DATA: birthDate: %v now: %v EXPECTED: %v, GOT: %v",
				d.birthDate,
				d.now,
				d.exp,
				res,
			)
		}
	}
}

func Test_isLeapYear(t *testing.T) {
	tests := []struct {
		date time.Time
		exp  bool
	}{
		{GetDate(2000, 4, 5), true},
		{GetDate(2016, 4, 5), true},
		{GetDate(2020, 4, 5), true},
		{GetDate(2024, 4, 5), true},
		{GetDate(2028, 4, 5), true},
		{GetDate(2001, 4, 5), false},
		{GetDate(1900, 4, 5), false},
		{GetDate(2100, 4, 5), false},
		{GetDate(2200, 4, 5), false},
		{GetDate(1999, 4, 5), false},
	}
	for _, d := range tests {
		res := isLeapYear(d.date.Year())
		if res != d.exp {
			t.Errorf(
				"DATA: %v EXPECTED: %v, GOT: %v",
				d.date,
				d.exp,
				res,
			)
		}
	}
}
