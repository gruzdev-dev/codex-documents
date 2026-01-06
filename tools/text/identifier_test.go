package text

import "testing"

func TestIsValidGoIdentifier(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid identifier",
			input: "Patient",
			want:  true,
		},
		{
			name:  "valid identifier with underscore",
			input: "Test_Resource",
			want:  true,
		},
		{
			name:  "valid identifier starting with underscore",
			input: "_private",
			want:  true,
		},
		{
			name:  "invalid identifier starting with number",
			input: "123Invalid",
			want:  false,
		},
		{
			name:  "invalid identifier with dash",
			input: "test-resource",
			want:  false,
		},
		{
			name:  "invalid identifier with space",
			input: "test resource",
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "valid identifier with numbers",
			input: "Field123",
			want:  true,
		},
		{
			name:  "valid single letter",
			input: "A",
			want:  true,
		},
		{
			name:  "valid single underscore",
			input: "_",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidGoIdentifier(tt.input)
			if got != tt.want {
				t.Errorf("IsValidGoIdentifier(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
