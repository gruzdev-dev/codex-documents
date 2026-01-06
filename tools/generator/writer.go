package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func (g *Generator) WriteResource(def StructureDefinition) error {
	var buf bytes.Buffer

	structMap := g.ProcessElements(def.Name, def.Snapshot.Element)

	needsJSON := false
	for _, fields := range structMap {
		for _, f := range fields {
			if strings.Contains(f.GoType, "json.RawMessage") {
				needsJSON = true
				break
			}
		}
		if needsJSON {
			break
		}
	}

	fmt.Fprintf(&buf, "package models\n\n")
	if needsJSON {
		fmt.Fprintf(&buf, "import \"encoding/json\"\n\n")
	}

	usedTypesInFile := make(map[string]bool)
	for _, fields := range structMap {
		for _, f := range fields {
			baseType := extractBaseType(f.GoType)
			if baseType != "" && !isBuiltinType(baseType) {
				usedTypesInFile[baseType] = true
			}
		}
	}

	for usedType := range usedTypesInFile {
		if _, exists := structMap[usedType]; !exists {
			// Don't create empty structures for types from Definitions
			// They will be generated as separate files through generator.go:75-84
			// Only create empty structures for types that are not in Definitions (unknown types)
			if _, defined := g.Definitions[usedType]; !defined {
				structMap[usedType] = []FieldInfo{}
			}
		}
	}

	g.writeStruct(&buf, def.Name, def.Description, structMap[def.Name])

	for sName, fields := range structMap {
		if sName == def.Name {
			continue
		}
		g.writeStruct(&buf, sName, "", fields)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		_ = os.WriteFile("debug_failed.go", buf.Bytes(), 0644)
		lineNum := extractLineNumber(err.Error())
		return fmt.Errorf("format error for %s at line %s: %w. Check debug_failed.go", def.Name, lineNum, err)
	}

	fileName := toSnakeCase(def.Name) + ".go"
	return os.WriteFile(filepath.Join(g.OutputPath, fileName), formatted, 0644)
}

func (g *Generator) writeStruct(buf *bytes.Buffer, name, comment string, fields []FieldInfo) {
	// Пропускаем пустые структуры, если они не определены в Definitions
	// (типы из Definitions должны генерироваться даже если пустые, так как они могут быть зависимостями)
	if len(fields) == 0 {
		if _, defined := g.Definitions[name]; !defined {
			return
		}
		// Для типов из Definitions генерируем пустую структуру с комментарием
		if comment == "" {
			comment = "Empty structure for " + name
		}
	}

	if comment != "" {
		fmt.Fprintf(buf, "// %s\n", sanitizeComment(comment))
	}
	fmt.Fprintf(buf, "type %s struct {\n", name)
	for _, f := range fields {
		goType := f.GoType
		if after, ok := strings.CutPrefix(goType, "[]"); ok {
			baseType := after
			if !isValidGoIdentifier(baseType) && baseType != "any" && baseType != "json.RawMessage" {
				goType = "[]any"
			}
		} else if strings.HasPrefix(goType, "*") {
			baseType := strings.TrimPrefix(goType, "*")
			if !isValidGoIdentifier(baseType) && baseType != "any" && baseType != "json.RawMessage" {
				goType = "*any"
			}
		} else {
			if !isValidGoIdentifier(goType) && goType != "any" && goType != "bool" && goType != "json.RawMessage" {
				goType = "any"
			}
		}

		commentPart := ""
		if f.Comment != "" {
			commentPart = " // " + sanitizeComment(f.Comment)
		}
		fmt.Fprintf(buf, "\t%s %s %s%s\n", f.Name, goType, f.JSONTag, commentPart)
	}
	fmt.Fprintf(buf, "}\n\n")
}

func sanitizeComment(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "`", "'")
	return strings.TrimSpace(s)
}

func extractLineNumber(errMsg string) string {
	re := regexp.MustCompile(`:(\d+):\d+:`)
	matches := re.FindStringSubmatch(errMsg)
	if len(matches) > 1 {
		return matches[1]
	}
	return "unknown"
}

func toSnakeCase(s string) string {
	var res strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			res.WriteRune('_')
		}
		res.WriteRune(unicode.ToLower(r))
	}
	return res.String()
}

func extractBaseType(goType string) string {
	res := goType
	for {
		var ok bool
		if res, ok = strings.CutPrefix(res, "[]"); ok {
			continue
		}
		if res, ok = strings.CutPrefix(res, "*"); ok {
			continue
		}
		break
	}
	return res
}

func isBuiltinType(t string) bool {
	switch t {
	case "string", "bool", "int", "int64", "float64", "any", "interface{}", "byte", "rune", "json.RawMessage":
		return true
	}
	return false
}
