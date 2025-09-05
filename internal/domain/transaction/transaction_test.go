package transaction

import (
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// TestNew verifies that a new transaction is created with the correct default values.
func TestNew(t *testing.T) {
	// --- Arrange ---
	amount := decimal.NewFromFloat(150.75)
	cat := category.Build(1, "Food", "Restaurant expenses")
	subCat := sub_category.Build(10, 1, "Dinner")
	date := time.Date(2025, time.September, 4, 21, 14, 0, 0, time.UTC)
	description := "Dinner with friends"
	st, _ := status.New("Completed")
	curr, _ := currency.New("BRL")
	essential := false

	// --- Act ---
	transaction := New(amount, *cat, *subCat, date, description, *st, *curr, essential)

	// --- Assert ---
	if transaction == nil {
		t.Fatal("New() returned a nil transaction")
	}

	// Check for the default ID
	if transaction.ID != -1 {
		t.Errorf("expected ID to be -1, but got %d", transaction.ID)
	}

	// Check that all other fields were assigned correctly
	if !transaction.Amount.Equal(amount) {
		t.Errorf("expected amount %s, but got %s", amount, transaction.Amount)
	}
	if !reflect.DeepEqual(transaction.Category, cat) {
		t.Errorf("expected category %+v, but got %+v", cat, transaction.Category)
	}
	if !reflect.DeepEqual(transaction.SubCategory, subCat) {
		t.Errorf("expected subCategory %+v, but got %+v", subCat, transaction.SubCategory)
	}
	if !transaction.Date.Equal(date) {
		t.Errorf("expected date %v, but got %v", date, transaction.Date)
	}
	if transaction.Description != description {
		t.Errorf("expected description '%s', but got '%s'", description, transaction.Description)
	}
	if !reflect.DeepEqual(transaction.Status, st) {
		t.Errorf("expected status %+v, but got %+v", st, transaction.Status)
	}
	if !reflect.DeepEqual(transaction.Currency, curr) {
		t.Errorf("expected currency %+v, but got %+v", curr, transaction.Currency)
	}
}

// TestBuild verifies that a transaction is reconstructed correctly from existing data.
func TestBuild(t *testing.T) {
	// --- Arrange ---
	id := int8(42)
	amount := decimal.NewFromFloat(99.99)
	cat := category.Build(2, "Shopping", "Clothing")
	subCat := sub_category.Build(20, 2, "T-shirt")
	date := time.Date(2025, time.September, 4, 10, 0, 0, 0, time.UTC)
	description := "New shirt"
	st, _ := status.New("Pending")
	curr, _ := currency.New("USD")
	essential := true

	// --- Act ---
	transaction := Build(id, amount, *cat, *subCat, date, description, *st, *curr, essential)

	// --- Assert ---
	if transaction == nil {
		t.Fatal("Build() returned a nil transaction")
	}

	// Check that all fields, including the ID, were assigned correctly
	if transaction.ID != id {
		t.Errorf("expected ID to be %d, but got %d", id, transaction.ID)
	}
	if !transaction.Amount.Equal(amount) {
		t.Errorf("expected amount %s, but got %s", amount, transaction.Amount)
	}
	if !reflect.DeepEqual(transaction.Category, cat) {
		t.Errorf("expected category %+v, but got %+v", cat, transaction.Category)
	}
	if !reflect.DeepEqual(transaction.SubCategory, subCat) {
		t.Errorf("expected subCategory %+v, but got %+v", subCat, transaction.SubCategory)
	}
	if !transaction.Date.Equal(date) {
		t.Errorf("expected date %v, but got %v", date, transaction.Date)
	}
	if transaction.Description != description {
		t.Errorf("expected description '%s', but got '%s'", description, transaction.Description)
	}
	if !reflect.DeepEqual(transaction.Status, st) {
		t.Errorf("expected status %+v, but got %+v", st, transaction.Status)
	}
	if !reflect.DeepEqual(transaction.Currency, curr) {
		t.Errorf("expected currency %+v, but got %+v", curr, transaction.Currency)
	}
}
