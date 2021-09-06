package email

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/app/sp"
	"github.com/batijo/random-person/utils"
)

type Email struct {
	*models.Person
}

func New(person models.Person) Email {
	return Email{Person: &person}
}

func (e *Email) Random() string {
	return ""
}

// [fn] - inserts full persons name
// [fs] - inserts full persons surname
// [nws] - inserts name without suffix
// [sws] - inserts surname without suffix
// [by] - inserts birth year
// [pby] - inserts partial birth year (if year is 1985, inserts 85)
// [command{3/2}] - command can be any command , number 3 represents which element, 2 how many time multiply it
// e.g. Surname is Kazlauskas so [sws{4/3}] is Kazlllausk
// number 3 can be replaced with e for last letter
// e.g. Surname is Kazlauskas so [fs{e/3}] is Kazlauskasss
// if you add number x after e it will multiply x letter from end
// e.g. Surname is Kazlauskas so [sws{e3/4}] is Kazlauuuusk
// if you add v (vowel) then [command{v/2}] multiplies first vowel 2 times
// e.g. Surname is Kazlauskas so [sws{v/3}] is Kaaazlausk
// You can also add e before v
// e.g. Surname is Kazlauskas so [fs{ev/4}] is Kaazlauskaaaas
// everything what goes after @ symbol is added without checking
// if you don't add @ a random popular domain will be added
// e.g. @gmail.com @yahoo.com @outlook.com ...
func (e *Email) ParseWithTemplate(template string) {
	var (
		chars       = strings.Split(template, "")
		startOfCopy = 0
	)
	for i := 0; i < len(chars); i++ {
		if chars[i] == "[" {
			e.Email += utils.ArrToString(chars[startOfCopy:i])
			for j := i + 1; j < len(chars); j++ {
				if chars[j] == "]" {
					e.Email += e.parseTemplateCommand(utils.ArrToString(chars[i+1 : j]))
					startOfCopy = j + 1
					i = j
					break
				}
			}
		} else if chars[i] == "@" {
			e.Email += utils.ArrToString(chars[startOfCopy:i])
			e.Email += utils.ArrToString(chars[i:])
			break
		} else if len(chars) == i+1 && !(chars[i+1] == "@" || chars[i+1] == "]") {
			e.Email += utils.ArrToString(chars[startOfCopy:])
			// TODO: random email domain
			break
		}
	}
}

func (e *Email) parseTemplateCommand(command string) string {
	var (
		word        string
		newWord     = ""
		chars       = strings.Split(command, "")
		subCommands []string
	)
	for i := 0; i < len(chars); i++ {
		if chars[i] == "{" {
			word = e.getByCommand(utils.ArrToString(chars[:i]))
			chars = chars[i:]
			break
		} else if len(chars) == i+1 {
			word = e.getByCommand(utils.ArrToString(chars))
			chars = chars[i+1:]
			break
		}
	}
	for i := 0; i < len(chars); i++ {
		if chars[i] == "{" {
			for j := i + 1; j < len(chars); j++ {
				if chars[j] == "}" {
					subCommands = append(subCommands, utils.ArrToString(chars[i+1:j]))
					i = j
					break
				}
			}
		}
	}
	charMap := parseTemplateSubCommands(subCommands, word)
	for i, s := range []rune(word) {
		if charMap[i] != 0 {
			for j := 0; j < charMap[i]; j++ {
				newWord += string(s)
			}
		} else {
			newWord += string(s)
		}
	}
	return newWord
}

func (e *Email) getByCommand(command string) string {
	switch command {
	case "fn":
		return e.Name
	case "fs":
		return e.Surname
	case "nws":
		return sp.RemoveSuffix(e.Name)
	case "sws":
		return sp.RemoveSuffix(e.Surname)
	case "by":
		return fmt.Sprint(e.BirthDate.Year())
	case "pby":
		return utils.Trim(fmt.Sprint(e.BirthDate.Year()), 2, false)
	}
	return "boi"
}

func parseTemplateSubCommands(subCommans []string, word string) map[int]int {
	var charMap = make(map[int]int)
	if len(subCommans) < 1 {
		return charMap
	}
	for _, fsc := range subCommans {
		ssc := strings.Split(fsc, "/")
		if len(ssc) != 2 {
			continue
		}
		subCommand := ssc[0]
		iterations, err := strconv.Atoi(ssc[1])
		if err != nil {
			continue
		}
		if position, err := strconv.Atoi(subCommand); err == nil {
			if position >= len([]rune(word)) {
				continue
			}
			charMap[position] = iterations
		} else if char := utils.StrElem(subCommand, 0); char == "e" {
			if len([]rune(subCommand)) > 1 {
				elem := utils.Trim(subCommand, 1, false)
				if elem != "v" {
					position, err := strconv.Atoi(elem)
					if err != nil || position > len([]rune(word)) {
						continue
					}
					charMap[len([]rune(word))-position] = iterations
				} else {
					for i := len([]rune(word)); i >= 0; i-- {
						if utils.IsVowel(utils.StrElem(word, i)) {
							charMap[i] = iterations
							break
						}
					}
				}
			} else {
				charMap[len([]rune(word))-1] = iterations
			}
		} else if subCommand == "v" {
			for i := 0; i < len([]rune(word)); i++ {
				if utils.IsVowel(utils.StrElem(word, i)) {
					charMap[i] = iterations
					break
				}
			}
		}
	}
	return charMap
}
