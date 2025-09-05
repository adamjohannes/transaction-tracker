package sub_category

import "testing"

// TestNew validates the logic for creating a new sub-category.
func TestNew(t *testing.T) {
	testCases := []struct {
		name             string
		inputParentID    int8
		inputName        string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:          "Successful case with valid data",
			inputParentID: 1,
			inputName:     "Smartphones",
			expectError:   false,
		},
		{
			name:          "Successful case with zero parent ID",
			inputParentID: 0,
			inputName:     "Laptops",
			expectError:   false,
		},
		{
			name:             "Error case with negative parent ID",
			inputParentID:    -1,
			inputName:        "Tablets",
			expectError:      true,
			expectedErrorMsg: "parent id must not be negative",
		},
		{
			name:             "Error case with empty name",
			inputParentID:    2,
			inputName:        "",
			expectError:      true,
			expectedErrorMsg: "name is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subCategory, err := New(tc.inputParentID, tc.inputName)

			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error but got none")
				}
				if err.Error() != tc.expectedErrorMsg {
					t.Errorf("expected error '%s' but got '%s'", tc.expectedErrorMsg, err.Error())
				}
				if subCategory != nil {
					t.Error("expected subCategory to be nil on error")
				}
			} else { // No error expected
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				if subCategory == nil {
					t.Fatal("expected a subCategory instance but got nil")
				}
				if subCategory.ID != -1 {
					t.Errorf("expected ID to be -1 but got %d", subCategory.ID)
				}
				if subCategory.ParentID != tc.inputParentID {
					t.Errorf("expected ParentID to be '%d' but got '%d'", tc.inputParentID, subCategory.ParentID)
				}
				if subCategory.Name != tc.inputName {
					t.Errorf("expected Name to be '%s' but got '%s'", tc.inputName, subCategory.Name)
				}
			}
		})
	}
}

// TestBuild validates the logic for reconstructing a sub-category from existing data.
func TestBuild(t *testing.T) {
	// Since Build has no complex logic, a direct test is sufficient.
	id := int8(101)
	parentID := int8(20)
	name := "Rebuilt SubCategory"

	subCategory := Build(id, parentID, name)

	if subCategory == nil {
		t.Fatal("Build() returned a nil subCategory")
	}

	if subCategory.ID != id {
		t.Errorf("expected ID to be %d, got %d", id, subCategory.ID)
	}
	if subCategory.ParentID != parentID {
		t.Errorf("expected ParentID to be %d, got %d", parentID, subCategory.ParentID)
	}
	if subCategory.Name != name {
		t.Errorf("expected Name to be '%s', got '%s'", name, subCategory.Name)
	}
}
