package transaction

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/crypto"
	"monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/domain/transaction"
	"monthly-expenses-handler/internal/domain/transaction_type"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

// Repository
// Defines the interface for transaction data operations.
type Repository interface {
	Create(ctx context.Context, tx *transaction.Transaction, userID int64) (*transaction.Transaction, error)
	GetAllByUser(ctx context.Context, userID int64) ([]*transaction.Transaction, error)
	GetFiltered(ctx context.Context, filters *transaction.FilterCriteria, userID int64) ([]*transaction.Transaction, error)
	GetTransactionCountByTypeAndCategory(ctx context.Context) (map[string]map[string]int, error)
	GetTransactionCount(ctx context.Context, groupBy string) (map[string]map[string]int, error)
	GetSubCategoryAmounts(ctx context.Context) ([]SubCategoryAmount, error)
}

// SubCategoryAmount
// Holds the result of our new query
type SubCategoryAmount struct {
	TransactionType string
	CategoryName    string
	SubCategoryName string
	TotalAmount     decimal.Decimal
}

// GetSubCategoryAmounts
// Fetches the sum of amounts grouped by type, category, and sub-category.
func (r *postgresRepository) GetSubCategoryAmounts(ctx context.Context) ([]SubCategoryAmount, error) {
	query := `
		SELECT
			tt.name AS transaction_type,
			tc.name AS category_name,
			scat.name AS sub_category_name,
			SUM(t.amount) AS total_amount
		FROM transactions t
		JOIN transaction_types tt ON t.type = tt.id
		JOIN transaction_categories tc ON t.category = tc.id
		JOIN transaction_sub_categories scat ON t.sub_category = scat.id
		GROUP BY tt.name, tc.name, scat.name
		ORDER BY tt.name, tc.name, scat.name;
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query sub-category amounts: %w", err)
	}
	defer rows.Close()

	var results []SubCategoryAmount
	for rows.Next() {
		var item SubCategoryAmount
		if err := rows.Scan(&item.TransactionType, &item.CategoryName, &item.SubCategoryName, &item.TotalAmount); err != nil {
			return nil, fmt.Errorf("failed to scan sub-category amount row: %w", err)
		}
		results = append(results, item)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading sub-category amount rows: %w", rows.Err())
	}
	return results, nil
}

// postgresRepository
// Is the concrete implementation of the Repository interface for PostgreSQL.
type postgresRepository struct {
	db     *pgxpool.Pool
	crypto *crypto.CryptoService
}

// NewPostgresRepository
// Creates a new instance of the transaction repository.
// It takes the database connection pool as a dependency.
func NewPostgresRepository(db *pgxpool.Pool, cryptoSvc *crypto.CryptoService) Repository {
	return &postgresRepository{db: db, crypto: cryptoSvc}
}

// Create
// Inserts a new transaction record into the database.
// It uses subqueries to look up foreign key IDs from names.
func (r *postgresRepository) Create(ctx context.Context, tx *transaction.Transaction, userID int64) (*transaction.Transaction, error) {
	var id int64

	// Encrypt sensitive data before insertion
	encryptedAmount, err := r.crypto.Encrypt([]byte(tx.Amount.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt transaction amount: %w", err)
	}

	encryptedDesc, err := r.crypto.Encrypt([]byte(tx.Description))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt transaction description: %w", err)
	}

	query := `
		INSERT INTO transactions (user_id, amount, date, description, essential, type, status, currency, category, sub_category) 
		VALUES (
			$1, $2, $3, $4, $5,
			(SELECT id FROM transaction_types WHERE name = $6),
			(SELECT id FROM transaction_status WHERE name = $7),
			(SELECT code FROM currencies WHERE code = $8),
			(SELECT id FROM transaction_categories WHERE name = $9),
			(SELECT id FROM transaction_sub_categories WHERE name = $10 AND parent_category = (SELECT id FROM transaction_categories WHERE name = $9))
		) 
		RETURNING id`

	err = r.db.QueryRow(ctx, query,
		userID,
		encryptedAmount,
		tx.Date,
		encryptedDesc,
		tx.Essential,
		tx.Type.Name,
		tx.Status.Name,
		tx.Currency.Code,
		tx.Category.Name,
		tx.SubCategory.Name,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	tx.ID = id
	return tx, nil
}

// GetAllByUser
// Retrieves all transaction records from the database for a specific user.
func (r *postgresRepository) GetAllByUser(ctx context.Context, userID int64) ([]*transaction.Transaction, error) {
	query := `
		SELECT 
			t.id, t.amount, t.date, t.description, t.essential,
			tt.id AS type_id, tt.name AS type_name,
			ts.id AS status_id, ts.name AS status_name,
			cur.code AS currency_code,
			cat.id AS category_id, cat.name AS category_name, cat.description AS category_description,
			scat.id AS sub_category_id, scat.parent_category AS sub_category_parent_id, scat.name AS sub_category_name
		FROM transactions t
		JOIN transaction_types tt ON t.type = tt.id
		JOIN transaction_status ts ON t.status = ts.id
		JOIN currencies cur ON t.currency = cur.code
		JOIN transaction_categories cat ON t.category = cat.id
		JOIN transaction_sub_categories scat ON t.sub_category = scat.id
		WHERE t.user_id = $1
		ORDER BY t.date DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	return r.scanTransactions(rows)
}

func (r *postgresRepository) scanTransactions(rows pgx.Rows) ([]*transaction.Transaction, error) {
	var transactions []*transaction.Transaction
	for rows.Next() {
		var tx transaction.Transaction
		var txType transaction_type.TransactionType
		var txStatus status.Status
		var txCurrency currency.Currency
		var txCategory category.Category
		var txSubCategory sub_category.SubCategory

		var encryptedAmount, encryptedDesc []byte

		err := rows.Scan(
			&tx.ID, &encryptedAmount, &tx.Date, &encryptedDesc, &tx.Essential,
			&txType.ID, &txType.Name,
			&txStatus.ID, &txStatus.Name,
			&txCurrency.Code,
			&txCategory.Id, &txCategory.Name, &txCategory.Description,
			&txSubCategory.ID, &txSubCategory.ParentID, &txSubCategory.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		// Decrypt fields
		decryptedAmount, err := r.crypto.Decrypt(encryptedAmount)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt amount for tx %d: %w", tx.ID, err)
		}
		tx.Amount, err = decimal.NewFromString(string(decryptedAmount))
		if err != nil {
			return nil, fmt.Errorf("failed to parse decrypted amount for tx %d: %w", tx.ID, err)
		}

		decryptedDesc, err := r.crypto.Decrypt(encryptedDesc)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt description for tx %d: %w", tx.ID, err)
		}
		tx.Description = string(decryptedDesc)

		tx.Type = &txType
		tx.Status = &txStatus
		tx.Currency = &txCurrency
		tx.Category = &txCategory
		tx.SubCategory = &txSubCategory

		transactions = append(transactions, &tx)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading transaction rows: %w", rows.Err())
	}

	return transactions, nil
}

// GetAll
// Retrieves all transaction records from the database.
func (r *postgresRepository) GetAll(ctx context.Context) ([]*transaction.Transaction, error) {
	query := `
		SELECT 
			t.id, t.amount, t.date, t.description, t.essential,
			tt.id AS type_id, tt.name AS type_name,
			ts.id AS status_id, ts.name AS status_name,
			cur.code AS currency_code,
			cat.id AS category_id, cat.name AS category_name, cat.description AS category_description,
			scat.id AS sub_category_id, scat.parent_category AS sub_category_parent_id, scat.name AS sub_category_name
		FROM transactions t
		JOIN transaction_types tt ON t.type = tt.id
		JOIN transaction_status ts ON t.status = ts.id
		JOIN currencies cur ON t.currency = cur.code
		JOIN transaction_categories cat ON t.category = cat.id
		JOIN transaction_sub_categories scat ON t.sub_category = scat.id
		ORDER BY t.date DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	var transactions []*transaction.Transaction
	for rows.Next() {
		var tx transaction.Transaction
		var txType transaction_type.TransactionType
		var txStatus status.Status
		var txCurrency currency.Currency
		var txCategory category.Category
		var txSubCategory sub_category.SubCategory

		// Scan encrypted fields into byte slices
		var encryptedAmount, encryptedDesc []byte

		err := rows.Scan(
			&tx.ID, &encryptedAmount, &tx.Date, &encryptedDesc, &tx.Essential,
			&txType.ID, &txType.Name,
			&txStatus.ID, &txStatus.Name,
			&txCurrency.Code,
			&txCategory.Id, &txCategory.Name, &txCategory.Description,
			&txSubCategory.ID, &txSubCategory.ParentID, &txSubCategory.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		// Decrypt fields
		decryptedAmount, err := r.crypto.Decrypt(encryptedAmount)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt amount for tx %d: %w", tx.ID, err)
		}
		tx.Amount, err = decimal.NewFromString(string(decryptedAmount))
		if err != nil {
			return nil, fmt.Errorf("failed to parse decrypted amount for tx %d: %w", tx.ID, err)
		}

		decryptedDesc, err := r.crypto.Decrypt(encryptedDesc)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt description for tx %d: %w", tx.ID, err)
		}
		tx.Description = string(decryptedDesc)

		tx.Type = &txType
		tx.Status = &txStatus
		tx.Currency = &txCurrency
		tx.Category = &txCategory
		tx.SubCategory = &txSubCategory

		transactions = append(transactions, &tx)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading transaction rows: %w", rows.Err())
	}

	return transactions, nil
}

// GetFiltered
// Retrieves transactions based on a dynamic set of criteria.
func (r *postgresRepository) GetFiltered(ctx context.Context, filters *transaction.FilterCriteria, userID int64) ([]*transaction.Transaction, error) {
	baseQuery := `
		SELECT 
			t.id, t.amount, t.date, t.description, t.essential,
			tt.id AS type_id, tt.name AS type_name,
			ts.id AS status_id, ts.name AS status_name,
			cur.code AS currency_code,
			cat.id AS category_id, cat.name AS category_name, cat.description AS category_description,
			scat.id AS sub_category_id, scat.parent_category AS sub_category_parent_id, scat.name AS sub_category_name
		FROM transactions t
		JOIN transaction_types tt ON t.type = tt.id
		JOIN transaction_status ts ON t.status = ts.id
		JOIN currencies cur ON t.currency = cur.code
		JOIN transaction_categories cat ON t.category = cat.id
		JOIN transaction_sub_categories scat ON t.sub_category = scat.id
	`

	whereClauses := []string{"t.user_id = $1"}
	args := []any{userID}
	argCount := 2

	if filters.CategoryName != nil && *filters.CategoryName != "" {
		whereClauses = append(whereClauses, "cat.name = $"+strconv.Itoa(argCount))
		args = append(args, *filters.CategoryName)
		argCount++
	}
	if filters.SubCategoryName != nil && *filters.SubCategoryName != "" {
		whereClauses = append(whereClauses, "scat.name = $"+strconv.Itoa(argCount))
		args = append(args, *filters.SubCategoryName)
		argCount++
	}
	if filters.Essential != nil {
		whereClauses = append(whereClauses, "t.essential = $"+strconv.Itoa(argCount))
		args = append(args, *filters.Essential)
		argCount++
	}
	if filters.StartDate != nil {
		whereClauses = append(whereClauses, "t.date >= $"+strconv.Itoa(argCount))
		args = append(args, *filters.StartDate)
		argCount++
	}
	if filters.EndDate != nil {
		whereClauses = append(whereClauses, "t.date <= $"+strconv.Itoa(argCount))
		args = append(args, *filters.EndDate)
		argCount++
	}
	if filters.TypeName != nil && *filters.TypeName != "" {
		whereClauses = append(whereClauses, "tt.name = $"+strconv.Itoa(argCount))
		args = append(args, *filters.TypeName)
		argCount++
	}
	if filters.CurrencyCode != nil && *filters.CurrencyCode != "" {
		whereClauses = append(whereClauses, "cur.code = $"+strconv.Itoa(argCount))
		args = append(args, *filters.CurrencyCode)
		argCount++
	}
	if filters.StatusName != nil && *filters.StatusName != "" {
		whereClauses = append(whereClauses, "ts.name = $"+strconv.Itoa(argCount))
		args = append(args, *filters.StatusName)
		argCount++
	}

	finalQuery := baseQuery + " WHERE " + strings.Join(whereClauses, " AND ")
	finalQuery += " ORDER BY t.date DESC"

	rows, err := r.db.Query(ctx, finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query filtered transactions: %w", err)
	}
	defer rows.Close()

	return r.scanTransactions(rows)
}

// scanTransactions
// Is a helper function that scans rows from a pgx.Rows object into
// a slice of transaction domain objects. It helps reduce code duplication.
func scanTransactions(rows pgx.Rows) ([]*transaction.Transaction, error) {
	var transactions []*transaction.Transaction
	for rows.Next() {
		var tx transaction.Transaction
		var txType transaction_type.TransactionType
		var txStatus status.Status
		var txCurrency currency.Currency
		var txCategory category.Category
		var txSubCategory sub_category.SubCategory

		err := rows.Scan(
			&tx.ID, &tx.Amount, &tx.Date, &tx.Description, &tx.Essential,
			&txType.ID, &txType.Name,
			&txStatus.ID, &txStatus.Name,
			&txCurrency.Code,
			&txCategory.Id, &txCategory.Name, &txCategory.Description,
			&txSubCategory.ID, &txSubCategory.ParentID, &txSubCategory.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		tx.Type = &txType
		tx.Status = &txStatus
		tx.Currency = &txCurrency
		tx.Category = &txCategory
		tx.SubCategory = &txSubCategory

		transactions = append(transactions, &tx)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading transaction rows: %w", rows.Err())
	}

	return transactions, nil
}

// GetTransactionCountByTypeAndCategory
// Retrieves the number of transactions for each category, grouped by transaction type.
func (r *postgresRepository) GetTransactionCountByTypeAndCategory(ctx context.Context) (map[string]map[string]int, error) {
	query := `
		SELECT
			tt.name AS transaction_type,
			tc.name AS category_name,
			COUNT(t.id) AS transaction_count
		FROM transactions t
		JOIN transaction_types tt ON t.type = tt.id
		JOIN transaction_categories tc ON t.category = tc.id
		GROUP BY tt.name, tc.name
		ORDER BY tt.name, category_name;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transaction counts: %w", err)
	}
	defer rows.Close()

	results := make(map[string]map[string]int)

	for rows.Next() {
		var transactionType, categoryName string
		var transactionCount int

		if err := rows.Scan(&transactionType, &categoryName, &transactionCount); err != nil {
			return nil, fmt.Errorf("failed to scan transaction count row: %w", err)
		}

		if _, ok := results[transactionType]; !ok {
			results[transactionType] = make(map[string]int)
		}
		results[transactionType][categoryName] = transactionCount
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading transaction count rows: %w", rows.Err())
	}

	return results, nil
}

// GetTransactionCount
// Retrieves counts grouped by category or sub-category.
func (r *postgresRepository) GetTransactionCount(ctx context.Context, groupBy string) (map[string]map[string]int, error) {
	var query string
	// Dynamically set the query based on the groupBy parameter
	if groupBy == "sub_category" {
		query = `
			SELECT
				tt.name AS transaction_type,
				scat.name AS group_name,
				COUNT(t.id) AS transaction_count
			FROM transactions t
			JOIN transaction_types tt ON t.type = tt.id
			JOIN transaction_sub_categories scat ON t.sub_category = scat.id
			GROUP BY tt.name, group_name
			ORDER BY tt.name, group_name;
		`
	} else { // Default to grouping by category
		query = `
			SELECT
				tt.name AS transaction_type,
				tc.name AS group_name,
				COUNT(t.id) AS transaction_count
			FROM transactions t
			JOIN transaction_types tt ON t.type = tt.id
			JOIN transaction_categories tc ON t.category = tc.id
			GROUP BY tt.name, group_name
			ORDER BY tt.name, group_name;
		`
	}

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transaction counts: %w", err)
	}
	defer rows.Close()

	results := make(map[string]map[string]int)

	for rows.Next() {
		var transactionType, groupName string
		var transactionCount int

		// Scan into groupName, which is an alias for either category or sub-category name
		if err := rows.Scan(&transactionType, &groupName, &transactionCount); err != nil {
			return nil, fmt.Errorf("failed to scan transaction count row: %w", err)
		}

		if _, ok := results[transactionType]; !ok {
			results[transactionType] = make(map[string]int)
		}
		results[transactionType][groupName] = transactionCount
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading transaction count rows: %w", rows.Err())
	}

	return results, nil
}
