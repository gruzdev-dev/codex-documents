package generator

import (
	"fmt"
	"strings"

	"codex-documents/tools/text"
)

type FieldInfo struct {
	Name    string
	GoType  string
	JSONTag string
	BSONTag string
	Comment string
}

func (g *Generator) ProcessElements(name string, elements []ElementDefinition) map[string][]FieldInfo {
	structs := make(map[string][]FieldInfo)

	for _, el := range elements {
		if el.Path == name || strings.Contains(el.ID, "._") {
			continue
		}

		parts := strings.Split(el.Path, ".")
		lastPart := parts[len(parts)-1]

		if strings.HasPrefix(lastPart, "_") || lastPart == "extension" || lastPart == "modifierExtension" {
			continue
		}

		var structName string
		if len(parts) == 1 {
			structName = name
		} else {
			parentPath := strings.Join(parts[:len(parts)-1], ".")
			structName = g.deriveNestedTypeName(parentPath)
		}

		if el.ContentReference != "" && strings.HasPrefix(el.ContentReference, "#") {
			refPath := strings.TrimPrefix(el.ContentReference, "#")
			goType := g.deriveNestedTypeName(refPath)
			if goType == "" {
				goType = "any"
			}

			cleanName := strings.ReplaceAll(lastPart, "[x]", "")
			cleanName = strings.ReplaceAll(cleanName, "-", "_")
			cleanName = text.TitleCase(cleanName)

			if cleanName == "" {
				continue
			}

			if el.Max == "*" {
				goType = "[]" + goType
			} else if el.Min == 0 && !strings.HasPrefix(goType, "[]") && goType != "bool" && goType != "json.RawMessage" && goType != "any" {
				goType = "*" + goType
			} else if el.Min > 0 && !strings.HasPrefix(goType, "[]") {
				isPrimitive := goType == "bool" || goType == "string" || goType == "int" || goType == "int64" || goType == "float64" || goType == "json.RawMessage" || goType == "any"
				if !isPrimitive {
					goType = "*" + goType
				}
			}

			escapedTag := strings.ReplaceAll(lastPart, "\"", "\\\"")
			jsonTag := fmt.Sprintf("`json:\"%s\"`", escapedTag)
			if el.Min == 0 {
				jsonTag = fmt.Sprintf("`json:\"%s,omitempty\"`", escapedTag)
			}
			bsonTag := generateBSONTag(cleanName, el.Min == 0)

			structs[structName] = append(structs[structName], FieldInfo{
				Name:    cleanName,
				GoType:  goType,
				JSONTag: jsonTag,
				BSONTag: bsonTag,
				Comment: el.Short,
			})
			continue
		}

		if strings.Contains(lastPart, "[x]") {
			baseName := strings.ReplaceAll(lastPart, "[x]", "")
			for _, fhirType := range el.Type {
				singleTypeEl := el
				singleTypeEl.Type = []ElementDataType{fhirType}

				goType := g.mapGoType(singleTypeEl)

				typeSuffix := text.TitleCase(fhirType.Code)
				fieldName := text.TitleCase(baseName) + typeSuffix
				jsonTag := fmt.Sprintf("`json:\"%s%s\"`", baseName, typeSuffix)
				if el.Min == 0 {
					jsonTag = fmt.Sprintf("`json:\"%s%s,omitempty\"`", baseName, typeSuffix)
				}
				bsonTag := generateBSONTag(fieldName, el.Min == 0)

				if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") && goType != "any" {
					goType = "*" + goType
				}

				structs[structName] = append(structs[structName], FieldInfo{
					Name:    fieldName,
					GoType:  goType,
					JSONTag: jsonTag,
					BSONTag: bsonTag,
					Comment: el.Short,
				})
			}
			continue
		}

		goType := g.mapGoType(el)

		if len(el.Type) > 0 &&
			(el.Type[0].Code == "BackboneElement" || el.Type[0].Code == "Element") {
			nestedStructName := g.deriveNestedTypeName(el.Path)
			if nestedStructName != "" {
				if _, exists := structs[nestedStructName]; !exists {
					structs[nestedStructName] = []FieldInfo{}
				}
			}
		}

		cleanName := strings.ReplaceAll(lastPart, "[x]", "")
		cleanName = strings.ReplaceAll(cleanName, "-", "_")
		cleanName = text.TitleCase(cleanName)

		if cleanName == "" {
			continue
		}

		if el.Max == "*" {
			goType = "[]" + goType
		} else if el.Min == 0 && !strings.HasPrefix(goType, "[]") && goType != "bool" && goType != "json.RawMessage" && goType != "any" {
			goType = "*" + goType
		} else if el.Min > 0 && !strings.HasPrefix(goType, "[]") {
			isPrimitive := goType == "bool" || goType == "string" || goType == "int" || goType == "int64" || goType == "float64" || goType == "json.RawMessage" || goType == "any"
			if !isPrimitive {
				goType = "*" + goType
			}
		}

		escapedTag := strings.ReplaceAll(lastPart, "\"", "\\\"")
		jsonTag := fmt.Sprintf("`json:\"%s\"`", escapedTag)
		if el.Min == 0 {
			jsonTag = fmt.Sprintf("`json:\"%s,omitempty\"`", escapedTag)
		}
		bsonTag := generateBSONTag(cleanName, el.Min == 0)

		structs[structName] = append(structs[structName], FieldInfo{
			Name:    cleanName,
			GoType:  goType,
			JSONTag: jsonTag,
			BSONTag: bsonTag,
			Comment: el.Short,
		})
	}
	return structs
}

func generateBSONTag(fieldName string, hasOmitEmpty bool) string {
	snakeName := text.ToSnakeCase(fieldName)
	bsonTag := fmt.Sprintf("`bson:\"%s\"`", snakeName)
	if hasOmitEmpty {
		bsonTag = fmt.Sprintf("`bson:\"%s,omitempty\"`", snakeName)
	}
	return bsonTag
}
