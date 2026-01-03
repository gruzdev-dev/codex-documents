package generator

import (
	"strings"
	"unicode"
)

var usedTypes = make(map[string]bool)

func (g *Generator) mapGoType(el ElementDefinition) string {
	if len(el.Type) > 1 {
		return "any"
	}

	if len(el.Type) == 0 {
		typeName := g.deriveNestedTypeName(el.Path)
		if typeName == "" {
			return "any"
		}
		return typeName
	}

	fhirType := el.Type[0].Code

	switch fhirType {
	case "string", "uri", "code", "id", "markdown", "base64Binary", "oid", "canonical", "url", "xhtml":
		return "string"
	case "boolean":
		return "bool"
	case "integer", "unsignedInt", "positiveInt":
		return "int"
	case "integer64":
		return "int64"
	case "decimal":
		return "float64"
	case "dateTime", "date", "instant", "time":
		return "string"
	case "Resource", "ResourceList":
		return "json.RawMessage"
	case "BackboneElement", "Element":
		return g.deriveNestedTypeName(el.Path)
	default:
		if strings.Contains(fhirType, "://") || strings.Contains(fhirType, "/") {
			parts := strings.FieldsFunc(fhirType, func(r rune) bool {
				return r == '/' || r == '.'
			})
			if len(parts) > 0 {
				lastPart := parts[len(parts)-1]
				switch lastPart {
				case "String":
					return "string"
				case "Boolean":
					return "bool"
				case "Integer", "Integer64":
					return "int"
				case "Decimal":
					return "float64"
				case "Date", "DateTime", "Time":
					return "string"
				default:
					if isValidGoIdentifier(lastPart) {
						usedTypes[lastPart] = true
						return lastPart
					}
				}
			}
			return "any"
		}

		if !isValidGoIdentifier(fhirType) {
			return "any"
		}

		usedTypes[fhirType] = true
		return fhirType
	}
}

func (g *Generator) deriveNestedTypeName(path string) string {
	if path == "" {
		return ""
	}
	parts := strings.Split(path, ".")
	for i := range parts {
		// handle value[x]
		parts[i] = strings.ReplaceAll(parts[i], "[x]", "")
		parts[i] = titleCase(parts[i])
	}
	return strings.Join(parts, "")
}

func titleCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func isValidGoIdentifier(s string) bool {
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
