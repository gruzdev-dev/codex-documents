package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"codex-documents/tools/text"
)

func (g *Generator) WriteResource(def StructureDefinition) error {
	var buf bytes.Buffer

	structMap := g.ProcessElements(def.Name, def.Snapshot.Element)

	needsJSON := false
	needsRegexp := false
	needsFmt := g.needsFmt(structMap)

	for _, fields := range structMap {
		for _, f := range fields {
			if strings.Contains(f.GoType, "json.RawMessage") {
				needsJSON = true
			}
			if f.Pattern != "" {
				needsRegexp = true
			}
		}
	}

	fmt.Fprintf(&buf, "package models\n\n")
	var imports []string
	if needsJSON {
		imports = append(imports, "\"encoding/json\"")
	}
	if needsRegexp {
		imports = append(imports, "\"regexp\"")
	}
	if needsFmt {
		imports = append(imports, "\"fmt\"")
	}
	if len(imports) > 0 {
		fmt.Fprintf(&buf, "import (\n")
		for _, imp := range imports {
			fmt.Fprintf(&buf, "\t%s\n", imp)
		}
		fmt.Fprintf(&buf, ")\n\n")
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
			if _, defined := g.Definitions[usedType]; !defined {
				structMap[usedType] = []FieldInfo{}
			}
		}
	}

	g.writeStruct(&buf, def.Name, def.Description, structMap[def.Name])
	g.writeValidateMethod(&buf, def.Name, structMap[def.Name], structMap)

	for sName, fields := range structMap {
		if sName == def.Name {
			continue
		}
		g.writeStruct(&buf, sName, "", fields)
		g.writeValidateMethod(&buf, sName, fields, structMap)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		_ = os.WriteFile("debug_failed.go", buf.Bytes(), 0644)
		lineNum := extractLineNumber(err.Error())
		return fmt.Errorf("format error for %s at line %s: %w. Check debug_failed.go", def.Name, lineNum, err)
	}

	fileName := text.ToSnakeCase(def.Name) + ".go"
	return os.WriteFile(filepath.Join(g.OutputPath, fileName), formatted, 0644)
}

func (g *Generator) writeStruct(buf *bytes.Buffer, name, comment string, fields []FieldInfo) {
	if len(fields) == 0 {
		if _, defined := g.Definitions[name]; !defined {
			return
		}
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
			if !text.IsValidGoIdentifier(baseType) && baseType != "any" && baseType != "json.RawMessage" {
				goType = "[]any"
			}
		} else if strings.HasPrefix(goType, "*") {
			baseType := strings.TrimPrefix(goType, "*")
			if !text.IsValidGoIdentifier(baseType) && baseType != "any" && baseType != "json.RawMessage" {
				goType = "*any"
			}
		} else {
			if !text.IsValidGoIdentifier(goType) && goType != "any" && goType != "bool" && goType != "json.RawMessage" {
				goType = "any"
			}
		}

		commentPart := ""
		if f.Comment != "" {
			commentPart = " // " + sanitizeComment(f.Comment)
		}

		jsonTagValue := strings.TrimPrefix(strings.TrimSuffix(f.JSONTag, "`"), "`json:")
		jsonTagValue = strings.Trim(jsonTagValue, "\"")

		var tagParts []string
		tagParts = append(tagParts, fmt.Sprintf("json:\"%s\"", jsonTagValue))

		if f.BSONTag != "" {
			bsonTagValue := strings.TrimPrefix(strings.TrimSuffix(f.BSONTag, "`"), "`bson:")
			bsonTagValue = strings.Trim(bsonTagValue, "\"")
			tagParts = append(tagParts, fmt.Sprintf("bson:\"%s\"", bsonTagValue))
		}

		tags := "`" + strings.Join(tagParts, " ") + "`"
		fmt.Fprintf(buf, "\t%s %s %s%s\n", f.Name, goType, tags, commentPart)
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

func (g *Generator) needsFmt(structMap map[string][]FieldInfo) bool {
	for _, fields := range structMap {
		if len(fields) == 0 {
			continue
		}

		for _, f := range fields {
			baseType := extractBaseType(f.GoType)
			isArray := strings.HasPrefix(f.GoType, "[]")
			isPointer := strings.HasPrefix(f.GoType, "*")
			isBuiltin := isBuiltinType(baseType)

			if baseType == "json.RawMessage" || baseType == "any" {
				continue
			}

			if f.IsRequired {
				return true
			}

			if f.MaxLength != nil {
				return true
			}

			if f.Pattern != "" {
				return true
			}

			if f.Fixed != nil {
				return true
			}

			if !isBuiltin && !isArray && isPointer {
				if _, exists := structMap[baseType]; exists {
					return true
				}
				if def, ok := g.Definitions[baseType]; ok && def.Name != "" {
					return true
				}
			} else if !isBuiltin && isArray {
				if _, exists := structMap[baseType]; exists {
					return true
				}
				if def, ok := g.Definitions[baseType]; ok && def.Name != "" {
					return true
				}
			}
		}
	}
	return false
}

func (g *Generator) writeValidateMethod(buf *bytes.Buffer, structName string, fields []FieldInfo, structMap map[string][]FieldInfo) {
	fmt.Fprintf(buf, "func (r *%s) Validate() error {\n", structName)

	if len(fields) == 0 {
		fmt.Fprintf(buf, "\treturn nil\n")
		fmt.Fprintf(buf, "}\n\n")
		return
	}

	for _, f := range fields {
		baseType := extractBaseType(f.GoType)
		isArray := strings.HasPrefix(f.GoType, "[]")
		isPointer := strings.HasPrefix(f.GoType, "*")
		isBuiltin := isBuiltinType(baseType)

		if baseType == "json.RawMessage" || baseType == "any" {
			continue
		}

		if f.IsRequired {
			if isArray {
				fmt.Fprintf(buf, "\tif len(r.%s) < %d {\n", f.Name, f.Min)
				fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' must have at least %d elements\")\n", f.Name, f.Min)
				fmt.Fprintf(buf, "\t}\n")
			} else if isPointer {
				fmt.Fprintf(buf, "\tif r.%s == nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' is required\")\n", f.Name)
				fmt.Fprintf(buf, "\t}\n")
			} else if isBuiltin {
				if baseType == "string" {
					fmt.Fprintf(buf, "\tif r.%s == \"\" {\n", f.Name)
					fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' is required\")\n", f.Name)
					fmt.Fprintf(buf, "\t}\n")
				} else if baseType == "bool" {
				} else {
					fmt.Fprintf(buf, "\tif r.%s == 0 {\n", f.Name)
					fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' is required\")\n", f.Name)
					fmt.Fprintf(buf, "\t}\n")
				}
			}
		}

		if f.MaxLength != nil && (baseType == "string" || (isPointer && baseType == "string")) {
			if isPointer {
				fmt.Fprintf(buf, "\tif r.%s != nil && len(*r.%s) > %d {\n", f.Name, f.Name, *f.MaxLength)
				fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' exceeds maxLength %d\")\n", f.Name, *f.MaxLength)
				fmt.Fprintf(buf, "\t}\n")
			} else {
				fmt.Fprintf(buf, "\tif len(r.%s) > %d {\n", f.Name, *f.MaxLength)
				fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' exceeds maxLength %d\")\n", f.Name, *f.MaxLength)
				fmt.Fprintf(buf, "\t}\n")
			}
		}

		if f.Pattern != "" && (baseType == "string" || (isPointer && baseType == "string")) {
			if isPointer {
				fmt.Fprintf(buf, "\tif r.%s != nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\tmatched, _ := regexp.MatchString(`%s`, *r.%s)\n", f.Pattern, f.Name)
				fmt.Fprintf(buf, "\t\tif !matched {\n")
				fmt.Fprintf(buf, "\t\t\treturn fmt.Errorf(\"field '%s' does not match pattern '%s'\")\n", f.Name, f.Pattern)
				fmt.Fprintf(buf, "\t\t}\n")
				fmt.Fprintf(buf, "\t}\n")
			} else {
				fmt.Fprintf(buf, "\tmatched, _ := regexp.MatchString(`%s`, r.%s)\n", f.Pattern, f.Name)
				fmt.Fprintf(buf, "\tif !matched {\n")
				fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' does not match pattern '%s'\")\n", f.Name, f.Pattern)
				fmt.Fprintf(buf, "\t}\n")
			}
		}

		if f.Fixed != nil {
			if isPointer {
				fmt.Fprintf(buf, "\tif r.%s != nil {\n", f.Name)
				g.writeFixedValidation(buf, f.Name, f.Fixed, baseType, true)
				fmt.Fprintf(buf, "\t}\n")
			} else {
				g.writeFixedValidation(buf, f.Name, f.Fixed, baseType, false)
			}
		}

		if !isBuiltin && !isArray && isPointer {
			if _, exists := structMap[baseType]; exists {
				fmt.Fprintf(buf, "\tif r.%s != nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\tif err := r.%s.Validate(); err != nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\t\treturn fmt.Errorf(\"%s: %%w\", err)\n", f.Name)
				fmt.Fprintf(buf, "\t\t}\n")
				fmt.Fprintf(buf, "\t}\n")
			} else if def, ok := g.Definitions[baseType]; ok && def.Name != "" {
				fmt.Fprintf(buf, "\tif r.%s != nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\tif err := r.%s.Validate(); err != nil {\n", f.Name)
				fmt.Fprintf(buf, "\t\t\treturn fmt.Errorf(\"%s: %%w\", err)\n", f.Name)
				fmt.Fprintf(buf, "\t\t}\n")
				fmt.Fprintf(buf, "\t}\n")
			}
		} else if !isBuiltin && isArray {
			if _, exists := structMap[baseType]; exists {
				fmt.Fprintf(buf, "\tfor i, item := range r.%s {\n", f.Name)
				fmt.Fprintf(buf, "\t\tif err := item.Validate(); err != nil {\n")
				fmt.Fprintf(buf, "\t\t\treturn fmt.Errorf(\"%s[%%d]: %%w\", i, err)\n", f.Name)
				fmt.Fprintf(buf, "\t\t}\n")
				fmt.Fprintf(buf, "\t}\n")
			} else if def, ok := g.Definitions[baseType]; ok && def.Name != "" {
				fmt.Fprintf(buf, "\tfor i, item := range r.%s {\n", f.Name)
				fmt.Fprintf(buf, "\t\tif err := item.Validate(); err != nil {\n")
				fmt.Fprintf(buf, "\t\t\treturn fmt.Errorf(\"%s[%%d]: %%w\", i, err)\n", f.Name)
				fmt.Fprintf(buf, "\t\t}\n")
				fmt.Fprintf(buf, "\t}\n")
			}
		}
	}

	fmt.Fprintf(buf, "\treturn nil\n")
	fmt.Fprintf(buf, "}\n\n")
}

func (g *Generator) writeFixedValidation(buf *bytes.Buffer, fieldName string, fixed any, baseType string, isPointer bool) {
	var valueStr string
	switch v := fixed.(type) {
	case string:
		valueStr = fmt.Sprintf(`"%s"`, v)
	case float64:
		if baseType == "int" || baseType == "int64" {
			valueStr = fmt.Sprintf("%.0f", v)
		} else {
			valueStr = fmt.Sprintf("%v", v)
		}
	case bool:
		valueStr = fmt.Sprintf("%v", v)
	default:
		valueStr = fmt.Sprintf("%v", v)
	}

	if isPointer {
		fmt.Fprintf(buf, "\t\tif *r.%s != %s {\n", fieldName, valueStr)
	} else {
		fmt.Fprintf(buf, "\tif r.%s != %s {\n", fieldName, valueStr)
	}
	fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"field '%s' must be %v\")\n", fieldName, fixed)
	fmt.Fprintf(buf, "\t}\n")
}
