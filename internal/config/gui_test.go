package config

import "testing"

func TestBuildGUIConfig(t *testing.T) {
	// --- Arrange ---
	// Define the input values for the test.
	expectedWidth := 1920
	expectedHeight := 1080
	expectedResizable := false

	// --- Act ---
	// Call the function we are testing.
	guiConfig := BuildGUIConfig(expectedWidth, expectedHeight, expectedResizable)

	// --- Assert ---
	// Check that the returned object is not nil.
	if guiConfig == nil {
		t.Fatal("BuildGUIConfig() returned a nil pointer")
	}

	// Check that each field was assigned the correct value.
	if guiConfig.Width != expectedWidth {
		t.Errorf("expected Width to be %d, but got %d", expectedWidth, guiConfig.Width)
	}

	if guiConfig.Height != expectedHeight {
		t.Errorf("expected Height to be %d, but got %d", expectedHeight, guiConfig.Height)
	}

	if guiConfig.Resizable != expectedResizable {
		t.Errorf("expected Resizable to be %v, but got %v", expectedResizable, guiConfig.Resizable)
	}
}
