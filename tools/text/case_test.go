package text

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple case",
			input: "Patient",
			want:  "patient",
		},
		{
			name:  "multiple words",
			input: "DocumentReference",
			want:  "document_reference",
		},
		{
			name:  "single word lowercase",
			input: "observation",
			want:  "observation",
		},
		{
			name:  "mixed case",
			input: "CodeableConcept",
			want:  "codeable_concept",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "already snake_case",
			input: "test_field",
			want:  "test_field",
		},
		{
			name:  "single uppercase letter",
			input: "A",
			want:  "a",
		},
		{
			name:  "XMLParser",
			input: "XMLParser",
			want:  "x_m_l_parser",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSnakeCase(tt.input)
			if got != tt.want {
				t.Errorf("ToSnakeCase(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestTitleCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple case",
			input: "test",
			want:  "Test",
		},
		{
			name:  "field",
			input: "field",
			want:  "Field",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "already capitalized",
			input: "Test",
			want:  "Test",
		},
		{
			name:  "single letter",
			input: "a",
			want:  "A",
		},
		{
			name:  "uppercase",
			input: "TEST",
			want:  "TEST",
		},
		{
			name:  "with numbers",
			input: "field123",
			want:  "Field123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TitleCase(tt.input)
			if got != tt.want {
				t.Errorf("TitleCase(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
