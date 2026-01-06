package text

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	var res strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			res.WriteRune('_')
		}
		res.WriteRune(unicode.ToLower(r))
	}
	return res.String()
}

func TitleCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}
