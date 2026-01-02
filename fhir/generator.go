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
	"Uri": {}, "Code": {}, "Markdown": {}, "Id": {},
}

var typesToGenerate = make(map[string]*ast.TypeSpec)
var allTypes = make(map[string]*ast.TypeSpec)
var resourceTypes = make(map[string]struct{})

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
					if isResource(ts) {
						resourceTypes[ts.Name.Name] = struct{}{}
					}
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

	var coreBuf bytes.Buffer
	fmt.Fprintf(&coreBuf, "package fhir\n\n")

	for name, spec := range typesToGenerate {
		cleanSpec(spec)
		if _, isCore := coreTypes[name]; isCore || !isResourceName(name) {
			if _, isWhite := whiteList[name]; !isWhite {
				fmt.Fprintf(&coreBuf, "type ")
				format.Node(&coreBuf, token.NewFileSet(), spec)
				fmt.Fprintf(&coreBuf, "\n\n")
				continue
			}
		}
		generateResourceFile(name, spec, outputDir)
	}

	formattedCore, _ := format.Source(coreBuf.Bytes())
	os.WriteFile(filepath.Join(outputDir, "core_types.go"), formattedCore, 0644)

	generateMapperInterface(whiteList, "adapters/fhir")
	fmt.Printf("Success. Models generated: %d\n", len(typesToGenerate))
}

func isResource(ts *ast.TypeSpec) bool {
	st, ok := ts.Type.(*ast.StructType)
	if !ok { return false }
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			if name.Name == "ResourceType" { return true }
		}
	}
	return false
}

func isResourceName(name string) bool {
	_, ok := resourceTypes[name]
	return ok
}

func collectDeps(name string) {
	if _, exists := typesToGenerate[name]; exists { return }
	spec, ok := allTypes[name]
	if !ok { return }

	typesToGenerate[name] = spec

	ast.Inspect(spec.Type, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok { return true }

		depName := ident.Name
		if _, exists := allTypes[depName]; !exists { return true }

		if isResourceName(depName) {
			if _, isWhite := whiteList[depName]; !isWhite {
				return false 
			}
		}

		collectDeps(depName)
		return true
	})
}

func cleanSpec(spec *ast.TypeSpec) {
	spec.Doc = nil
	if st, ok := spec.Type.(*ast.StructType); ok {
		var filteredFields []*ast.Field
		for _, field := range st.Fields.List {
			field.Doc = nil
			field.Comment = nil

			for _, name := range field.Names {
				if name.Name == "ResourceType" {
					field.Type = &ast.Ident{Name: "interface{}"}
					break
				}
			}

			fieldType := getFieldTypeName(field.Type)
			if isResourceName(fieldType) {
				if _, isWhite := whiteList[fieldType]; !isWhite {
					continue
				}
			}

			field.Tag = nil
			filteredFields = append(filteredFields, field)
		}
		st.Fields.List = filteredFields
	}
}

func sanitizeModelsFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil { return }
	re := regexp.MustCompile(`const (\w+)_Undefined (\w+) = "([^"]+)"`)
	sanitized := re.ReplaceAllFunc(content, func(match []byte) []byte {
		sub := re.FindSubmatch(match)
		op := string(sub[3])
		suffix := "Undefined"
		switch op {
		case ">=": suffix = "GE"
		case ">":  suffix = "GT"
		case "<=": suffix = "LE"
		case "<":  suffix = "LT"
		}
		return []byte(fmt.Sprintf("const %s_%s %s = %q", string(sub[1]), suffix, string(sub[2]), op))
	})
	os.WriteFile(path, sanitized, 0644)
	fmt.Println("Models sanitized: AgeComparator constants fixed.")
}

func generateResourceFile(name string, spec *ast.TypeSpec, dir string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package fhir\n\ntype ")
	format.Node(&buf, token.NewFileSet(), spec)
	formatted, _ := format.Source(buf.Bytes())
	os.WriteFile(filepath.Join(dir, toSnakeCase(name)+".go"), formatted, 0644)
}

func generateMapperInterface(resources map[string]struct{}, outputDir string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package fhir\n\nimport (\n\tdto \"codex-documents/fhir\"\n\tdomain \"codex-documents/core/domain/fhir\"\n)\n\n")
	fmt.Fprintf(&buf, "// goverter:converter\n")
	fmt.Fprintf(&buf, "// goverter:output:file ./generated.go\n")
	fmt.Fprintf(&buf, "// goverter:output:package codex-documents/adapters/fhir\n")
	fmt.Fprintf(&buf, "// goverter:extend MapAny\n")
	fmt.Fprintf(&buf, "type Converter interface {\n")
	for res := range resources {
		fmt.Fprintf(&buf, "\tToDomain%s(source *dto.%s) *domain.%s\n", res, res, res)
		fmt.Fprintf(&buf, "\tFromDomain%s(source *domain.%s) *dto.%s\n\n", res, res, res)
	}
	fmt.Fprintf(&buf, "}\n\n")
	fmt.Fprintf(&buf, "func MapAny(v interface{}) interface{} { return v }\n")

	os.MkdirAll(outputDir, 0755)
	formatted, _ := format.Source(buf.Bytes())
	os.WriteFile(filepath.Join(outputDir, "mapper_interface.go"), formatted, 0644)
	fmt.Println("Mapper interface generated in adapters/fhir/mapper_interface.go")
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
	case *ast.Ident: return t.Name
	case *ast.StarExpr: return getFieldTypeName(t.X)
	case *ast.ArrayType: return getFieldTypeName(t.Elt)
	default: return ""
	}
}