// sp stands for suffix parser
package sp

import "github.com/batijo/random-person/utils"

// FeminizeMaleSurname takes Lithuanian male surname and marital status uint8
// (0 - unmarried, 1 - married, 2 - without marital status)
// Returns empty string if marital status is not in range (0-2)
// TODO: Some surnames are made with -(i)uvienė and -ūtė suffix but here never used.
func FeminizeMaleSurname(s string, status uint8) string {
	if status > 2 {
		return ""
	}
	surname := removeSuffix(s)
	switch status {
	case 0:
		if utils.LastElem(surname) == "k" && utils.StrElemEnd(surname, 2) == "t" {
			surname += "utė"
		} else if utils.IsCharInString(utils.LastElem(surname), "čšž") {
			surname += "iūtė"
		} else if utils.IsVowel(utils.StrElemEnd(surname, 2)) && !utils.IsCharInString(utils.LastElem(surname), "mn") {
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
