package tests

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

type ValidationTestCase struct {
	Name    string
	Data    string
	WantErr bool
	ErrMsg  string
}

func compileGeneratedCode(outputDir string) error {
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = outputDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go build failed: %w\nOutput: %s", err, string(output))
	}
	return nil
}

func createValidationTestFile(outputDir, structName, fileName string, testCases []ValidationTestCase) error {
	var buf strings.Builder
	buf.WriteString("package models\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\t\"strings\"\n")
	buf.WriteString("\t\"testing\"\n")
	buf.WriteString(")\n\n")
	buf.WriteString(fmt.Sprintf("func Test%s_Validate(t *testing.T) {\n", structName))
	buf.WriteString("\ttests := []struct {\n")
	buf.WriteString("\t\tname    string\n")
	buf.WriteString("\t\tdata    *" + structName + "\n")
	buf.WriteString("\t\twantErr bool\n")
	buf.WriteString("\t\terrMsg  string\n")
	buf.WriteString("\t}{\n")

	for _, tc := range testCases {
		buf.WriteString("\t\t{\n")
		buf.WriteString(fmt.Sprintf("\t\t\tname: %q,\n", tc.Name))
		buf.WriteString(fmt.Sprintf("\t\t\tdata: %s,\n", tc.Data))
		buf.WriteString(fmt.Sprintf("\t\t\twantErr: %v,\n", tc.WantErr))
		if tc.ErrMsg != "" {
			buf.WriteString(fmt.Sprintf("\t\t\terrMsg: %q,\n", tc.ErrMsg))
		}
		buf.WriteString("\t\t},\n")
	}

	buf.WriteString("\t}\n\n")
	buf.WriteString("\tfor _, tt := range tests {\n")
	buf.WriteString("\t\tt.Run(tt.name, func(t *testing.T) {\n")
	buf.WriteString("\t\t\terr := tt.data.Validate()\n")
	buf.WriteString("\t\t\tif (err != nil) != tt.wantErr {\n")
	buf.WriteString("\t\t\t\tt.Errorf(\"Validate() error = %v, wantErr %v\", err, tt.wantErr)\n")
	buf.WriteString("\t\t\t}\n")
	buf.WriteString("\t\t\tif tt.wantErr && tt.errMsg != \"\" && err != nil {\n")
	buf.WriteString("\t\t\t\tif !strings.Contains(err.Error(), tt.errMsg) {\n")
	buf.WriteString("\t\t\t\t\tt.Errorf(\"Validate() error = %v, want error containing %q\", err, tt.errMsg)\n")
	buf.WriteString("\t\t\t\t}\n")
	buf.WriteString("\t\t\t}\n")
	buf.WriteString("\t\t})\n")
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")

	testFileName := strings.TrimSuffix(fileName, ".go") + "_validation_test.go"
	testFilePath := filepath.Join(outputDir, testFileName)
	return os.WriteFile(testFilePath, []byte(buf.String()), 0644)
}

func runValidationTests(outputDir string) error {
	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Dir = outputDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go test failed: %w\nOutput: %s", err, string(output))
	}
	return nil
}
