package v2

import (
	"encoding/json"
	"testing"
)

func TestCreateSecgroupResponse_Unmarshal_ToEntity(t *testing.T) {
	jsonData := `{
		"data": {
			"id": 1,
			"uuid": "sg-uuid-1",
			"createdAt": "2024-01-01T00:00:00Z",
			"deletedAt": null,
			"status": "ACTIVE",
			"secgroupId": 42,
			"secgroupName": "my-secgroup",
			"projectUuid": "proj-1",
			"description": "test sg",
			"updatedAt": null,
			"isSystem": false,
			"type": "custom"
		}
	}`

	var resp CreateSecgroupResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	sg := resp.ToEntitySecgroup()
	if sg.ID != "sg-uuid-1" {
		t.Fatalf("ID: got %q", sg.ID)
	}
	if sg.Name != "my-secgroup" {
		t.Fatalf("Name: got %q", sg.Name)
	}
	if sg.Description != "test sg" {
		t.Fatalf("Description: got %q", sg.Description)
	}
	if sg.Status != "ACTIVE" {
		t.Fatalf("Status: got %q", sg.Status)
	}
}

func TestGetSecgroupByIDResponse_Unmarshal_ToEntity(t *testing.T) {
	jsonData := `{
		"data": {
			"id": "sg-id-2",
			"name": "default",
			"description": "default security group",
			"status": "ACTIVE",
			"createdAt": "2024-06-01T10:00:00Z",
			"isSystem": true
		}
	}`

	var resp GetSecgroupByIDResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	sg := resp.ToEntitySecgroup()
	if sg.ID != "sg-id-2" {
		t.Fatalf("ID: got %q", sg.ID)
	}
	if sg.Name != "default" {
		t.Fatalf("Name: got %q", sg.Name)
	}
	if sg.Description != "default security group" {
		t.Fatalf("Description: got %q", sg.Description)
	}
}

func TestListSecgroupResponse_Unmarshal_ToEntity(t *testing.T) {
	jsonData := `{
		"listData": [
			{"id": "sg-1", "name": "sg-one", "description": "first", "status": "ACTIVE", "createdAt": "2024-01-01", "isSystem": false},
			{"id": "sg-2", "name": "sg-two", "description": "second", "status": "ACTIVE", "createdAt": "2024-01-02", "isSystem": true}
		],
		"page": 1,
		"pageSize": 10,
		"totalPage": 1,
		"totalItem": 2
	}`

	var resp ListSecgroupResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	lst := resp.ToListEntitySecgroups()
	if len(lst.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(lst.Items))
	}
	if lst.Items[0].ID != "sg-1" {
		t.Fatalf("first ID: got %q", lst.Items[0].ID)
	}
	if lst.Items[1].Name != "sg-two" {
		t.Fatalf("second Name: got %q", lst.Items[1].Name)
	}
}

func TestListSecgroupResponse_Empty(t *testing.T) {
	jsonData := `{"listData": [], "page": 1, "pageSize": 10, "totalPage": 0, "totalItem": 0}`
	var resp ListSecgroupResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	lst := resp.ToListEntitySecgroups()
	if len(lst.Items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(lst.Items))
	}
}
