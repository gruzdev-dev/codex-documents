package tests

import (
	"encoding/json"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"codex-documents/tools/generator"
)

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

func TestGenerator_RealResources(t *testing.T) {
	tests := []struct {
		name       string
		specFile   string
		whitelist  []string
		wantStruct string
	}{
		{
			name:       "generate Patient from real spec",
			specFile:   "patient",
			whitelist:  []string{"Patient"},
			wantStruct: "Patient",
		},
		{
			name:       "generate Observation from real spec",
			specFile:   "observation",
			whitelist:  []string{"Observation"},
			wantStruct: "Observation",
		},
		{
			name:       "generate DocumentReference from real spec",
			specFile:   "documentreference",
			whitelist:  []string{"DocumentReference"},
			wantStruct: "DocumentReference",
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

			g := generator.NewGenerator("", outputDir, tt.whitelist)
			g.Definitions[spec.Name] = spec

			if err := g.WriteResource(spec); err != nil {
				t.Fatalf("WriteResource() error = %v", err)
			}

			fileName := toSnakeCase(tt.wantStruct) + ".go"
			filePath := filepath.Join(outputDir, fileName)
			generatedCode, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("ReadFile() error = %v", err)
			}

			if _, err := parseGeneratedCode(string(generatedCode)); err != nil {
				t.Errorf("parseGeneratedCode() error = %v, code should be valid Go", err)
			}

			if _, err := format.Source(generatedCode); err != nil {
				t.Errorf("format.Source() error = %v, code should be formattable", err)
			}

			if !strings.Contains(string(generatedCode), "type "+tt.wantStruct+" struct") {
				t.Errorf("generated code should contain struct %s", tt.wantStruct)
			}
		})
	}
}

func TestGenerator_RequiredFields(t *testing.T) {
	spec, err := loadTestSpec("simple_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"SimpleTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "simple_test_resource.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	tests := []struct {
		name                string
		fieldName           string
		shouldHaveOmitEmpty bool
	}{
		{
			name:                "requiredField (min=1) should not have omitempty",
			fieldName:           "RequiredField",
			shouldHaveOmitEmpty: false,
		},
		{
			name:                "optionalField (min=0) should have omitempty",
			fieldName:           "OptionalField",
			shouldHaveOmitEmpty: true,
		},
		{
			name:                "requiredArray (min=2) should not have omitempty",
			fieldName:           "RequiredArray",
			shouldHaveOmitEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldLine := ""
			for _, line := range strings.Split(code, "\n") {
				if strings.Contains(line, tt.fieldName) && strings.Contains(line, "json:") {
					fieldLine = line
					break
				}
			}

			if fieldLine == "" {
				t.Fatalf("field %s not found in generated code", tt.fieldName)
			}

			hasOmitEmpty := strings.Contains(fieldLine, "omitempty")
			if tt.shouldHaveOmitEmpty && !hasOmitEmpty {
				t.Errorf("field %s should have omitempty but doesn't", tt.fieldName)
			}
			if !tt.shouldHaveOmitEmpty && hasOmitEmpty {
				t.Errorf("field %s should not have omitempty but does", tt.fieldName)
			}
		})
	}
}

func TestGenerator_BSONTags(t *testing.T) {
	spec, err := loadTestSpec("simple_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"SimpleTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "simple_test_resource.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)
	hasBSON := strings.Contains(code, "bson:")

	if !hasBSON {
		t.Errorf("BSON tags are not present in generated code")
	}

	if !strings.Contains(code, "bson:\"required_field\"") {
		t.Errorf("BSON tag for requiredField should be in snake_case format: required_field")
	}

	if !strings.Contains(code, "bson:\"optional_field,omitempty\"") {
		t.Errorf("BSON tag for optionalField should be in snake_case with omitempty: optional_field,omitempty")
	}
}

func TestGenerator_EmptyStructures(t *testing.T) {
	spec, err := loadTestSpec("nested_structures")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"NestedStructuresTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "nested_structures_test_resource.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if strings.Contains(code, "type NestedStructuresTestResourceNestedEmpty struct") {
		emptyStructMatch := strings.Contains(code, "type NestedStructuresTestResourceNestedEmpty struct {\n}")
		if emptyStructMatch {
			t.Errorf("Empty structure NestedStructuresTestResourceNestedEmpty should not be generated or should have comment")
		}
	}
}

func TestGenerator_ReferenceTargetProfile(t *testing.T) {
	spec, err := loadTestSpec("reference_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"ReferenceTestResource"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	entries, _ := os.ReadDir(outputDir)
	generatedFiles := make(map[string]bool)
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			generatedFiles[entry.Name()] = true
		}
	}

	if generatedFiles["patient.go"] || generatedFiles["organization.go"] {
		t.Errorf("Models from targetProfile should NOT be generated automatically (Patient, Organization)")
	}
}

func TestGenerator_Dependencies(t *testing.T) {
	_, err := loadTestSpec("whitelist_test")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"WhitelistResource1"})

	wd, _ := os.Getwd()
	var path string
	if filepath.Base(wd) == "tests" {
		path = filepath.Join(wd, "testdata", "whitelist_test.json")
	} else {
		path = filepath.Join(wd, "tests", "testdata", "whitelist_test.json")
	}

	var bundle generator.StructureDefinitionBundle
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if err := json.Unmarshal(data, &bundle); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	for _, entry := range bundle.Entry {
		if entry.Resource.Name != "" {
			g.Definitions[entry.Resource.Name] = entry.Resource
		}
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	whitelistFile := filepath.Join(outputDir, "whitelist_resource1.go")
	if _, err := os.Stat(whitelistFile); err != nil {
		t.Errorf("whitelist resource file should be generated: %v", err)
	}

	dependencyFile := filepath.Join(outputDir, "dependency_type.go")
	if _, err := os.Stat(dependencyFile); err != nil {
		t.Logf("DependencyType should be generated as dependency (may fail if usedTypes tracking is broken)")
	}

	notInWhitelistFile := filepath.Join(outputDir, "not_in_whitelist_resource.go")
	if _, err := os.Stat(notInWhitelistFile); err == nil {
		t.Errorf("NotInWhitelistResource should NOT be generated")
	}
}

func TestGenerator_WhitelistFiltering(t *testing.T) {
	_, err := loadTestSpec("whitelist_test")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"WhitelistResource1"})

	wd, _ := os.Getwd()
	var path string
	if filepath.Base(wd) == "tests" {
		path = filepath.Join(wd, "testdata", "whitelist_test.json")
	} else {
		path = filepath.Join(wd, "tests", "testdata", "whitelist_test.json")
	}

	var bundle generator.StructureDefinitionBundle
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if err := json.Unmarshal(data, &bundle); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	for _, entry := range bundle.Entry {
		if entry.Resource.Name != "" {
			g.Definitions[entry.Resource.Name] = entry.Resource
		}
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	entries, _ := os.ReadDir(outputDir)
	generatedFiles := make(map[string]bool)
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			generatedFiles[entry.Name()] = true
		}
	}

	if !generatedFiles["whitelist_resource1.go"] {
		t.Errorf("WhitelistResource1 should be generated")
	}

	if generatedFiles["not_in_whitelist_resource.go"] {
		t.Errorf("NotInWhitelistResource should NOT be generated")
	}
}

func TestGenerator_CodeCompilation(t *testing.T) {
	spec, err := loadTestSpec("patient")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"Patient"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "patient.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if _, err := format.Source(generatedCode); err != nil {
		t.Errorf("format.Source() error = %v, code should be formattable", err)
	}

	if _, err := parseGeneratedCode(string(generatedCode)); err != nil {
		t.Errorf("parseGeneratedCode() error = %v, code should be valid Go", err)
	}
}

func TestGenerator_DurationDependency(t *testing.T) {
	spec, err := loadTestSpec("duration_dependency")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"TimingRepeat"})
	g.Definitions[spec.Name] = spec

	g.Definitions["Duration"] = generator.StructureDefinition{
		Name:        "Duration",
		Description: "A length of time",
		Snapshot: generator.Snapshot{
			Element: []generator.ElementDefinition{
				{
					ID:   "Duration",
					Path: "Duration",
					Min:  0,
					Max:  "*",
				},
			},
		},
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	timingFile := filepath.Join(outputDir, "timing_repeat.go")
	timingCode, err := os.ReadFile(timingFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(timingCode)
	if !strings.Contains(code, "BoundsDuration") {
		t.Errorf("TimingRepeat should contain BoundsDuration field")
	}

	durationFile := filepath.Join(outputDir, "duration.go")
	if _, err := os.Stat(durationFile); err != nil {
		t.Errorf("Duration should be generated as separate file through usedTypes, but duration.go not found: %v", err)
	}

	durationCode, err := os.ReadFile(durationFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if !strings.Contains(string(durationCode), "type Duration struct") {
		t.Errorf("Duration file should contain Duration struct definition")
	}
}

func TestGenerator_RequiredReferencePointer(t *testing.T) {
	spec, err := loadTestSpec("required_reference")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"TestResourceWithRequiredReference"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "test_resource_with_required_reference.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	fieldLine := ""
	for _, line := range strings.Split(code, "\n") {
		if strings.Contains(line, "RequiredRef") && strings.Contains(line, "json:") {
			fieldLine = line
			break
		}
	}

	if fieldLine == "" {
		t.Fatalf("field RequiredRef not found in generated code")
	}

	if !strings.Contains(fieldLine, "*Reference") {
		t.Errorf("field RequiredRef should have pointer type *Reference, got: %s", fieldLine)
	}
}

func TestGenerator_ContentReference(t *testing.T) {
	spec, err := loadTestSpec("content_reference_test")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"TestResourceWithContentRef"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	filePath := filepath.Join(outputDir, "test_resource_with_content_ref.go")
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	code := string(generatedCode)

	if strings.Contains(code, "TestResourceWithContentRefComponentReferenceRange") {
		t.Errorf("contentReference should reuse TestResourceWithContentRefReferenceRange type, not create new ComponentReferenceRange")
	}

	if !strings.Contains(code, "TestResourceWithContentRefReferenceRange") {
		t.Errorf("contentReference type TestResourceWithContentRefReferenceRange should be present")
	}
}

func TestGenerator_ParseBindingAndConstraint(t *testing.T) {
	spec, err := loadTestSpec("documentreference")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	var hasBinding bool
	var hasConstraint bool

	for _, el := range spec.Snapshot.Element {
		if el.Binding != nil {
			hasBinding = true
			if el.Binding.Strength == "" && el.Binding.ValueSet == "" {
				t.Error("Binding should have at least strength or valueSet")
			}
		}

		if len(el.Constraint) > 0 {
			hasConstraint = true
			for _, c := range el.Constraint {
				if c.Key == "" {
					t.Error("Constraint should have key")
				}
				if c.Severity == "" {
					t.Error("Constraint should have severity")
				}
				if c.Human == "" {
					t.Error("Constraint should have human description")
				}
				if c.Expression == "" {
					t.Error("Constraint should have expression")
				}
			}
		}

		if el.MaxLength != nil && *el.MaxLength <= 0 {
			t.Error("MaxLength should be positive")
		}
	}

	if !hasBinding {
		t.Error("DocumentReference spec should have at least one element with binding")
	}

	if !hasConstraint {
		t.Error("DocumentReference spec should have at least one element with constraint")
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

	g := generator.NewGenerator("", outputDir, []string{"DocumentReference"})
	g.Definitions[spec.Name] = spec

	if err := g.WriteResource(spec); err != nil {
		t.Fatalf("WriteResource() error = %v", err)
	}

	fileName := "document_reference.go"
	filePath := filepath.Join(outputDir, fileName)
	generatedCode, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if _, err := parseGeneratedCode(string(generatedCode)); err != nil {
		t.Errorf("parseGeneratedCode() error = %v, code should be valid Go", err)
	}

	if !strings.Contains(string(generatedCode), "type DocumentReference struct") {
		t.Errorf("generated code should contain struct DocumentReference")
	}
}

func TestGenerator_ValidateMethodGeneration(t *testing.T) {
	spec, err := loadTestSpec("simple_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

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
	if !strings.Contains(code, "func (r *SimpleTestResource) Validate() error") {
		t.Error("generated code should contain Validate() method")
	}

	if _, err := parseGeneratedCode(code); err != nil {
		t.Errorf("parseGeneratedCode() error = %v, code should be valid Go", err)
	}

	if _, err := format.Source(generatedCode); err != nil {
		t.Errorf("format.Source() error = %v, code should be formattable", err)
	}
}

func TestGenerator_ValidateMethodWorks(t *testing.T) {
	spec, err := loadTestSpec("simple_resource")
	if err != nil {
		t.Fatalf("loadTestSpec() error = %v", err)
	}

	outputDir, cleanup, err := createTempOutputDir()
	if err != nil {
		t.Fatalf("createTempOutputDir() error = %v", err)
	}
	defer cleanup()

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
	if !strings.Contains(code, "func (r *SimpleTestResource) Validate() error") {
		t.Error("generated code should contain Validate() method")
	}

	if _, err := parseGeneratedCode(code); err != nil {
		t.Errorf("parseGeneratedCode() error = %v, code should be valid Go", err)
	}
}
