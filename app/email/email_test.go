package email

import (
	"reflect"
	"testing"

	"github.com/batijo/random-person/app/age"
	"github.com/batijo/random-person/app/models"
	wr "github.com/mroth/weightedrand"
)

func Test_ParseWithTemplate(t *testing.T) {
	tests := []struct {
		person   models.Person
		template string
		exp      string
	}{ // Birth year 2003
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[fn].[fs]@gmail.com",
			"Jonas.Kazlauskas@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[fn1].[fs2]@gmail.com",
			"J.Ka@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws].[sws]@gmail.com",
			"Jon.Kazlausk@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws].[sws].[by]@gmail.com",
			"Jon.Kazlausk.2003@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws].[sws].[pby]@gmail.com",
			"Jon.Kazlausk.03@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws].[sws{e/3}]@gmail.com",
			"Jon.Kazlauskkk@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws{ev/4}{e/3}].[sws]@gmail.com",
			"Joooonnn.Kazlausk@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws{4/3}].[sws{4/3}]@gmail.com",
			"Jon.Kazlaaausk@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws{}].[sws{}]@gmail.com",
			"Jon.Kazlausk@gmail.com",
		},
		{
			newTestPerson("Jonas", "Kazlauskas", 18, 18, 0),
			"[nws2{}].[sws3{}]@gmail.com",
			"Jo.Kaz@gmail.com",
		},
		{
			newTestPerson("ąčęėįš", "žūųšįėęč", 18, 18, 0),
			"[fn{ev/2}{e/3}].[fs{v/2}]@gmail.com",
			"ąčęėįįššš.žūūųšįėęč@gmail.com",
		},
	}
	mockDomainData()
	for _, d := range tests {
		ParseWithTemplate(d.template, &d.person)
		if d.exp != d.person.Email {
			t.Errorf(
				"DATA: template: %v EXPECTED: %v, GOT: %v",
				d.template,
				d.exp,
				d.person.Email,
			)
		}
	}
}

func Test_parseTemplateSubCommands(t *testing.T) {
	tests := []struct {
		subCommands []string
		word        string
		exp         map[int]int
	}{
		{[]string{"0/2", "2/2"}, "Kazlauskas", map[int]int{0: 2, 2: 2}},
		{[]string{"e/2", "v/2"}, "Kazlauskas", map[int]int{9: 2, 1: 2}},
		{[]string{"ev/2", "e3/2"}, "Kazlauskas", map[int]int{8: 2, 7: 2}},
		{[]string{"0/2", "2/2"}, "ąčęėįšųūž", map[int]int{0: 2, 2: 2}},
		{[]string{"e/2", "v/2"}, "ąčęėįšųūž", map[int]int{8: 2, 0: 2}},
		{[]string{"ev/2", "e3/2"}, "ąčęėįšųūž", map[int]int{7: 2, 6: 2}},
		{[]string{"ev/2"}, "ą", map[int]int{0: 2}},
		{[]string{"e/2"}, "ą", map[int]int{0: 2}},
		{[]string{"0/2"}, "ą", map[int]int{0: 2}},
		{[]string{"v/2"}, "ą", map[int]int{0: 2}},
		{[]string{"1/2"}, "ą", map[int]int{}},
		{[]string{""}, "ąsd", map[int]int{}},
		{[]string{"d/d"}, "ąsd", map[int]int{}},
		{[]string{"/"}, "ąsd", map[int]int{}},
		{[]string{"/2"}, "ąsd", map[int]int{}},
		{[]string{"v/2"}, "d", map[int]int{}},
		{[]string{"e4/"}, "ąsd", map[int]int{}},
		{[]string{"ev/2", "e15/2"}, "ąčęėįšųūž", map[int]int{7: 2}},
		{[]string{"15/2", "e3/2"}, "ąčęėįšųūž", map[int]int{6: 2}},
		{[]string{"ev/2", "e3/2", "19/2"}, "ąčęėįšųūž", map[int]int{7: 2, 6: 2}},
	}
	for _, d := range tests {
		res := parseTemplateSubCommands(d.subCommands, d.word)
		if !reflect.DeepEqual(res, d.exp) {
			t.Errorf(
				"DATA: subCommands: %v word: %v EXPECTED: %v, GOT: %v",
				d.subCommands,
				d.word,
				d.exp,
				res,
			)
		}
	}
}

func newTestPerson(nm, srnm string, agf, agl, gender int) models.Person {
	return models.Person{
		NameOnly:    &models.NameOnly{Name: nm},
		SurnameOnly: &models.SurnameOnly{Surname: srnm},
		BirthDate:   age.GetRandomBirthDateByAgeRangeAt(agf, agl, age.GetBirthDate(2021, 9, 3)),
		Gender:      uint(gender),
	}
}

func mockDomainData() {
	var wd = []weightData{
		{
			"gmail.com",
			10,
		},
		{
			"yahoo.com",
			8,
		},
	}
	domains, _ = wr.NewChooser(getChoices(wd)...)
}
