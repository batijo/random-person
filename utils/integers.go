package utils

import (
	"strconv"
	"strings"
)

// StringContainsInt checks if string contains integer i.
// Elements needs to be separated with space
func StringContainsInt(s string, i int) bool {
	for _, v := range strings.Fields(s) {
		if v == strconv.Itoa(i) {
			return true
		}
	}
	return false
}
