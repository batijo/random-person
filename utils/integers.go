package utils

import (
	"strconv"
	"strings"
)

// StringContainsInt checks if string contains integer i.
// Elements needs to be separated with space
func StringContainsInt(s string, i int) bool {
	s = strings.ReplaceAll(s, "\t", " ")
	arr := strings.Split(s, " ")
	for _, a := range arr {
		if num, err := strconv.Atoi(strings.TrimSpace(a)); err == nil {
			if num == i {
				return true
			}
		} else {
			continue
		}
	}
	return false
}
