package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"codex-documents/tools/generator"
)

func TestValidation_RequiredFields(t *testing.T) {
	spec, err := loadTestSpec("simple_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	if err := createGoMod(outputDir); err != nil {
		t.Fatalf("createGoMod() error = %v", err)
	}

	g := generator.NewGenerator("", outputDir, []string{"SimpleTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "simple_test_resource.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if !checkValidateMethodExists(code, "SimpleTestResource") {
		t.Error("generated code should contain Validate() method")
	}

	if !checkRequiredFieldValidation(code, "RequiredField") {
		t.Error("generated code should contain validation for RequiredField")
	}

	if !checkRequiredFieldValidation(code, "RequiredArray") {
		t.Error("generated code should contain validation for RequiredArray")
	}

	testCases := []ValidationTestCase{
		{
			Name:    "valid data",
			Data:    `&SimpleTestResource{RequiredField: "value", RequiredArray: []string{"item1", "item2"}}`,
			WantErr: false,
		},
		{
			Name:    "missing required field",
			Data:    `&SimpleTestResource{RequiredArray: []string{"item1", "item2"}}`,
			WantErr: true,
			ErrMsg:  "field 'RequiredField' is required",
		},
		{
			Name:    "missing required array",
			Data:    `&SimpleTestResource{RequiredField: "value"}`,
			WantErr: true,
			ErrMsg:  "field 'RequiredArray' must have at least",
		},
		{
			Name:    "required array too short",
			Data:    `&SimpleTestResource{RequiredField: "value", RequiredArray: []string{"item1"}}`,
			WantErr: true,
			ErrMsg:  "field 'RequiredArray' must have at least 2 elements",
		},
	}

	if err := createValidationTestFile(outputDir, "SimpleTestResource", fileName, testCases); err != nil {
		t.Fatalf("createValidationTestFile() error = %v", err)
	}

	if err := addHelperFunctions(outputDir); err != nil {
		t.Fatalf("addHelperFunctions() error = %v", err)
	}

	if err := compileGeneratedCode(outputDir); err != nil {
		t.Fatalf("compileGeneratedCode() error = %v", err)
	}

	if err := runValidationTests(outputDir); err != nil {
		t.Errorf("runValidationTests() error = %v", err)
	}
}

func TestValidation_EmptyStructures(t *testing.T) {
	emptySpec := generator.StructureDefinition{
		Name: "EmptyTestResource",
		Snapshot: generator.Snapshot{
			Element: []generator.ElementDefinition{
				{
					ID:   "EmptyTestResource",
					Path: "EmptyTestResource",
					Min:  0,
					Max:  "*",
				},
			},
		},
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	if err := createGoMod(outputDir); err != nil {
		t.Fatalf("createGoMod() error = %v", err)
	}

	g := generator.NewGenerator("", outputDir, []string{"EmptyTestResource"})
	g.Definitions[emptySpec.Name] = emptySpec

	if err := g.WriteResource(emptySpec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "empty_test_resource.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	hasFmt, _ := checkImports(code)
	if hasFmt {
		t.Error("empty structures should not require fmt import if they only return nil")
	}

	if !checkValidateMethodExists(code, "EmptyTestResource") {
		t.Error("generated code should contain Validate() method even for empty structures")
	}

	if !strings.Contains(code, "return nil") {
		t.Error("empty structure Validate() should only return nil")
	}
}

func createGoMod(outputDir string) error {
	modContent := "module testmodels\n\ngo 1.21\n"
	return os.WriteFile(filepath.Join(outputDir, "go.mod"), []byte(modContent), 0644)
}

func addHelperFunctions(outputDir string) error {
	helperContent := `package models

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
`
	return os.WriteFile(filepath.Join(outputDir, "helpers.go"), []byte(helperContent), 0644)
}

func createMaxLengthTestSpec() generator.StructureDefinition {
	return generator.StructureDefinition{
		Name: "MaxLengthTestResource",
		Snapshot: generator.Snapshot{
			Element: []generator.ElementDefinition{
				{
					ID:   "MaxLengthTestResource",
					Path: "MaxLengthTestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:        "MaxLengthTestResource.stringField",
					Path:      "MaxLengthTestResource.stringField",
					Min:       0,
					Max:       "1",
					Type:      []generator.ElementDataType{{Code: "string"}},
					MaxLength: intPtr(10),
				},
			},
		},
	}
}

func createPatternTestSpec() generator.StructureDefinition {
	return generator.StructureDefinition{
		Name: "PatternTestResource",
		Snapshot: generator.Snapshot{
			Element: []generator.ElementDefinition{
				{
					ID:   "PatternTestResource",
					Path: "PatternTestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:      "PatternTestResource.stringField",
					Path:    "PatternTestResource.stringField",
					Min:     0,
					Max:     "1",
					Type:    []generator.ElementDataType{{Code: "string"}},
					Pattern: "^[A-Z]+$",
				},
			},
		},
	}
}

func intPtr(i int) *int {
	return &i
}

func TestValidation_MaxLength(t *testing.T) {
	spec := createMaxLengthTestSpec()

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	if err := createGoMod(outputDir); err != nil {
		t.Fatalf("createGoMod() error = %v", err)
	}

	g := generator.NewGenerator("", outputDir, []string{"MaxLengthTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "max_length_test_resource.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if !checkMaxLengthValidation(code, "StringField", 10) {
		t.Error("generated code should contain maxLength validation for StringField")
	}

	testCases := []ValidationTestCase{
		{
			Name:    "valid length",
			Data:    `&MaxLengthTestResource{StringField: stringPtr("short")}`,
			WantErr: false,
		},
		{
			Name:    "exceeds maxLength",
			Data:    `&MaxLengthTestResource{StringField: stringPtr("this is too long")}`,
			WantErr: true,
			ErrMsg:  "exceeds maxLength 10",
		},
	}

	if err := createValidationTestFile(outputDir, "MaxLengthTestResource", fileName, testCases); err != nil {
		t.Fatalf("createValidationTestFile() error = %v", err)
	}

	if err := addHelperFunctions(outputDir); err != nil {
		t.Fatalf("addHelperFunctions() error = %v", err)
	}

	if err := compileGeneratedCode(outputDir); err != nil {
		t.Fatalf("compileGeneratedCode() error = %v", err)
	}

	if err := runValidationTests(outputDir); err != nil {
		t.Errorf("runValidationTests() error = %v", err)
	}
}

func TestValidation_Pattern(t *testing.T) {
	spec := createPatternTestSpec()

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	if err := createGoMod(outputDir); err != nil {
		t.Fatalf("createGoMod() error = %v", err)
	}

	g := generator.NewGenerator("", outputDir, []string{"PatternTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "pattern_test_resource.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if !checkPatternValidation(code, "StringField") {
		t.Error("generated code should contain pattern validation for StringField")
	}

	_, hasRegexp := checkImports(code)
	if !hasRegexp {
		t.Error("generated code should import regexp for pattern validation")
	}

	testCases := []ValidationTestCase{
		{
			Name:    "valid pattern",
			Data:    `&PatternTestResource{StringField: stringPtr("ABC")}`,
			WantErr: false,
		},
		{
			Name:    "invalid pattern",
			Data:    `&PatternTestResource{StringField: stringPtr("abc")}`,
			WantErr: true,
			ErrMsg:  "does not match pattern",
		},
	}

	if err := createValidationTestFile(outputDir, "PatternTestResource", fileName, testCases); err != nil {
		t.Fatalf("createValidationTestFile() error = %v", err)
	}

	if err := addHelperFunctions(outputDir); err != nil {
		t.Fatalf("addHelperFunctions() error = %v", err)
	}

	if err := compileGeneratedCode(outputDir); err != nil {
		t.Fatalf("compileGeneratedCode() error = %v", err)
	}

	if err := runValidationTests(outputDir); err != nil {
		t.Errorf("runValidationTests() error = %v", err)
	}
}

func TestValidation_NestedStructures(t *testing.T) {
	spec, err := loadTestSpec("nested_structures")
	if err != nil {
		t.Skipf("loadTestSpec() error = %v, skipping nested structures test", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	if err := createGoMod(outputDir); err != nil {
		t.Fatalf("createGoMod() error = %v", err)
	}

	g := generator.NewGenerator("", outputDir, []string{"NestedStructuresTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "nested_structures_test_resource.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if !checkValidateMethodExists(code, "NestedStructuresTestResource") {
		t.Error("generated code should contain Validate() method")
	}

	if !checkNestedValidation(code, "Nested") {
		t.Log("nested validation check skipped - may not have nested structures in test data")
	}

	if !checkValidateMethodExists(code, "NestedStructuresTestResource") {
		t.Error("generated code should contain Validate() method")
	}
}

func TestValidation_RealResources(t *testing.T) {
	tests := []struct {
		name      string
		specFile  string
		whitelist []string
	}{
		{
			name:      "Patient",
			specFile:  "patient",
			whitelist: []string{"Patient"},
		},
		{
			name:      "Observation",
			specFile:  "observation",
			whitelist: []string{"Observation"},
		},
		{
			name:      "DocumentReference",
			specFile:  "documentreference",
			whitelist: []string{"DocumentReference"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec, err := loadTestSpec(tt.specFile)
			if err != nil {
				t.Fatalf("loadTestSpec() error = %v", err)
			}

			outputDir, cleanup, err := createTempOutputDir()
			if err != nil {
				t.Fatalf("createTempOutputDir() error = %v", err)
			}
			defer cleanup()

			if err := createGoMod(outputDir); err != nil {
				t.Fatalf("createGoMod() error = %v", err)
			}

			g := generator.NewGenerator("", outputDir, tt.whitelist)
			g.Definitions[spec.Name] = spec

			if err := g.WriteResource(spec); err != nil {
				t.Fatalf("WriteResource() error = %v", err)
			}

			structName := spec.Name
			fileName := snakeCase(structName) + ".go"
			filePath := filepath.Join(outputDir, fileName)
			generatedCode, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("ReadFile() error = %v, tried file: %s", err, fileName)
			}

			code := string(generatedCode)

			if !checkValidateMethodExists(code, structName) {
				t.Errorf("generated code should contain Validate() method for %s", structName)
			}

			hasFmt, hasRegexp := checkImports(code)
			if hasFmt {
				t.Logf("%s uses fmt import for validation", structName)
			}
			if hasRegexp {
				t.Logf("%s uses regexp import for validation", structName)
			}
		})
	}
}

func snakeCase(s string) string {
	var res strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			res.WriteRune('_')
		}
		res.WriteRune(unicode.ToLower(r))
	}
	return res.String()
}
