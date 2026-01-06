package generator

import (
	"encoding/json"
	"testing"
)

func TestElementDefinition_ParseBinding(t *testing.T) {
	jsonData := `{
		"id": "TestElement.binding",
		"path": "TestElement.binding",
		"min": 0,
		"max": "1",
		"binding": {
			"strength": "extensible",
			"description": "Security Labels from the Healthcare Privacy and Security Classification System.",
			"valueSet": "http://hl7.org/fhir/ValueSet/security-labels",
			"extension": [
				{
					"url": "http://hl7.org/fhir/StructureDefinition/elementdefinition-bindingName",
					"valueString": "SecurityLabels"
				}
			]
		}
	}`

	var el ElementDefinition
	if err := json.Unmarshal([]byte(jsonData), &el); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if el.Binding == nil {
		t.Fatal("Binding should not be nil")
	}

	if el.Binding.Strength != "extensible" {
		t.Errorf("Binding.Strength = %v, want extensible", el.Binding.Strength)
	}

	if el.Binding.Description != "Security Labels from the Healthcare Privacy and Security Classification System." {
		t.Errorf("Binding.Description = %v, want Security Labels from the Healthcare Privacy and Security Classification System.", el.Binding.Description)
	}

	if el.Binding.ValueSet != "http://hl7.org/fhir/ValueSet/security-labels" {
		t.Errorf("Binding.ValueSet = %v, want http://hl7.org/fhir/ValueSet/security-labels", el.Binding.ValueSet)
	}

	if len(el.Binding.Extension) == 0 {
		t.Error("Binding.Extension should not be empty")
	}
}

func TestElementDefinition_ParseConstraint(t *testing.T) {
	jsonData := `{
		"id": "TestElement.constraint",
		"path": "TestElement.constraint",
		"min": 0,
		"max": "1",
		"constraint": [
			{
				"key": "ele-1",
				"severity": "error",
				"human": "All FHIR elements must have a @value or children",
				"expression": "hasValue() or (children().count() > id.count())",
				"source": "http://hl7.org/fhir/StructureDefinition/Element"
			},
			{
				"key": "docRef-1",
				"severity": "warning",
				"human": "facilityType SHALL only be present if context is not an encounter",
				"expression": "facilityType.empty() or context.where(resolve() is Encounter).empty()"
			}
		]
	}`

	var el ElementDefinition
	if err := json.Unmarshal([]byte(jsonData), &el); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if len(el.Constraint) != 2 {
		t.Fatalf("len(el.Constraint) = %v, want 2", len(el.Constraint))
	}

	if el.Constraint[0].Key != "ele-1" {
		t.Errorf("Constraint[0].Key = %v, want ele-1", el.Constraint[0].Key)
	}

	if el.Constraint[0].Severity != "error" {
		t.Errorf("Constraint[0].Severity = %v, want error", el.Constraint[0].Severity)
	}

	if el.Constraint[0].Human != "All FHIR elements must have a @value or children" {
		t.Errorf("Constraint[0].Human = %v, want All FHIR elements must have a @value or children", el.Constraint[0].Human)
	}

	if el.Constraint[0].Expression != "hasValue() or (children().count() > id.count())" {
		t.Errorf("Constraint[0].Expression = %v, want hasValue() or (children().count() > id.count())", el.Constraint[0].Expression)
	}

	if el.Constraint[0].Source != "http://hl7.org/fhir/StructureDefinition/Element" {
		t.Errorf("Constraint[0].Source = %v, want http://hl7.org/fhir/StructureDefinition/Element", el.Constraint[0].Source)
	}

	if el.Constraint[1].Key != "docRef-1" {
		t.Errorf("Constraint[1].Key = %v, want docRef-1", el.Constraint[1].Key)
	}

	if el.Constraint[1].Severity != "warning" {
		t.Errorf("Constraint[1].Severity = %v, want warning", el.Constraint[1].Severity)
	}

	if el.Constraint[1].Source != "" {
		t.Errorf("Constraint[1].Source = %v, want empty string", el.Constraint[1].Source)
	}
}

func TestElementDefinition_ParseMaxLengthPatternFixed(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		check    func(*testing.T, ElementDefinition)
	}{
		{
			name: "maxLength",
			jsonData: `{
				"id": "TestElement.maxLength",
				"path": "TestElement.maxLength",
				"min": 0,
				"max": "1",
				"maxLength": 1048576
			}`,
			check: func(t *testing.T, el ElementDefinition) {
				if el.MaxLength == nil {
					t.Fatal("MaxLength should not be nil")
				}
				if *el.MaxLength != 1048576 {
					t.Errorf("MaxLength = %v, want 1048576", *el.MaxLength)
				}
			},
		},
		{
			name: "pattern",
			jsonData: `{
				"id": "TestElement.pattern",
				"path": "TestElement.pattern",
				"min": 0,
				"max": "1",
				"pattern": "[A-Z][a-z]+"
			}`,
			check: func(t *testing.T, el ElementDefinition) {
				if el.Pattern != "[A-Z][a-z]+" {
					t.Errorf("Pattern = %v, want [A-Z][a-z]+", el.Pattern)
				}
			},
		},
		{
			name: "fixed string",
			jsonData: `{
				"id": "TestElement.fixed",
				"path": "TestElement.fixed",
				"min": 0,
				"max": "1",
				"fixed": "test-value"
			}`,
			check: func(t *testing.T, el ElementDefinition) {
				if el.Fixed == nil {
					t.Fatal("Fixed should not be nil")
				}
				fixedStr, ok := el.Fixed.(string)
				if !ok {
					t.Fatalf("Fixed should be string, got %T", el.Fixed)
				}
				if fixedStr != "test-value" {
					t.Errorf("Fixed = %v, want test-value", fixedStr)
				}
			},
		},
		{
			name: "fixed number",
			jsonData: `{
				"id": "TestElement.fixed",
				"path": "TestElement.fixed",
				"min": 0,
				"max": "1",
				"fixed": 42
			}`,
			check: func(t *testing.T, el ElementDefinition) {
				if el.Fixed == nil {
					t.Fatal("Fixed should not be nil")
				}
				fixedNum, ok := el.Fixed.(float64)
				if !ok {
					t.Fatalf("Fixed should be number, got %T", el.Fixed)
				}
				if fixedNum != 42 {
					t.Errorf("Fixed = %v, want 42", fixedNum)
				}
			},
		},
		{
			name: "fixed boolean",
			jsonData: `{
				"id": "TestElement.fixed",
				"path": "TestElement.fixed",
				"min": 0,
				"max": "1",
				"fixed": true
			}`,
			check: func(t *testing.T, el ElementDefinition) {
				if el.Fixed == nil {
					t.Fatal("Fixed should not be nil")
				}
				fixedBool, ok := el.Fixed.(bool)
				if !ok {
					t.Fatalf("Fixed should be boolean, got %T", el.Fixed)
				}
				if fixedBool != true {
					t.Errorf("Fixed = %v, want true", fixedBool)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var el ElementDefinition
			if err := json.Unmarshal([]byte(tt.jsonData), &el); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			tt.check(t, el)
		})
	}
}

func TestElementDefinition_ParseFull(t *testing.T) {
	jsonData := `{
		"id": "TestElement.full",
		"path": "TestElement.full",
		"min": 0,
		"max": "1",
		"binding": {
			"strength": "required",
			"description": "Test binding description",
			"valueSet": "http://hl7.org/fhir/ValueSet/test"
		},
		"constraint": [
			{
				"key": "test-1",
				"severity": "error",
				"human": "Test constraint",
				"expression": "test()",
				"source": "http://test.com"
			}
		],
		"maxLength": 100,
		"pattern": "[0-9]+",
		"fixed": "fixed-value"
	}`

	var el ElementDefinition
	if err := json.Unmarshal([]byte(jsonData), &el); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if el.Binding == nil {
		t.Fatal("Binding should not be nil")
	}
	if el.Binding.Strength != "required" {
		t.Errorf("Binding.Strength = %v, want required", el.Binding.Strength)
	}

	if len(el.Constraint) != 1 {
		t.Fatalf("len(el.Constraint) = %v, want 1", len(el.Constraint))
	}
	if el.Constraint[0].Key != "test-1" {
		t.Errorf("Constraint[0].Key = %v, want test-1", el.Constraint[0].Key)
	}

	if el.MaxLength == nil {
		t.Fatal("MaxLength should not be nil")
	}
	if *el.MaxLength != 100 {
		t.Errorf("MaxLength = %v, want 100", *el.MaxLength)
	}

	if el.Pattern != "[0-9]+" {
		t.Errorf("Pattern = %v, want [0-9]+", el.Pattern)
	}

	if el.Fixed == nil {
		t.Fatal("Fixed should not be nil")
	}
	fixedStr, ok := el.Fixed.(string)
	if !ok || fixedStr != "fixed-value" {
		t.Errorf("Fixed = %v, want fixed-value", el.Fixed)
	}
}

func TestElementDefinition_ParseWithoutOptionalFields(t *testing.T) {
	jsonData := `{
		"id": "TestElement.minimal",
		"path": "TestElement.minimal",
		"min": 0,
		"max": "1"
	}`

	var el ElementDefinition
	if err := json.Unmarshal([]byte(jsonData), &el); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if el.Binding != nil {
		t.Error("Binding should be nil")
	}

	if len(el.Constraint) != 0 {
		t.Errorf("len(el.Constraint) = %v, want 0", len(el.Constraint))
	}

	if el.MaxLength != nil {
		t.Error("MaxLength should be nil")
	}

	if el.Pattern != "" {
		t.Errorf("Pattern = %v, want empty string", el.Pattern)
	}

	if el.Fixed != nil {
		t.Error("Fixed should be nil")
	}
}
