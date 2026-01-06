package generator

import (
	"testing"
)

func TestNewGenerator(t *testing.T) {
	tests := []struct {
		name      string
		specPath  string
		outPath   string
		whitelist []string
		want      map[string]struct{}
	}{
		{
			name:      "create generator with whitelist",
			specPath:  "test/spec",
			outPath:   "test/output",
			whitelist: []string{"Patient", "Observation"},
			want: map[string]struct{}{
				"Patient":     {},
				"Observation": {},
			},
		},
		{
			name:      "create generator with empty whitelist",
			specPath:  "test/spec",
			outPath:   "test/output",
			whitelist: []string{},
			want:      map[string]struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenerator(tt.specPath, tt.outPath, tt.whitelist)

			if got.SpecPath != tt.specPath {
				t.Errorf("SpecPath = %v, want %v", got.SpecPath, tt.specPath)
			}

			if got.OutputPath != tt.outPath {
				t.Errorf("OutputPath = %v, want %v", got.OutputPath, tt.outPath)
			}

			if len(got.Whitelist) != len(tt.want) {
				t.Errorf("Whitelist length = %v, want %v", len(got.Whitelist), len(tt.want))
			}

			for key := range tt.want {
				if _, ok := got.Whitelist[key]; !ok {
					t.Errorf("Whitelist missing key: %v", key)
				}
			}
		})
	}
}

func TestExtractBaseType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "pointer type",
			input: "*string",
			want:  "string",
		},
		{
			name:  "slice type",
			input: "[]string",
			want:  "string",
		},
		{
			name:  "pointer to slice",
			input: "*[]string",
			want:  "string",
		},
		{
			name:  "simple type",
			input: "string",
			want:  "string",
		},
		{
			name:  "nested pointer",
			input: "**string",
			want:  "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractBaseType(tt.input)
			if got != tt.want {
				t.Errorf("extractBaseType(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsBuiltinType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "string is builtin",
			input: "string",
			want:  true,
		},
		{
			name:  "bool is builtin",
			input: "bool",
			want:  true,
		},
		{
			name:  "int is builtin",
			input: "int",
			want:  true,
		},
		{
			name:  "float64 is builtin",
			input: "float64",
			want:  true,
		},
		{
			name:  "any is builtin",
			input: "any",
			want:  true,
		},
		{
			name:  "json.RawMessage is builtin",
			input: "json.RawMessage",
			want:  true,
		},
		{
			name:  "custom type is not builtin",
			input: "CustomType",
			want:  false,
		},
		{
			name:  "Reference is not builtin",
			input: "Reference",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBuiltinType(tt.input)
			if got != tt.want {
				t.Errorf("isBuiltinType(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

