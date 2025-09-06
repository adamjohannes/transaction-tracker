package transaction

import (
	"fmt"
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/domain/transaction_type"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          int8
	Amount      decimal.Decimal
	Category    *category.Category
	SubCategory *sub_category.SubCategory
	Date        time.Time
	Description string
	Status      *status.Status
	Currency    *currency.Currency
	Essential   bool
	Type        transaction_type.TransactionType
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
		Type:        transactionType,
	}
}

func Build(transaction map[string]any) (*Transaction, error) {
	if err := validateRequiredFields(transaction); err != nil {
		return nil, err
	}

	var id int8
	var amount decimal.Decimal
	var transactionCategory category.Category
	var transactionSubcategory sub_category.SubCategory
	var date time.Time
	var description string
	var transactionStatus status.Status
	var transactionCurrency currency.Currency
	var essential bool
	var transactionType transaction_type.TransactionType

	// Build ID, if present
	if transaction["id"] != nil {
		if tempId, err := strconv.Atoi(transaction["id"].(string)); err != nil {
			return nil, err
		} else {
			id = int8(tempId)
		}
	}

	// Build Amount
	if tempAmount, err := decimal.NewFromString(transaction["amount"].(string)); err != nil {
		return nil, err
	} else {
		amount = tempAmount
	}

	// Build Date
	if tempDate, err := time.Parse("2006-01-02", transaction["date"].(string)); err != nil {
		return nil, err
	} else {
		date = tempDate
	}

	// Build Type
	if transaction["type"] != nil {
		if tempType, err := transaction_type.New(transaction["type"].(string)); err != nil {
			return nil, err
		} else {
			transactionType = *tempType
		}
	}

	if transaction["description"] != nil {
		description = transaction["description"].(string)
	}

	if transaction["category"] != nil {
		if tempCategory, err := category.New(transaction["category"].(string), ""); err != nil {
			return nil, err
		} else {
			transactionCategory = *tempCategory
		}
	}

	return &Transaction{
		ID:          id,
		Amount:      amount,
		Category:    &transactionCategory,
		SubCategory: &transactionSubcategory,
		Date:        date,
		Description: description,
		Status:      &transactionStatus,
		Currency:    &transactionCurrency,
		Essential:   essential,
		Type:        transactionType,
	}, nil
}

// -- Helper functions --

// validateRequiredFields
// Validates if all required fields are present in the transaction.
func validateRequiredFields(transaction map[string]any) error {
	if transaction == nil {
		return fmt.Errorf("transaction is required")
	}

	if transaction["amount"] == nil {
		return fmt.Errorf("amount is required")
	}

	if transaction["date"] == nil {
		return fmt.Errorf("date is required")
	}

	if transaction["type"] == nil {
		return fmt.Errorf("type is required")
	}

	if transaction["essential"] == nil {
		return fmt.Errorf("essential is required")
	}

	if transaction["status"] == nil {
		return fmt.Errorf("status is required")
	}

	if transaction["currency"] == nil {
		return fmt.Errorf("currency is required")
	}

	if transaction["category"] == nil {
		return fmt.Errorf("category is required")
	}

	return nil
}
