package common

import (
	"testing"
)

// ---------------------------------------------------------------------------
// Ptr
// ---------------------------------------------------------------------------

func TestPtr(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"int", func(t *testing.T) {
			p := Ptr(42)
			if *p != 42 {
				t.Fatalf("got %d, want 42", *p)
			}
		}},
		{"string", func(t *testing.T) {
			p := Ptr("hello")
			if *p != "hello" {
				t.Fatalf("got %q, want %q", *p, "hello")
			}
		}},
		{"bool", func(t *testing.T) {
			p := Ptr(true)
			if *p != true {
				t.Fatal("got false, want true")
			}
		}},
		{"struct", func(t *testing.T) {
			type S struct{ X int }
			p := Ptr(S{X: 7})
			if p.X != 7 {
				t.Fatalf("got %d, want 7", p.X)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.fn)
	}
}

// ---------------------------------------------------------------------------
// StructToMap
// ---------------------------------------------------------------------------

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		wantKeys []string
		wantVals map[string]any
	}{
		{
			name: "struct with json tags",
			input: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{"Alice", 30},
			wantKeys: []string{"name", "age"},
			wantVals: map[string]any{"name": "Alice", "age": float64(30)},
		},
		{
			name:     "empty struct",
			input:    struct{}{},
			wantKeys: nil,
			wantVals: map[string]any{},
		},
		{
			name: "omitempty skips zero values",
			input: struct {
				A string `json:"a,omitempty"`
				B string `json:"b"`
			}{A: "", B: ""},
			wantKeys: []string{"b"},
			wantVals: map[string]any{"b": ""},
		},
		{
			name:     "nil input",
			input:    nil,
			wantKeys: nil,
			wantVals: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := StructToMap(tt.input)

			if tt.wantVals == nil {
				if m != nil {
					t.Fatalf("expected nil map, got %v", m)
				}
				return
			}

			for _, k := range tt.wantKeys {
				if _, ok := m[k]; !ok {
					t.Errorf("missing key %q in map", k)
				}
			}

			for k, want := range tt.wantVals {
				got, ok := m[k]
				if !ok {
					t.Errorf("missing key %q", k)
					continue
				}
				if got != want {
					t.Errorf("key %q: got %v, want %v", k, got, want)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Project
// ---------------------------------------------------------------------------

func TestProject(t *testing.T) {
	var p Project
	if p.ID != "" {
		t.Fatal("zero-value project should return empty string")
	}

	p.ID = "proj-123"
	if p.ID != "proj-123" {
		t.Fatalf("got %q, want %q", p.ID, "proj-123")
	}
}

// ---------------------------------------------------------------------------
// Paging
// ---------------------------------------------------------------------------

func TestPaging(t *testing.T) {
	var p Paging
	if p.Page != 0 || p.Size != 0 {
		t.Fatal("zero-value paging should return 0")
	}

	p.Page = 3
	p.Size = 25
	if p.Page != 3 {
		t.Fatalf("page: got %d, want 3", p.Page)
	}
	if p.Size != 25 {
		t.Fatalf("size: got %d, want 25", p.Size)
	}
}

// ---------------------------------------------------------------------------
// PortalUser
// ---------------------------------------------------------------------------

func TestPortalUser(t *testing.T) {
	var pu PortalUser
	if pu.ID != "" {
		t.Fatal("zero-value should return empty string")
	}

	pu.ID = "user-456"
	if pu.ID != "user-456" {
		t.Fatalf("got %q, want %q", pu.ID, "user-456")
	}
}

// ---------------------------------------------------------------------------
// Tag
// ---------------------------------------------------------------------------

func TestTag(t *testing.T) {
	tag := Tag{Key: "env", Value: "prod", IsEdited: true}
	if tag.Key != "env" || tag.Value != "prod" || !tag.IsEdited {
		t.Fatalf("unexpected tag: %+v", tag)
	}
}

// ---------------------------------------------------------------------------
// Zone constants
// ---------------------------------------------------------------------------

func TestZoneConstants(t *testing.T) {
	zones := []Zone{HCM_03_1A_ZONE, HCM_03_1B_ZONE, HCM_03_1C_ZONE, HCM_03_BKK_01_ZONE, HAN_01_1A_ZONE}
	for _, z := range zones {
		if z == "" {
			t.Fatal("zone constant should not be empty")
		}
	}
}
