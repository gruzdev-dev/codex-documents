package tests

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"codex-documents/tools/generator"
)

func loadTestSpec(name string) (generator.StructureDefinition, error) {
	wd, err := os.Getwd()
	if err != nil {
		return generator.StructureDefinition{}, fmt.Errorf("getwd: %w", err)
	}

	var path string
	if filepath.Base(wd) == "tests" {
		path = filepath.Join(wd, "testdata", name+".json")
	} else {
		path = filepath.Join(wd, "tests", "testdata", name+".json")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return generator.StructureDefinition{}, fmt.Errorf("read file %s: %w", path, err)
	}

	var bundle generator.StructureDefinitionBundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return generator.StructureDefinition{}, fmt.Errorf("unmarshal %s: %w", name, err)
	}

	if len(bundle.Entry) == 0 {
		return generator.StructureDefinition{}, fmt.Errorf("no entries in bundle %s", name)
	}

	return bundle.Entry[0].Resource, nil
}

func parseGeneratedCode(code string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", code, parser.ParseComments)
}

func createTempOutputDir() (string, func(), error) {
	dir, err := os.MkdirTemp("", "fhir_generator_test_*")
	if err != nil {
		return "", nil, err
	}

	cleanup := func() {
		os.RemoveAll(dir)
	}

	return dir, cleanup, nil
}
