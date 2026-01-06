package text

import "unicode"

func IsValidGoIdentifier(s string) bool {
	if s == "" {
		return false
	}
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	first := runes[0]
	if !unicode.IsLetter(first) && first != '_' {
		return false
	}
	for i := 1; i < len(runes); i++ {
		r := runes[i]
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}
