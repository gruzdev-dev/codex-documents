package generator

import (
	"testing"
)

func TestMapGoType_PrimitiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		element  ElementDefinition
		want     string
		wantUsed bool
	}{
		{
			name: "string type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "string"}},
			},
			want:     "string",
			wantUsed: false,
		},
		{
			name: "boolean type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "boolean"}},
			},
			want:     "bool",
			wantUsed: false,
		},
		{
			name: "integer type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "integer"}},
			},
			want:     "int",
			wantUsed: false,
		},
		{
			name: "integer64 type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "integer64"}},
			},
			want:     "int64",
			wantUsed: false,
		},
		{
			name: "decimal type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "decimal"}},
			},
			want:     "float64",
			wantUsed: false,
		},
		{
			name: "date type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "date"}},
			},
			want:     "string",
			wantUsed: false,
		},
		{
			name: "dateTime type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "dateTime"}},
			},
			want:     "string",
			wantUsed: false,
		},
		{
			name: "uri type",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "uri"}},
			},
			want:     "string",
			wantUsed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			got := g.mapGoType(tt.element)

			if got != tt.want {
				t.Errorf("mapGoType() = %v, want %v", got, tt.want)
			}

			if tt.wantUsed {
				if !g.usedTypes[tt.want] {
					t.Errorf("type %s should be added to usedTypes", tt.want)
				}
			} else {
				if g.usedTypes[tt.want] {
					t.Errorf("primitive type %s should not be added to usedTypes", tt.want)
				}
			}
		})
	}
}

func TestMapGoType_ReferenceTypes(t *testing.T) {
	tests := []struct {
		name     string
		element  ElementDefinition
		want     string
		wantUsed bool
	}{
		{
			name: "Reference without targetProfile",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "Reference"}},
			},
			want:     "Reference",
			wantUsed: true,
		},
		{
			name: "Reference with targetProfile (should still be Reference)",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{
					{
						Code: "Reference",
					},
				},
			},
			want:     "Reference",
			wantUsed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			got := g.mapGoType(tt.element)

			if got != tt.want {
				t.Errorf("mapGoType() = %v, want %v", got, tt.want)
			}

			if tt.wantUsed {
				if !g.usedTypes[tt.want] {
					t.Errorf("type %s should be added to usedTypes", tt.want)
				}
			}
		})
	}
}

func TestMapGoType_ReferenceTargetProfile(t *testing.T) {
	tests := []struct {
		name        string
		element     ElementDefinition
		wantType    string
		shouldNotGenerateTargets bool
	}{
		{
			name: "Reference with targetProfile should not generate target models",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{
					{
						Code: "Reference",
					},
				},
			},
			wantType: "Reference",
			shouldNotGenerateTargets: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			got := g.mapGoType(tt.element)

			if got != tt.wantType {
				t.Errorf("mapGoType() = %v, want %v", got, tt.wantType)
			}

			if tt.shouldNotGenerateTargets {
				if g.usedTypes["Patient"] || g.usedTypes["Organization"] {
					t.Errorf("targetProfile models should not be added to usedTypes automatically")
				}
			}
		})
	}
}

func TestMapGoType_BackboneElement(t *testing.T) {
	tests := []struct {
		name    string
		element ElementDefinition
		want    string
	}{
		{
			name: "BackboneElement creates nested type name",
			element: ElementDefinition{
				Path: "TestResource.nested",
				Type: []ElementDataType{{Code: "BackboneElement"}},
			},
			want: "TestResourceNested",
		},
		{
			name: "Element creates nested type name",
			element: ElementDefinition{
				Path: "TestResource.element",
				Type: []ElementDataType{{Code: "Element"}},
			},
			want: "TestResourceElement",
		},
		{
			name: "deeply nested BackboneElement",
			element: ElementDefinition{
				Path: "TestResource.level1.level2",
				Type: []ElementDataType{{Code: "BackboneElement"}},
			},
			want: "TestResourceLevel1Level2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			got := g.mapGoType(tt.element)

			if got != tt.want {
				t.Errorf("mapGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapGoType_MultipleTypes(t *testing.T) {
	tests := []struct {
		name    string
		element ElementDefinition
		want    string
	}{
		{
			name: "multiple types returns any",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{
					{Code: "string"},
					{Code: "integer"},
				},
			},
			want: "any",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			got := g.mapGoType(tt.element)

			if got != tt.want {
				t.Errorf("mapGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapGoType_UsedTypesTracking(t *testing.T) {
	tests := []struct {
		name     string
		element  ElementDefinition
		wantType string
		wantUsed bool
	}{
		{
			name: "non-primitive type should be added to usedTypes",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "CustomType"}},
			},
			wantType: "CustomType",
			wantUsed: true,
		},
		{
			name: "primitive type should not be added to usedTypes",
			element: ElementDefinition{
				Path: "TestResource.field",
				Type: []ElementDataType{{Code: "string"}},
			},
			wantType: "string",
			wantUsed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator("", "", nil)
			_ = g.mapGoType(tt.element)

			if tt.wantUsed {
				if !g.usedTypes[tt.wantType] {
					t.Errorf("type %s should be added to usedTypes", tt.wantType)
				}
			} else {
				if g.usedTypes[tt.wantType] {
					t.Errorf("primitive type %s should not be added to usedTypes", tt.wantType)
				}
			}
		})
	}
}

func TestMapGoType_UsedTypesIsolation(t *testing.T) {
	t.Run("usedTypes should be isolated between calls", func(t *testing.T) {
		g1 := NewGenerator("", "", nil)
		_ = g1.mapGoType(ElementDefinition{
			Path: "TestResource.field1",
			Type: []ElementDataType{{Code: "TypeA"}},
		})

		g2 := NewGenerator("", "", nil)
		_ = g2.mapGoType(ElementDefinition{
			Path: "TestResource.field2",
			Type: []ElementDataType{{Code: "TypeB"}},
		})

		if g2.usedTypes["TypeA"] {
			t.Errorf("usedTypes should be isolated - TypeA should not be present in second generator")
		}
		if !g2.usedTypes["TypeB"] {
			t.Errorf("usedTypes should contain TypeB after second call")
		}
	})
}
