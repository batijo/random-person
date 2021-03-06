package utils

import (
	"strings"
	"unicode/utf8"
)

// IsVowel returns true if character is Lithuanian vowel
func IsVowel(s string) bool {
	return IsCharInString(s, "aeiouyąęėįųū")
}

// IsCharInString receives single character and a string of characters
// and checks if those characters contains this element
// regardless of whether the letters are uppercase or lowercase
func IsCharInString(s, elements string) bool {
	s = strings.ToLower(s)
	elements = strings.ToLower(elements)
	r := []rune(s)
	if len(r) != 1 {
		return false
	}
	return strings.ContainsRune(elements, r[0])
}

// LastElemN return N elements from the end of a string
func LastElemN(s string, n int) string {
	return Trim(s, len([]rune(s))-n, false)
}

// LastElem returns last character of a string
func LastElem(s string) string {
	return StrElem(s, utf8.RuneCountInString(s)-1)
}

// StrElemEnd returns character at given position from end, same as strElem just inverted.
// position represent which element from end to return. 0 would be last element
func StrElemEnd(s string, position int) string {
	return StrElem(s, utf8.RuneCountInString(s)-position-1)
}

// StrElem returns one character of string at give position.
// It returns empty string if element you want to access is out of range
func StrElem(s string, position int) string {
	if position < 0 {
		return ""
	}
	r := []rune(s)
	if position >= len(r) {
		return ""
	}
	return string(r[position])
}

// TrimLastElem trims last character from end of a string
func TrimLastElem(s string) string {
	return Trim(s, 1, true)
}

// Trim funcion trims characters from end or from beginning.
// elemCount - how many characters to trim, if right equals to true it trims from end
func Trim(s string, elemCount int, right bool) string {
	if elemCount < 0 {
		return s
	} else if elemCount > utf8.RuneCountInString(s) {
		return ""
	}
	if right {
		for i := 0; i < elemCount; i++ {
			_, size := utf8.DecodeLastRuneInString(s)
			s = s[:len(s)-size]
		}
		return s
	}
	for i := 0; i < elemCount; i++ {
		_, size := utf8.DecodeRuneInString(s)
		s = s[size:]
	}
	return s
}

// TrimUntil trims characters until n left
func TrimUntil(s string, n int, right bool) string {
	if len([]rune(s)) <= n {
		return s
	}
	return Trim(s, len([]rune(s))-n, right)
}

// FilterChars removes all characters from string which are in chars string
func FilterChars(s, chars string) string {
	for _, c := range chars {
		s = strings.Replace(s, string(c), "", -1)
	}
	return s
}

// FilterNumbers removes all numbers from string
func FilterNumbers(s string) string {
	return FilterChars(s, "0123456789")
}

// FilterLetters removes all letters from string
func FilterLetters(s string) string {
	return FilterChars(s, "ąčęėįšųūžqwertyuiopasdfghjklzxcvbnm")
}

func ArrContains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func ArrToString(arr []string) string {
	s := ""
	for _, e := range arr {
		s += e
	}
	return s
}
