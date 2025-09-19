package status

import "testing"

func TestNew(t *testing.T) {
	// Table-driven tests for various scenarios.
	testCases := []struct {
		name             string
		inputName        string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:        "Successful case with valid name",
			inputName:   "Active",
			expectError: false,
		},
		{
			name:             "Error case with empty name",
			inputName:        "", // Name is required
			expectError:      true,
			expectedErrorMsg: "name is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := New(tc.inputName)

			if tc.expectError {
				// We expected an error.
				if err == nil {
					t.Fatal("expected an error but got none")
				}
				if err.Error() != tc.expectedErrorMsg {
					t.Errorf("expected error message '%s' but got '%s'", tc.expectedErrorMsg, err.Error())
				}
				if status != nil {
					t.Error("expected status to be nil on error")
				}
			} else {
				// We did NOT expect an error.
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				if status == nil {
					t.Fatal("expected a status instance but got nil")
				}
				if status.ID != -1 {
					t.Errorf("expected ID to be initialized to -1, but got %d", status.ID)
				}
				if status.Name != tc.inputName {
					t.Errorf("expected name to be '%s', but got '%s'", tc.inputName, status.Name)
				}
			}
		})
	}
}
