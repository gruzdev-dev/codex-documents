package generator

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestWriteStruct_JSONTags(t *testing.T) {
	tests := []struct {
		name     string
		fields   []FieldInfo
		wantJSON map[string]string
	}{
		{
			name: "required field without omitempty",
			fields: []FieldInfo{
				{
					Name:    "RequiredField",
					GoType:  "string",
					JSONTag: "`json:\"requiredField\"`",
					Comment: "Required field",
				},
			},
			wantJSON: map[string]string{
				"RequiredField": "requiredField",
			},
		},
		{
			name: "optional field with omitempty",
			fields: []FieldInfo{
				{
					Name:    "OptionalField",
					GoType:  "*string",
					JSONTag: "`json:\"optionalField,omitempty\"`",
					Comment: "Optional field",
				},
			},
			wantJSON: map[string]string{
				"OptionalField": "optionalField,omitempty",
			},
		},
		{
			name: "array with omitempty",
			fields: []FieldInfo{
				{
					Name:    "ArrayField",
					GoType:  "[]string",
					JSONTag: "`json:\"arrayField,omitempty\"`",
					Comment: "Array field",
				},
			},
			wantJSON: map[string]string{
				"ArrayField": "arrayField,omitempty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			var buf bytes.Buffer
			g.writeStruct(&buf, "TestStruct", "", tt.fields)

			output := buf.String()
			for fieldName, wantTag := range tt.wantJSON {
				if !strings.Contains(output, fieldName) {
					t.Errorf("field %s not found in output", fieldName)
				}
				if !strings.Contains(output, wantTag) {
					t.Errorf("JSON tag %s not found for field %s, output: %s", wantTag, fieldName, output)
				}
			}
		})
	}
}

func TestWriteStruct_BSONTags(t *testing.T) {
	tests := []struct {
		name     string
		fields   []FieldInfo
		wantBSON map[string]string
	}{
		{
			name: "BSON tags should be present",
			fields: []FieldInfo{
				{
					Name:    "TestField",
					GoType:  "string",
					JSONTag: "`json:\"testField\"`",
					BSONTag: "`bson:\"test_field\"`",
					Comment: "Test field",
				},
			},
			wantBSON: map[string]string{
				"TestField": "test_field",
			},
		},
		{
			name: "BSON tags with omitempty",
			fields: []FieldInfo{
				{
					Name:    "OptionalField",
					GoType:  "*string",
					JSONTag: "`json:\"optionalField,omitempty\"`",
					BSONTag: "`bson:\"optional_field,omitempty\"`",
					Comment: "Optional field",
				},
			},
			wantBSON: map[string]string{
				"OptionalField": "optional_field,omitempty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			var buf bytes.Buffer
			g.writeStruct(&buf, "TestStruct", "", tt.fields)

			output := buf.String()
			for fieldName, wantBSON := range tt.wantBSON {
				if !strings.Contains(output, fieldName) {
					t.Errorf("field %s not found in output", fieldName)
				}
				bsonTagPattern := fmt.Sprintf("bson:\"%s\"", wantBSON)
				if !strings.Contains(output, bsonTagPattern) {
					t.Errorf("BSON tag for field %s not found, want: %s, got output: %s", fieldName, wantBSON, output)
				}
			}
		})
	}
}

func TestWriteStruct_EmptyStructures(t *testing.T) {
	tests := []struct {
		name        string
		fields      []FieldInfo
		shouldExist bool
	}{
		{
			name:        "empty structure should not be generated",
			fields:      []FieldInfo{},
			shouldExist: false,
		},
		{
			name: "structure with fields should be generated",
			fields: []FieldInfo{
				{
					Name:    "Field1",
					GoType:  "string",
					JSONTag: "`json:\"field1\"`",
				},
			},
			shouldExist: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			var buf bytes.Buffer
			g.writeStruct(&buf, "TestStruct", "Test description", tt.fields)

			output := buf.String()
			hasStruct := strings.Contains(output, "type TestStruct struct")
			if tt.shouldExist && !hasStruct {
				t.Errorf("struct TestStruct should be generated but was not found")
			}
			if !tt.shouldExist && hasStruct {
				t.Errorf("struct TestStruct should not be generated but was found")
			}
		})
	}
}

func TestWriteStruct_NestedStructures(t *testing.T) {
	tests := []struct {
		name   string
		fields []FieldInfo
		want   []string
	}{
		{
			name: "nested structure names should be correct",
			fields: []FieldInfo{
				{
					Name:    "Nested",
					GoType:  "TestResourceNested",
					JSONTag: "`json:\"nested\"`",
				},
			},
			want: []string{"TestResourceNested"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			var buf bytes.Buffer
			g.writeStruct(&buf, "TestResource", "", tt.fields)

			output := buf.String()
			for _, wantType := range tt.want {
				if !strings.Contains(output, wantType) {
					t.Errorf("type %s not found in output", wantType)
				}
			}
		})
	}
}

func TestWriteResource_Imports(t *testing.T) {
	tests := []struct {
		name        string
		elements    []ElementDefinition
		wantImports []string
	}{
		{
			name: "json.RawMessage should add encoding/json import",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.contained",
					Path: "TestResource.contained",
					Min:  0,
					Max:  "*",
					Type: []ElementDataType{
						{Code: "Resource"},
					},
					Short: "Contained resources",
				},
			},
			wantImports: []string{"encoding/json"},
		},
		{
			name: "no json.RawMessage should not add encoding/json import",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.stringField",
					Path: "TestResource.stringField",
					Min:  0,
					Max:  "1",
					Type: []ElementDataType{
						{Code: "string"},
					},
					Short: "String field",
				},
			},
			wantImports: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			def := StructureDefinition{
				Name:        "TestResource",
				Description: "Test resource",
				Snapshot: Snapshot{
					Element: tt.elements,
				},
			}

			g := NewGenerator("", "", nil)
			structMap := g.ProcessElements("TestResource", def.Snapshot.Element)

			var buf bytes.Buffer
			fmt.Fprintf(&buf, "package models\n\n")
			needsJSON := false
			for _, fields := range structMap {
				for _, f := range fields {
					if strings.Contains(f.GoType, "json.RawMessage") {
						needsJSON = true
						break
					}
				}
				if needsJSON {
					break
				}
			}

			if needsJSON {
				fmt.Fprintf(&buf, "import \"encoding/json\"\n\n")
			}

			output := buf.String()
			for _, wantImport := range tt.wantImports {
				if !strings.Contains(output, wantImport) {
					t.Errorf("import %s not found in output", wantImport)
				}
			}

			for _, notWantImport := range []string{"encoding/json"} {
				if len(tt.wantImports) == 0 && strings.Contains(output, notWantImport) {
					t.Errorf("import %s should not be present", notWantImport)
				}
			}
		})
	}
}
