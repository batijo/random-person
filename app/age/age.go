package age

import (
	"math/rand"
	"time"

	"github.com/batijo/random-person/app/models"
)

// TODO: implement
func Random(p *models.Person) {
}

func GetDate(year, month, day int) time.Time {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

func GetRandomBirthDateByAgeRangeAt(fromAge, toAge int, now time.Time) time.Time {
	if fromAge > toAge || fromAge < 0 || toAge < 0 {
		return GetDate(0, 0, 0)
	}
	var (
		year, month, day int
		monthDays        = map[int]int{
			1:  31,
			2:  28,
			3:  31,
			4:  30,
			5:  31,
			6:  30,
			7:  31,
			8:  31,
			9:  30,
			10: 31,
			11: 30,
			12: 31,
		}
	)
	rand.Seed(time.Now().UnixNano())
	if fromAge == toAge {
		year = now.Year() - fromAge
	} else {
		year = rand.Intn(toAge-fromAge) + (now.Year() - toAge)
	}
	rand.Seed(time.Now().UnixNano())
	if now.Year()-fromAge == year {
		month = rand.Intn(12-int(now.Month())) + int(now.Month())
	} else {
		month = rand.Intn(11) + 1
	}
	rand.Seed(time.Now().UnixNano())
	if month == 2 && isLeapYear(year) {
		day = rand.Intn(29-1) + 1
	} else {
		day = rand.Intn(monthDays[month]-1) + 1
	}
	return GetDate(year, month, day)
}

func GetRandomBirthDateByAge(age int) time.Time {
	return GetRandomBirthDateByAgeRange(age, age)
}

func GetRandomBirthDateByAgeRange(fromAge, toAge int) time.Time {
	return GetRandomBirthDateByAgeRangeAt(fromAge, toAge, time.Now())
}

func GetAgeAt(birthDate, now time.Time) int {
	years := now.Year() - birthDate.Year()
	birthDay := getBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years -= 1
	}
	return years
}

func GetAge(birthDate time.Time) int {
	return GetAgeAt(birthDate, time.Now())
}

func getBirthDay(birthDate, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeapYear(birthDate.Year()) && !isLeapYear(now.Year()) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeapYear(now.Year()) && !isLeapYear(birthDate.Year()) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

func isLeapYear(year int) bool {
	// https://www.timeanddate.com/date/leapyear.html
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
