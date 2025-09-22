package transaction

import (
	"monthly-expenses-handler/internal/apierror"
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/domain/transaction_type"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          int64
	Amount      decimal.Decimal
	Category    *category.Category
	SubCategory *sub_category.SubCategory
	Date        time.Time
	Description string
	Status      *status.Status
	Currency    *currency.Currency
	Essential   bool
	Type        *transaction_type.TransactionType
}

func New(
	amount decimal.Decimal, transactionCategory category.Category,
	transactionSubcategory sub_category.SubCategory,
	date time.Time, description string, transactionStatus status.Status,
	transactionCurrency currency.Currency, essential bool,
	transactionType transaction_type.TransactionType) *Transaction {
	return &Transaction{
		ID:          -1,
		Amount:      amount,
		Category:    &transactionCategory,
		SubCategory: &transactionSubcategory,
		Date:        date,
		Description: description,
		Status:      &transactionStatus,
		Currency:    &transactionCurrency,
		Essential:   essential,
		Type:        &transactionType,
	}
}

func BuildTransaction(transaction map[string]any) (*Transaction, error) {
	if err := validateRequiredFields(transaction); err != nil {
		return nil, err
	}

	// --- Amount ---
	amountStr, _ := transaction["amount"].(string)
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return nil, apierror.NewValidationError("invalid amount format: %v", transaction["amount"])
	}

	// --- Date ---
	dateStr, _ := transaction["date"].(string)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, apierror.NewValidationError("invalid date format, use YYYY-MM-DD")
	}

	// --- Essential ---
	essential, ok := transaction["essential"].(bool)
	if !ok {
		return nil, apierror.NewValidationError("field 'essential' must be a boolean (true/false)")
	}

	// --- Description (Optional) ---
	description := ""
	if transaction["description"] != nil {
		description, _ = transaction["description"].(string)
	}

	// --- Build Domain Objects ---
	typeStr, _ := transaction["type"].(string)
	transactionType, err := transaction_type.New(typeStr)
	if err != nil {
		return nil, err
	}

	statusStr, _ := transaction["status"].(string)
	transactionStatus, err := status.New(statusStr)
	if err != nil {
		return nil, err
	}

	currencyStr, _ := transaction["currency"].(string)
	transactionCurrency, err := currency.New(currencyStr)
	if err != nil {
		return nil, err
	}

	categoryStr, _ := transaction["category"].(string)
	transactionCategory, err := category.New(categoryStr, "") // Description is optional
	if err != nil {
		return nil, err
	}

	subCategoryStr, _ := transaction["subCategory"].(string)
	// The parent ID is unknown here, so we use a placeholder. The repository will handle the lookup.
	transactionSubcategory, err := sub_category.New(0, subCategoryStr)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		ID:          -1, // ID is set by the database
		Amount:      amount,
		Category:    transactionCategory,
		SubCategory: transactionSubcategory,
		Date:        date,
		Description: description,
		Status:      transactionStatus,
		Currency:    transactionCurrency,
		Essential:   essential,
		Type:        transactionType,
	}, nil
}

// -- Helper functions --

// validateRequiredFields
// Validates if all required fields are present in the transaction.
func validateRequiredFields(transaction map[string]any) error {
	if transaction == nil {
		return apierror.NewValidationError("transaction payload is required")
	}

	requiredFields := []string{
		"amount", "date", "type", "essential",
		"status", "currency", "category",
	}

	for _, field := range requiredFields {
		val, ok := transaction[field]
		if !ok || val == nil {
			return apierror.NewValidationError("field '%s' is required", field)
		}
		// Check for empty strings as well
		if strVal, ok := val.(string); ok && strVal == "" {
			return apierror.NewValidationError("field '%s' cannot be empty", field)
		}
	}

	return nil
}

// FilterCriteria
// Holds all possible filters for a transaction query.
// Pointers are used to indicate optional filter fields.
type FilterCriteria struct {
	CategoryName    *string
	SubCategoryName *string
	Essential       *bool
	StartDate       *time.Time
	EndDate         *time.Time
	TypeName        *string
	CurrencyCode    *string
	StatusName      *string
}
