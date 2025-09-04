package category

import "testing"

// TestNew tests the New function, which is for creating a new category entity.
func TestNew(t *testing.T) {
	// Table-driven tests for various scenarios.
	testCases := []struct {
		name             string
		inputName        string
		inputDescription string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:             "Successful case with all fields",
			inputName:        "Electronics",
			inputDescription: "All kinds of electronic gadgets.",
			expectError:      false,
		},
		{
			name:             "Successful case with empty description",
			inputName:        "Books",
			inputDescription: "", // Description is optional
			expectError:      false,
		},
		{
			name:             "Error case with empty name",
			inputName:        "", // Name is required
			inputDescription: "This should fail",
			expectError:      true,
			expectedErrorMsg: "name is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			category, err := New(tc.inputName, tc.inputDescription)

			if tc.expectError {
				// We expected an error.
				if err == nil {
					t.Fatal("expected an error but got none")
				}
				if err.Error() != tc.expectedErrorMsg {
					t.Errorf("expected error message '%s' but got '%s'", tc.expectedErrorMsg, err.Error())
				}
				if category != nil {
					t.Error("expected category to be nil on error")
				}
			} else {
				// We did NOT expect an error.
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				if category == nil {
					t.Fatal("expected a category instance but got nil")
				}
				if category.Id != -1 {
					t.Errorf("expected Id to be initialized to -1, but got %d", category.Id)
				}
				if category.Name != tc.inputName {
					t.Errorf("expected name to be '%s', but got '%s'", tc.inputName, category.Name)
				}
				if category.Description != tc.inputDescription {
					t.Errorf("expected description to be '%s', but got '%s'", tc.inputDescription, category.Description)
				}
			}
		})
	}
}

// TestBuild tests the Build function, which is for reconstructing a category from storage.
func TestBuild(t *testing.T) {
	// Since Build has no validation, a single, straightforward test is sufficient.
	id := int8(42)
	name := "Rehydrated Category"
	description := "This category was built from existing data."

	category, err := Build(id, name, description)

	if err != nil {
		t.Fatalf("Build() returned an unexpected error: %v", err)
	}
	if category == nil {
		t.Fatal("Build() returned a nil category")
	}

	// Verify that all fields match the input exactly.
	if category.Id != id {
		t.Errorf("expected Id to be %d, but got %d", id, category.Id)
	}
	if category.Name != name {
		t.Errorf("expected Name to be '%s', but got '%s'", name, category.Name)
	}
	if category.Description != description {
		t.Errorf("expected Description to be '%s', but got '%s'", description, category.Description)
	}
}
