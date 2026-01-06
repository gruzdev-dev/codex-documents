package tests

import (
	"fmt"
	"regexp"
	"strings"
)

func checkValidateMethodExists(code, structName string) bool {
	pattern := `func\s+\(r\s+\*` + regexp.QuoteMeta(structName) + `\)\s+Validate\(\)\s+error`
	matched, _ := regexp.MatchString(pattern, code)
	return matched
}

func checkRequiredFieldValidation(code, fieldName string) bool {
	patterns := []string{
		`if\s+r\.` + regexp.QuoteMeta(fieldName) + `\s+==\s+nil`,
		`if\s+r\.` + regexp.QuoteMeta(fieldName) + `\s+==\s+""`,
		`if\s+r\.` + regexp.QuoteMeta(fieldName) + `\s+==\s+0`,
		`if\s+len\(r\.` + regexp.QuoteMeta(fieldName) + `\)\s+<\s+\d+`,
	}
	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, code)
		if matched {
			return true
		}
	}
	return false
}

func checkMaxLengthValidation(code, fieldName string, maxLength int) bool {
	if strings.Contains(code, fmt.Sprintf("exceeds maxLength %d", maxLength)) {
		return true
	}
	pattern := `len\(.*r\.` + regexp.QuoteMeta(fieldName) + `.*\)\s+>\s+` + fmt.Sprintf("%d", maxLength)
	matched, _ := regexp.MatchString(pattern, code)
	return matched
}

func checkPatternValidation(code, fieldName string) bool {
	return strings.Contains(code, "regexp.MatchString") &&
		strings.Contains(code, fmt.Sprintf("field '%s' does not match pattern", fieldName))
}

func checkNestedValidation(code, fieldName string) bool {
	pattern := `r\.` + regexp.QuoteMeta(fieldName) + `\.Validate\(\)`
	matched, _ := regexp.MatchString(pattern, code)
	return matched
}

func checkImports(code string) (hasFmt, hasRegexp bool) {
	hasFmt = strings.Contains(code, `import`) &&
		(strings.Contains(code, `"fmt"`) || strings.Contains(code, `fmt`))
	hasRegexp = strings.Contains(code, `import`) &&
		(strings.Contains(code, `"regexp"`) || strings.Contains(code, `regexp`))
	return
}
