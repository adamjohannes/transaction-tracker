package transaction

import (
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/domain/transaction_type"
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestNew(t *testing.T) {
	amount := decimal.NewFromFloat(150.75)
	cat := category.Build(1, "Food", "Restaurant expenses")
	subCat := sub_category.Build(10, 1, "Dinner")
	date := time.Date(2025, time.September, 4, 21, 14, 0, 0, time.UTC)
	description := "Dinner with friends"
	st, _ := status.New("Completed")
	curr, _ := currency.New("BRL")
	essential := false
	tt, _ := transaction_type.New("Expense")

	transaction := New(amount, *cat, *subCat, date, description, *st, *curr, essential, *tt)

	if transaction == nil {
		t.Fatal("New() returned a nil transaction")
	}

	if transaction.ID != -1 {
		t.Errorf("expected ID to be -1, but got %d", transaction.ID)
	}
	if !transaction.Amount.Equal(amount) {
		t.Errorf("expected amount %s, but got %s", amount, transaction.Amount)
	}
	if !reflect.DeepEqual(transaction.Category, cat) {
		t.Errorf("expected category %+v, but got %+v", cat, transaction.Category)
	}
	if !reflect.DeepEqual(transaction.Type, tt) {
		t.Errorf("expected type %+v, but got %+v", tt, transaction.Type)
	}
}

func TestBuild(t *testing.T) {
	transactionData := map[string]any{
		"id":          1,
		"amount":      "99.99",
		"date":        "2025-09-04",
		"description": "New shirt",
		"essential":   true,
		"type":        "Expense",
		"status":      "Pending",
		"currency":    "USD",
		"category":    "Shopping",
		"subCategory": "T-shirt",
	}

	transaction, err := BuildTransaction(transactionData)

	if err != nil {
		t.Fatalf("BuildTransaction() returned an unexpected error: %v", err)
	}
	if transaction == nil {
		t.Fatal("BuildTransaction() returned a nil transaction")
	}

	if transaction.ID != -1 {
		t.Errorf("expected ID to be %d, but got %d", -1, transaction.ID)
	}

	expectedAmount := decimal.NewFromFloat(99.99)
	if !transaction.Amount.Equal(expectedAmount) {
		t.Errorf("expected amount %s, but got %s", expectedAmount, transaction.Amount)
	}

	expectedDate := time.Date(2025, time.September, 4, 0, 0, 0, 0, time.UTC)
	if !transaction.Date.Equal(expectedDate) {
		t.Errorf("expected date %v, but got %v", expectedDate, transaction.Date)
	}

	if transaction.Category.Name != "Shopping" {
		t.Errorf("expected category 'Shopping', but got '%s'", transaction.Category.Name)
	}

	if transaction.SubCategory.Name != "T-shirt" {
		t.Errorf("expected subCategory 'T-shirt', but got '%s'", transaction.SubCategory.Name)
	}

	if transaction.Status.Name != "Pending" {
		t.Errorf("expected status 'Pending', but got '%s'", transaction.Status.Name)
	}
}
