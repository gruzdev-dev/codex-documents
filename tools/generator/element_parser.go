package generator

import (
	"fmt"
	"strings"
)

type FieldInfo struct {
	Name    string
	GoType  string
	JSONTag string
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

		// Handle contentReference before processing type
		if el.ContentReference != "" && strings.HasPrefix(el.ContentReference, "#") {
			// Extract path from contentReference (e.g., "#Observation.referenceRange" -> "Observation.referenceRange")
			refPath := strings.TrimPrefix(el.ContentReference, "#")
			// Use the referenced path to determine the type
			goType := g.deriveNestedTypeName(refPath)
			if goType == "" {
				goType = "any"
			}

			cleanName := strings.ReplaceAll(lastPart, "[x]", "")
			cleanName = strings.ReplaceAll(cleanName, "-", "_")
			cleanName = titleCase(cleanName)

			if cleanName == "" {
				continue
			}

			if el.Max == "*" {
				goType = "[]" + goType
			} else if el.Min == 0 && !strings.HasPrefix(goType, "[]") && goType != "bool" && goType != "json.RawMessage" && goType != "any" {
				goType = "*" + goType
			} else if el.Min > 0 && !strings.HasPrefix(goType, "[]") {
				// For required non-primitive types (like Reference, custom types), always use pointer except for arrays and built-in types
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

			structs[structName] = append(structs[structName], FieldInfo{
				Name:    cleanName,
				GoType:  goType,
				JSONTag: jsonTag,
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

				typeSuffix := titleCase(fhirType.Code)
				fieldName := titleCase(baseName) + typeSuffix
				jsonTag := fmt.Sprintf("`json:\"%s%s\"`", baseName, typeSuffix)
				if el.Min == 0 {
					jsonTag = fmt.Sprintf("`json:\"%s%s,omitempty\"`", baseName, typeSuffix)
				}

				if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") && goType != "any" {
					goType = "*" + goType
				}

				structs[structName] = append(structs[structName], FieldInfo{
					Name:    fieldName,
					GoType:  goType,
					JSONTag: jsonTag,
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
		cleanName = titleCase(cleanName)

		if cleanName == "" {
			continue
		}

		if el.Max == "*" {
			goType = "[]" + goType
		} else if el.Min == 0 && !strings.HasPrefix(goType, "[]") && goType != "bool" && goType != "json.RawMessage" && goType != "any" {
			goType = "*" + goType
		} else if el.Min > 0 && !strings.HasPrefix(goType, "[]") {
			// For required non-primitive types (like Reference, custom types), always use pointer except for arrays and built-in types
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

		structs[structName] = append(structs[structName], FieldInfo{
			Name:    cleanName,
			GoType:  goType,
			JSONTag: jsonTag,
			Comment: el.Short,
		})
	}
	return structs
}
