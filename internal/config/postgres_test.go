package config

import "testing"

func TestNew(t *testing.T) {
	// Table-driven tests to cover all validation scenarios.
	testCases := []struct {
		name             string
		inputURL         string
		inputName        string
		inputUser        string
		inputPassword    string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:          "Successful case with valid data",
			inputURL:      "localhost:5432",
			inputName:     "expenses_db",
			inputUser:     "admin",
			inputPassword: "secure_password",
			expectError:   false,
		},
		{
			name:             "Error case with empty URL",
			inputURL:         "",
			inputName:        "expenses_db",
			inputUser:        "admin",
			inputPassword:    "secure_password",
			expectError:      true,
			expectedErrorMsg: "postgres url required",
		},
		{
			name:             "Error case with empty database name",
			inputURL:         "localhost:5432",
			inputName:        "",
			inputUser:        "admin",
			inputPassword:    "secure_password",
			expectError:      true,
			expectedErrorMsg: "postgres database name required",
		},
		{
			name:             "Error case with empty username",
			inputURL:         "localhost:5432",
			inputName:        "expenses_db",
			inputUser:        "",
			inputPassword:    "secure_password",
			expectError:      true,
			expectedErrorMsg: "postgres username required",
		},
		{
			name:             "Error case with empty password",
			inputURL:         "localhost:5432",
			inputName:        "expenses_db",
			inputUser:        "admin",
			inputPassword:    "",
			expectError:      true,
			expectedErrorMsg: "postgres password required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewPostgresConfig(tc.inputURL, tc.inputName, tc.inputUser, tc.inputPassword)

			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error but got none")
				}
				if err.Error() != tc.expectedErrorMsg {
					t.Errorf("expected error message '%s' but got '%s'", tc.expectedErrorMsg, err.Error())
				}
				if config != nil {
					t.Error("expected config to be nil on error")
				}
			} else { // No error expected
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				if config == nil {
					t.Fatal("expected a config instance but got nil")
				}
				if config.DatabaseURL != tc.inputURL {
					t.Errorf("expected DatabaseURL '%s', but got '%s'", tc.inputURL, config.DatabaseURL)
				}
				if config.DatabaseName != tc.inputName {
					t.Errorf("expected DatabaseName '%s', but got '%s'", tc.inputName, config.DatabaseName)
				}
				if config.DatabaseUser != tc.inputUser {
					t.Errorf("expected DatabaseUser '%s', but got '%s'", tc.inputUser, config.DatabaseUser)
				}
				if config.DatabasePass != tc.inputPassword {
					t.Errorf("expected DatabasePass to be a non-empty string, but it was empty")
				}
			}
		})
	}
}
