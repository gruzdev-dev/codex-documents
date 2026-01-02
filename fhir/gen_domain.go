//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var whiteList = map[string]struct{}{
	"Patient":           {},
	"DocumentReference": {},
	"Binary":            {},
	"Observation":       {},
	"Bundle":            {},
	"OperationOutcome":  {},
	"Encounter":         {},
	"Practitioner":      {},
}

var coreTypes = map[string]struct{}{
	"CodeableConcept": {}, "Coding": {}, "Identifier": {},
	"Reference": {}, "Quantity": {}, "Period": {},
	"Attachment": {}, "Meta": {}, "HumanName": {},
	"Address": {}, "ContactPoint": {}, "Money": {},
	"Extension": {}, "Element": {}, "String": {},
	"Boolean": {}, "Integer": {}, "PositiveInt": {},
	"DateTime": {}, "Instant": {}, "Range": {}, "Ratio": {},
}

var typesToGenerate = make(map[string]*ast.TypeSpec)
var allTypes = make(map[string]*ast.TypeSpec)

func main() {
	sanitizeModelsFile("fhir/models.go")

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "fhir/models.go", nil, 0)
	if err != nil {
		panic(err)
	}

	for _, decl := range node.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
			for _, spec := range gen.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					allTypes[ts.Name.Name] = ts
				}
			}
		}
	}

	for res := range whiteList {
		collectDeps(res)
	}

	outputDir := "core/domain/fhir"
	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, 0755)

	// Буфер для всех базовых типов
	var coreBuf bytes.Buffer
	fmt.Fprintf(&coreBuf, "package fhir\n\n")

	for name, spec := range typesToGenerate {
		cleanSpec(spec)

		if _, isCore := coreTypes[name]; isCore {
			fmt.Fprintf(&coreBuf, "type ")
			format.Node(&coreBuf, token.NewFileSet(), spec)
			fmt.Fprintf(&coreBuf, "\n\n")
		} else {
			generateResourceFile(name, spec, outputDir)
		}
	}

	formattedCore, _ := format.Source(coreBuf.Bytes())
	os.WriteFile(filepath.Join(outputDir, "core_types.go"), formattedCore, 0644)

	fmt.Printf("Done. Generated models: %d\n", len(typesToGenerate))
}

func collectDeps(name string) {
	if _, exists := typesToGenerate[name]; exists {
		return
	}
	spec, ok := allTypes[name]
	if !ok {
		return
	}
	typesToGenerate[name] = spec

	ast.Inspect(spec.Type, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			if _, isCore := coreTypes[ident.Name]; isCore {
				collectDeps(ident.Name)
			} else if _, isWhite := whiteList[ident.Name]; isWhite {
				collectDeps(ident.Name)
			}
		}
		return true
	})
}

func cleanSpec(spec *ast.TypeSpec) {
	if st, ok := spec.Type.(*ast.StructType); ok {
		var filteredFields []*ast.Field
		for _, field := range st.Fields.List {
			fieldType := getFieldTypeName(field.Type)
			_, isWhite := whiteList[fieldType]
			_, isCore := coreTypes[fieldType]

			if isWhite || isCore || isSimpleType(fieldType) {
				field.Tag = nil
				filteredFields = append(filteredFields, field)
			}
		}
		st.Fields.List = filteredFields
	}
}

func generateResourceFile(name string, spec *ast.TypeSpec, dir string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package fhir\n\ntype ")
	format.Node(&buf, token.NewFileSet(), spec)

	formatted, _ := format.Source(buf.Bytes())
	fileName := toSnakeCase(name) + ".go"
	os.WriteFile(filepath.Join(dir, fileName), formatted, 0644)
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getFieldTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return getFieldTypeName(t.X)
	case *ast.ArrayType:
		return getFieldTypeName(t.Elt)
	default:
		return ""
	}
}

func isSimpleType(name string) bool {
	return name == "string" || name == "bool" || name == "int" || strings.HasSuffix(name, "Enum")
}

func sanitizeModelsFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`const (\w+)_Undefined (\w+) = "([^"]+)"`)

	sanitized := re.ReplaceAllFunc(content, func(match []byte) []byte {
		submatches := re.FindSubmatch(match)
		prefix := string(submatches[1])
		typeName := string(submatches[2])
		operator := string(submatches[3])

		suffix := "Undefined"
		switch operator {
		case ">=":
			suffix = "GE"
		case ">":
			suffix = "GT"
		case "<=":
			suffix = "LE"
		case "<":
			suffix = "LT"
		default:
			suffix = fmt.Sprintf("Val_%x", operator)
		}

		return []byte(fmt.Sprintf("const %s_%s %s = %q", prefix, suffix, typeName, operator))
	})

	err = os.WriteFile(path, sanitized, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sanitize done!")
}
