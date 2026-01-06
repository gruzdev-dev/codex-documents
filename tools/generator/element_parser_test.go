package generator

import (
	"strings"
	"testing"
)

func TestProcessElements_RequiredFields(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		want     map[string]struct {
			fieldName    string
			hasOmitEmpty bool
		}
	}{
		{
			name: "required field min=1 should not have omitempty",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.requiredField",
					Path:  "TestResource.requiredField",
					Min:   1,
					Max:   "1",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Required field",
				},
			},
			want: map[string]struct {
				fieldName    string
				hasOmitEmpty bool
			}{
				"TestResource": {
					fieldName:    "RequiredField",
					hasOmitEmpty: false,
				},
			},
		},
		{
			name: "optional field min=0 should have omitempty",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.optionalField",
					Path:  "TestResource.optionalField",
					Min:   0,
					Max:   "1",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Optional field",
				},
			},
			want: map[string]struct {
				fieldName    string
				hasOmitEmpty bool
			}{
				"TestResource": {
					fieldName:    "OptionalField",
					hasOmitEmpty: true,
				},
			},
		},
		{
			name: "required array min=2 should not have omitempty",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.requiredArray",
					Path:  "TestResource.requiredArray",
					Min:   2,
					Max:   "*",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Required array",
				},
			},
			want: map[string]struct {
				fieldName    string
				hasOmitEmpty bool
			}{
				"TestResource": {
					fieldName:    "RequiredArray",
					hasOmitEmpty: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			for structName, expected := range tt.want {
				fields, ok := structs[structName]
				if !ok {
					t.Fatalf("struct %s not found", structName)
				}

				var found bool
				for _, field := range fields {
					if field.Name == expected.fieldName {
						found = true
						hasOmitEmpty := strings.Contains(field.JSONTag, "omitempty")
						if hasOmitEmpty != expected.hasOmitEmpty {
							t.Errorf("field %s: hasOmitEmpty = %v, want %v", field.Name, hasOmitEmpty, expected.hasOmitEmpty)
						}
						break
					}
				}

				if !found {
					t.Errorf("field %s not found in struct %s", expected.fieldName, structName)
				}
			}
		})
	}
}

func TestProcessElements_ChoiceTypes(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		want     []struct {
			fieldName string
			goType    string
			jsonTag   string
		}
	}{
		{
			name: "choice type [x] creates separate fields",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.value[x]",
					Path: "TestResource.value[x]",
					Min:  0,
					Max:  "1",
					Type: []ElementDataType{
						{Code: "string"},
						{Code: "integer"},
					},
					Short: "Choice type field",
				},
			},
			want: []struct {
				fieldName string
				goType    string
				jsonTag   string
			}{
				{
					fieldName: "ValueString",
					goType:    "*string",
					jsonTag:   "valueString",
				},
				{
					fieldName: "ValueInteger",
					goType:    "*int",
					jsonTag:   "valueInteger",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			fields, ok := structs["TestResource"]
			if !ok {
				t.Fatalf("struct TestResource not found")
			}

			if len(fields) < len(tt.want) {
				t.Fatalf("got %d fields, want at least %d", len(fields), len(tt.want))
			}

			for _, expected := range tt.want {
				var found bool
				for _, field := range fields {
					if field.Name == expected.fieldName {
						found = true
						if !strings.Contains(field.GoType, expected.goType) {
							t.Errorf("field %s: GoType = %s, want contains %s", field.Name, field.GoType, expected.goType)
						}
						if !strings.Contains(field.JSONTag, expected.jsonTag) {
							t.Errorf("field %s: JSONTag = %s, want contains %s", field.Name, field.JSONTag, expected.jsonTag)
						}
						break
					}
				}

				if !found {
					t.Errorf("field %s not found", expected.fieldName)
				}
			}
		})
	}
}

func TestProcessElements_ChoiceTypesJSONTags(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		wantJSON string
	}{
		{
			name: "choice type [x] JSON tags should be valueString, valueInteger (not just value)",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.value[x]",
					Path: "TestResource.value[x]",
					Min:  0,
					Max:  "1",
					Type: []ElementDataType{
						{Code: "string"},
						{Code: "integer"},
					},
					Short: "Choice type field",
				},
			},
			wantJSON: "valueString",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			fields, ok := structs["TestResource"]
			if !ok {
				t.Fatalf("struct TestResource not found")
			}

			var found bool
			for _, field := range fields {
				if strings.Contains(field.JSONTag, tt.wantJSON) {
					found = true
					if strings.Contains(field.JSONTag, "value,") || strings.Contains(field.JSONTag, "\"value\"") {
						t.Errorf("choice type JSON tag should be %s, not just 'value': %s", tt.wantJSON, field.JSONTag)
					}
					break
				}
			}

			if !found {
				t.Errorf("field with JSON tag %s not found", tt.wantJSON)
			}
		})
	}
}

func TestProcessElements_EmptyStructures(t *testing.T) {
	tests := []struct {
		name            string
		elements        []ElementDefinition
		wantEmptyStruct string
		shouldExist     bool
	}{
		{
			name: "empty BackboneElement should create empty structure",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.emptyNested",
					Path:  "TestResource.emptyNested",
					Min:   0,
					Max:   "*",
					Type:  []ElementDataType{{Code: "BackboneElement"}},
					Short: "Empty nested structure",
				},
			},
			wantEmptyStruct: "TestResourceEmptyNested",
			shouldExist:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			fields, exists := structs[tt.wantEmptyStruct]
			if tt.shouldExist && !exists {
				t.Errorf("empty structure %s should be created", tt.wantEmptyStruct)
			}
			if exists && len(fields) > 0 {
				t.Logf("Empty structure %s has %d fields (this may be a problem if structure is truly empty)", tt.wantEmptyStruct, len(fields))
			}
		})
	}
}

func TestProcessElements_NestedStructures(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		want     []string
	}{
		{
			name: "BackboneElement creates nested structure",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.nested",
					Path:  "TestResource.nested",
					Min:   0,
					Max:   "*",
					Type:  []ElementDataType{{Code: "BackboneElement"}},
					Short: "Nested structure",
				},
				{
					ID:    "TestResource.nested.field1",
					Path:  "TestResource.nested.field1",
					Min:   0,
					Max:   "1",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Field in nested",
				},
			},
			want: []string{"TestResource", "TestResourceNested"},
		},
		{
			name: "empty BackboneElement should still create structure",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.emptyNested",
					Path:  "TestResource.emptyNested",
					Min:   0,
					Max:   "*",
					Type:  []ElementDataType{{Code: "BackboneElement"}},
					Short: "Empty nested structure",
				},
			},
			want: []string{"TestResource", "TestResourceEmptyNested"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			for _, wantStruct := range tt.want {
				if _, ok := structs[wantStruct]; !ok {
					t.Errorf("struct %s not found, got structs: %v", wantStruct, getStructNames(structs))
				}
			}
		})
	}
}

func TestProcessElements_SkipElements(t *testing.T) {
	tests := []struct {
		name             string
		elements         []ElementDefinition
		wantSkip         []string
		shouldHaveStruct bool
	}{
		{
			name: "skip elements with ._ in ID",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource._implicitRules",
					Path: "TestResource.implicitRules",
					Min:  0,
					Max:  "1",
					Type: []ElementDataType{{Code: "uri"}},
				},
				{
					ID:    "TestResource.validField",
					Path:  "TestResource.validField",
					Min:   0,
					Max:   "1",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Valid field",
				},
			},
			wantSkip:         []string{"ImplicitRules"},
			shouldHaveStruct: true,
		},
		{
			name: "skip extension and modifierExtension",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.extension",
					Path: "TestResource.extension",
					Min:  0,
					Max:  "*",
					Type: []ElementDataType{{Code: "Extension"}},
				},
				{
					ID:   "TestResource.modifierExtension",
					Path: "TestResource.modifierExtension",
					Min:  0,
					Max:  "*",
					Type: []ElementDataType{{Code: "Extension"}},
				},
				{
					ID:    "TestResource.validField",
					Path:  "TestResource.validField",
					Min:   0,
					Max:   "1",
					Type:  []ElementDataType{{Code: "string"}},
					Short: "Valid field",
				},
			},
			wantSkip:         []string{"Extension"},
			shouldHaveStruct: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			if tt.shouldHaveStruct {
				fields, ok := structs["TestResource"]
				if !ok {
					t.Fatalf("struct TestResource not found")
				}

				for _, skipField := range tt.wantSkip {
					for _, field := range fields {
						if field.Name == skipField {
							t.Errorf("field %s should be skipped but was found", skipField)
						}
					}
				}
			}
		})
	}
}

func getStructNames(structs map[string][]FieldInfo) []string {
	names := make([]string, 0, len(structs))
	for name := range structs {
		names = append(names, name)
	}
	return names
}

func TestProcessElements_RequiredReferencePointer(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		want     struct {
			fieldName string
			goType    string
		}
	}{
		{
			name: "required Reference field (min=1) should have pointer",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.requiredRef",
					Path:  "TestResource.requiredRef",
					Min:   1,
					Max:   "1",
					Type:  []ElementDataType{{Code: "Reference"}},
					Short: "Required reference field",
				},
			},
			want: struct {
				fieldName string
				goType    string
			}{
				fieldName: "RequiredRef",
				goType:    "*Reference",
			},
		},
		{
			name: "required array Reference (min=1) should not have pointer (array itself)",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:    "TestResource.requiredRefArray",
					Path:  "TestResource.requiredRefArray",
					Min:   1,
					Max:   "*",
					Type:  []ElementDataType{{Code: "Reference"}},
					Short: "Required reference array",
				},
			},
			want: struct {
				fieldName string
				goType    string
			}{
				fieldName: "RequiredRefArray",
				goType:    "[]Reference",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			fields, ok := structs["TestResource"]
			if !ok {
				t.Fatalf("struct TestResource not found")
			}

			var found bool
			for _, field := range fields {
				if field.Name == tt.want.fieldName {
					found = true
					if field.GoType != tt.want.goType {
						t.Errorf("field %s: GoType = %s, want %s", field.Name, field.GoType, tt.want.goType)
					}
					break
				}
			}

			if !found {
				t.Errorf("field %s not found in struct TestResource", tt.want.fieldName)
			}
		})
	}
}

func TestProcessElements_ContentReference(t *testing.T) {
	tests := []struct {
		name     string
		elements []ElementDefinition
		want     struct {
			structName string
			fieldName  string
			goType     string
		}
	}{
		{
			name: "contentReference should use referenced type",
			elements: []ElementDefinition{
				{
					ID:   "TestResource",
					Path: "TestResource",
					Min:  0,
					Max:  "*",
				},
				{
					ID:   "TestResource.referenceRange",
					Path: "TestResource.referenceRange",
					Min:  0,
					Max:  "*",
					Type: []ElementDataType{{Code: "BackboneElement"}},
				},
				{
					ID:   "TestResource.referenceRange.low",
					Path: "TestResource.referenceRange.low",
					Min:  0,
					Max:  "1",
					Type: []ElementDataType{{Code: "Quantity"}},
				},
				{
					ID:               "TestResource.component",
					Path:             "TestResource.component",
					Min:              0,
					Max:              "*",
					Type:             []ElementDataType{{Code: "BackboneElement"}},
					ContentReference: "",
				},
				{
					ID:               "TestResource.component.referenceRange",
					Path:             "TestResource.component.referenceRange",
					Min:              0,
					Max:              "*",
					ContentReference: "#TestResource.referenceRange",
				},
			},
			want: struct {
				structName string
				fieldName  string
				goType     string
			}{
				structName: "TestResourceComponent",
				fieldName:  "ReferenceRange",
				goType:     "[]TestResourceReferenceRange",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			structs := g.ProcessElements("TestResource", tt.elements)

			fields, ok := structs[tt.want.structName]
			if !ok {
				t.Fatalf("struct %s not found. Available structs: %v", tt.want.structName, getStructNames(structs))
			}

			var found bool
			for _, field := range fields {
				if field.Name == tt.want.fieldName {
					found = true
					if field.GoType != tt.want.goType {
						t.Errorf("field %s: GoType = %s, want %s", field.Name, field.GoType, tt.want.goType)
					}
					break
				}
			}

			if !found {
				t.Errorf("field %s not found in struct %s. Available fields: %v", tt.want.fieldName, tt.want.structName, getFieldNames(fields))
			}
		})
	}
}

func getFieldNames(fields []FieldInfo) []string {
	names := make([]string, 0, len(fields))
	for _, f := range fields {
		names = append(names, f.Name)
	}
	return names
}
