// sp stands for suffix parser
package sp

import "github.com/batijo/random-person/utils"

// Removes suffix from Lithuanian male s or any name
func RemoveSuffix(s string) string {
	if utils.LastElem(s) == "s" {
		if !utils.IsVowel(utils.StrElemEnd(s, 1)) {
			return s
		}
		s = utils.TrimLastElem(s)
		for utils.IsVowel(utils.LastElem(s)) {
			s = utils.TrimLastElem(s)
		}
	} else if utils.IsVowel(utils.LastElem(s)) {
		for utils.IsVowel(utils.LastElem(s)) {
			s = utils.TrimLastElem(s)
		}
	}
	return s
}
