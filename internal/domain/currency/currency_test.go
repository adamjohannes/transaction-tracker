package currency

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Define a struct for our test cases for a table-driven test
	testCases := []struct {
		name          string // The name of the test case
		inputCode     string // The input to the New function
		expectError   bool   // Whether we expect an error
		expectedError string // The expected error message, if any
	}{
		{
			name:        "Valid 3-Letter Code",
			inputCode:   "USD",
			expectError: false,
		},
		{
			name:          "Error on Short Code",
			inputCode:     "US",
			expectError:   true,
			expectedError: "currency must be 3 characters long",
		},
		{
			name:          "Error on Long Code",
			inputCode:     "USDE",
			expectError:   true,
			expectedError: "currency must be 3 characters long",
		},
		{
			name:          "Error on Empty Code",
			inputCode:     "",
			expectError:   true,
			expectedError: "currency must be 3 characters long",
		},
		{
			name:          "Error on Code with Numbers",
			inputCode:     "U5D",
			expectError:   true,
			expectedError: "currency code must only contain alphabetic characters",
		},
		{
			name:          "Error on Code with Symbols",
			inputCode:     "U$D",
			expectError:   true,
			expectedError: "currency code must only contain alphabetic characters",
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		// t.Run allows running sub-tests, making output cleaner
		t.Run(tc.name, func(t *testing.T) {
			// Call the function we are testing
			currency, err := New(tc.inputCode)

			// Case 1: We expect an error, but didn't get one
			if tc.expectError && err == nil {
				t.Fatalf("expected an error but got none")
			}

			// Case 2: We don't expect an error, but got one
			if !tc.expectError && err != nil {
				t.Fatalf("did not expect an error but got: %v", err)
			}

			// Case 3: We got an error, check if it's the correct one
			if err != nil {
				if err.Error() != tc.expectedError {
					t.Errorf("expected error '%s' but got '%s'", tc.expectedError, err.Error())
				}
				// If we got the expected error, the test for this case is done.
				return
			}

			// Case 4: We got a valid currency object, check its value
			if currency.code != tc.inputCode {
				t.Errorf("expected currency code '%s' but got '%s'", tc.inputCode, currency.code)
			}
		})
	}
}
