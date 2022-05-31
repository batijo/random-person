package email

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/app/sp"
	"github.com/batijo/random-person/utils"
)

func Random(p *models.Person) {
	ParseWithTemplate(getRandomTemplate(), p)
	p.Email = strings.ToLower(p.Email)
	p.StringifyBirthDate()
}

// [fn] - inserts full persons name
// [fs] - inserts full persons surname
// [nws] - inserts name without suffix
// [sws] - inserts surname without suffix
// [by] - inserts birth year
// [pby] - inserts partial birth year (if year is 1985, inserts 85)
// if you add a number N after any command it will take N number of characters from the start of a result
// e.g. Name is Jonas so [fn2] is Jo
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
func ParseWithTemplate(template string, p *models.Person) {
	var (
		chars       = strings.Split(template, "")
		startOfCopy = 0
	)
	for i := 0; i < len(chars); i++ {
		if chars[i] == "[" {
			p.Email += utils.ArrToString(chars[startOfCopy:i])
			for j := i + 1; j < len(chars); j++ {
				if chars[j] == "]" {
					p.Email += parseTemplateCommand(utils.ArrToString(chars[i+1:j]), p)
					startOfCopy = j + 1
					i = j
					break
				}
			}
		} else if chars[i] == "@" {
			p.Email += utils.ArrToString(chars[startOfCopy:i])
			p.Email += utils.ArrToString(chars[i:])
			break
		}
		if len(chars) == i+1 {
			p.Email += utils.ArrToString(chars[startOfCopy:])
			p.Email += "@" + getRandomDomain()
			break
		}
	}
}

func parseTemplateCommand(command string, p *models.Person) string {
	var (
		word        string
		newWord     = ""
		chars       = strings.Split(command, "")
		subCommands []string
	)
	for i := 0; i < len(chars); i++ {
		if chars[i] == "{" {
			word = getByCommand(utils.ArrToString(chars[:i]), p)
			chars = chars[i:]
			break
		} else if len(chars) == i+1 {
			word = getByCommand(utils.ArrToString(chars), p)
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

func getByCommand(command string, p *models.Person) string {
	var str string
	switch utils.FilterNumbers(command) {
	case "fn":
		str = p.Name
	case "fs":
		str = p.Surname
	case "nws":
		str = sp.RemoveSuffix(p.Name)
	case "sws":
		str = sp.RemoveSuffix(p.Surname)
	case "by":
		str = fmt.Sprint(p.BirthDate.Year())
	case "pby":
		str = utils.Trim(fmt.Sprint(p.BirthDate.Year()), 2, false)
	}
	if sn := utils.FilterLetters(command); sn != "" {
		n, err := strconv.Atoi(sn)
		if err != nil {
			return str
		} else {
			return utils.TrimUntil(str, n, true)
		}
	}
	return str
}

// TODO: refactor
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

func getRandomDomain() string {
	rand.Seed(time.Now().UnixNano())
	return domains.Pick().(string)
}

func getRandomTemplate() string {
	rand.Seed(time.Now().UnixNano())
	return templates.Pick().(string)
}
