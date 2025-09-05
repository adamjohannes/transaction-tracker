package transaction

import (
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
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
}

func New(
	amount decimal.Decimal, transactionCategory category.Category,
	transactionSubcategory sub_category.SubCategory,
	date time.Time, description string, transactionStatus status.Status,
	transactionCurrency currency.Currency, essential bool) *Transaction {
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
	}
}

func Build(
	id int8, amount decimal.Decimal, transactionCategory category.Category,
	transactionSubcategory sub_category.SubCategory,
	date time.Time, description string, transactionStatus status.Status,
	transactionCurrency currency.Currency, essential bool) *Transaction {
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
	}
}
