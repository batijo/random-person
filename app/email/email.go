package email

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/batijo/random-person/app/models"
	"github.com/batijo/random-person/utils"
)

func New(person models.Person) string {

	return ""
}

func ParseWithTemplate(template string, person models.Person) (string, error) {
	if err := validateTemplate(template); err != nil {
		return "", err
	}
	return "", nil
}

// [fn] - inserts full persons name
// [fs] - inserts full persons surname
// [nws] - inserts name without suffix
// [sws] - inserts surname without suffix
// [by] - inserts birth year
// [pby] - inserts partial birth year (if year is 1985, inserts 85)
// [a] - inserts age
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
func validateTemplate(template string) error {
	commands := []string{"fn", "fs", "nws", "sws", "by", "pby", "a"}
	chars := strings.Split(template, "")
	i := 0
	eoc := true
	for i < len(chars) {
		eoc = true
		if chars[i] == "[" {
			eoc = false
			j := i + 1
			for j < len(chars) {
				if chars[j] == "]" {
					if chars[j-1] != "}" && !utils.ArrContains(commands, utils.ArrToString(chars[i+1:j])) {
						return fmt.Errorf("unknown command: %v", utils.ArrToString(chars[i+1:j]))
					}
					i = j + 1
					eoc = true
					break
				} else if chars[j] == "{" {
					if !utils.ArrContains(commands, utils.ArrToString(chars[i+1:j])) {
						return fmt.Errorf("unknown command: %v", utils.ArrToString(chars[i+1:j]))
					} else if len(chars) < j+6 {
						return fmt.Errorf("wrong sub command formating")
					}
					k := j + 1
					separator := 0
					for k < len(chars) {
						if chars[k] == "/" {
							separator = k
							if len(chars) < k+4 {
								return fmt.Errorf("wrong sub command formating")
							}
							subCommand := utils.ArrToString(chars[j+1 : k])
							if _, err := strconv.Atoi(subCommand); err != nil {
								switch len(subCommand) {
								case 1:
									if !utils.IsCharInString(subCommand, "ev") {
										return fmt.Errorf("unknown sub command: %v", subCommand)
									}
								case 2:
									if chars[j+1] != "e" {
										return fmt.Errorf("unknown sub command: %v", subCommand)
									} else if !(chars[j+2] == "v" || utils.IsCharInString(chars[j+2], "0123456789")) {
										return fmt.Errorf("unknown sub command: %v", subCommand)
									}
								default:
									if _, err := strconv.Atoi(utils.ArrToString(chars[j+2 : k])); !((chars[j+1] == "e" || chars[j+1] == "v") && err == nil) {
										return fmt.Errorf("unknown sub command: %v", subCommand)
									}
								}
							}
						} else if chars[k] == "}" {
							if _, err := strconv.Atoi(utils.ArrToString(chars[separator+1 : k])); err != nil {
								return fmt.Errorf("wrong sub command formating: %v", utils.ArrToString(chars[j+1:k]))
							}
							j = k
							break
						}
						k++
					}
				}
				j++
			}
		}
		if !eoc {
			return errors.New("command ending bracket is missing")
		}
		i++
	}
	return nil
}
