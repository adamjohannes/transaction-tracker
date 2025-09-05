package transaction_type

import "testing"

// TestNew validates the creation of a new transaction type.
func TestNew(t *testing.T) {
	testCases := []struct {
		name             string
		inputName        string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:        "Successful case with a valid name",
			inputName:   "Credit",
			expectError: false,
		},
		{
			name:             "Error case with an empty name",
			inputName:        "",
			expectError:      true,
			expectedErrorMsg: "name is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transactionType, err := New(tc.inputName)

			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error but got none")
				}
				if err.Error() != tc.expectedErrorMsg {
					t.Errorf("expected error message '%s' but got '%s'", tc.expectedErrorMsg, err.Error())
				}
				if transactionType != nil {
					t.Error("expected transactionType to be nil on error")
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				if transactionType == nil {
					t.Fatal("expected a transactionType instance but got nil")
				}
				if transactionType.ID != -1 {
					t.Errorf("expected ID to be initialized to -1, but got %d", transactionType.ID)
				}
				if transactionType.Name != tc.inputName {
					t.Errorf("expected name to be '%s', but got '%s'", tc.inputName, transactionType.Name)
				}
			}
		})
	}
}

// TestBuild validates the reconstruction of a transaction type from existing data.
func TestBuild(t *testing.T) {
	id := int8(1)
	name := "Debit"

	transactionType := Build(id, name)

	if transactionType == nil {
		t.Fatal("Build() returned a nil transactionType")
	}

	if transactionType.ID != id {
		t.Errorf("expected ID to be %d, but got %d", id, transactionType.ID)
	}
	if transactionType.Name != name {
		t.Errorf("expected Name to be '%s', but got '%s'", name, transactionType.Name)
	}
}
