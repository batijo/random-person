// ssp stands for surname suffix parser
package ssp

import "github.com/batijo/random-person/utils"

// Feminize takes Lithuanian male surname and marital status uint
// (0 - unmarried, 1 - married, 2 - without marital status)
// Returns empty string if marital status is not in range (0-2)
// TODO: Some surnames are made with -(i)uvienė and -ūtė suffix but here never used.
func Feminize(s string, status uint) string {
	if status > 2 {
		return ""
	}
	surname := removeSuffix(s)
	switch status {
	case 0:
		if utils.LastElem(surname) == "k" && utils.StrElemEnd(surname, 2) == "t" {
			surname += "utė"
		} else if utils.IsCharInElements(utils.LastElem(surname), "čšž") {
			surname += "iūtė"
		} else if utils.IsVowel(utils.StrElemEnd(surname, 2)) {
			surname += "ytė"
		} else {
			surname += "aitė"
		}
	case 1:
		surname += "ienė"
	case 2:
		surname += "ė"
	}
	return surname
}

// Removes suffix from Lithuanian male surnames
func removeSuffix(surname string) string {
	if utils.LastElem(surname) == "s" {
		if !utils.IsVowel(utils.StrElemEnd(surname, 2)) {
			return surname
		}
		surname = utils.TrimLastElem(surname)
		for utils.IsVowel(utils.LastElem(surname)) {
			surname = utils.TrimLastElem(surname)
		}
	} else if utils.IsVowel(utils.LastElem(surname)) {
		for utils.IsVowel(utils.LastElem(surname)) {
			surname = utils.TrimLastElem(surname)
		}
	}
	return surname
}
