package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"monthly-expenses-handler/internal/database"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: go run %s <path_to_csv_file>", os.Args[0])
	}

	csvFilePath := os.Args[1]
	log.Printf("🚀 Starting CSV import process for file: %s", csvFilePath)

	// Read all data from the CSV file
	header, records, err := readCSVFile(csvFilePath)
	if err != nil {
		log.Fatalf("❌ Failed to read CSV file: %v", err)
	}

	// Connect to the database
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Process each record
	var failedRows [][]string
	var successfulImports int

	query := `
		INSERT INTO transactions (date, category, sub_category, amount, essential, type, status, currency, description)
		VALUES (
			$1,
			(SELECT id FROM transaction_categories WHERE name = $2),
			(SELECT id FROM transaction_sub_categories WHERE name = $3),
			$4, $5,
			(SELECT id FROM transaction_types WHERE name = $6),
			(SELECT id FROM transaction_status WHERE name = 'Completed'),
			(SELECT code FROM currencies WHERE code = 'BRL'),
			''
		)`

	for _, record := range records {
		err := processRecord(context.Background(), pool, record, query)
		if err != nil {
			log.Printf("❌ Failed to process record %v: %v", record, err)
			failedRows = append(failedRows, record)
		} else {
			log.Printf("✅ Successfully inserted transaction for date %s", record[0])
			successfulImports++
		}
	}

	// Print the final summary of the operation
	printSummary(successfulImports, len(failedRows), failedRows, header)
}

// readCSVFile
// Opens and reads the entire CSV file, returning the header and data rows.
func readCSVFile(filePath string) (header []string, records [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	header, err = csvReader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("could not read header row: %w", err)
	}

	records, err = csvReader.ReadAll()

	if err != nil {
		return nil, nil, fmt.Errorf("could not read data rows: %w", err)
	}

	return header, records, nil
}

// processRecord
// Handles the logic for a single CSV record: parsing, validation, and DB insertion.
func processRecord(ctx context.Context, pool *pgxpool.Pool, record []string, query string) error {
	if len(record) < 8 || record[0] == "" {
		return nil // Skip empty or invalid rows silently
	}

	date, category, subCategory := record[0], record[1], record[2]
	amountStr, essentialStr, transactionType := record[3], record[4], record[7]
	amount, err := parseAmount(amountStr)

	if err != nil {
		return fmt.Errorf("could not parse amount '%s': %w", amountStr, err)
	}

	essential, err := strconv.ParseBool(strings.ToUpper(essentialStr))
	if err != nil {
		return fmt.Errorf("could not parse 'Essential' value '%s': %w", essentialStr, err)
	}

	_, err = pool.Exec(ctx, query,
		date, category, subCategory, amount, essential, transactionType,
	)
	if err != nil {
		return fmt.Errorf("db insert failed: %w", err)
	}

	return nil
}

// printSummary
// Logs the final results of the import process.
func printSummary(successful, failed int, failedRows [][]string, header []string) {
	log.Println("-------------------------------------------")
	log.Printf("📊 Import complete!")
	log.Printf("   - Successful inserts: %d", successful)
	log.Printf("   - Failed inserts: %d", failed)
	log.Println("-------------------------------------------")

	if len(failedRows) > 0 {
		log.Println("📋 Details of Rows that Failed DB Insertion:")
		log.Printf("Header: %v", header)
		for i, failedRecord := range failedRows {
			log.Printf("   - Row %d: %v", i+1, failedRecord)
		}
		log.Println("-------------------------------------------")
	}
}

// parseAmount
// Cleans and converts a currency string to a decimal.
func parseAmount(amountStr string) (decimal.Decimal, error) {
	cleanStr := strings.TrimSpace(amountStr)

	// Check if the string uses a comma for the decimal part (e.g., "1.806,46").
	if strings.Contains(cleanStr, ",") {
		// If it does, remove periods (which are thousand separators).
		cleanStr = strings.ReplaceAll(cleanStr, ".", "")
		// Then, replace the comma with a period for parsing.
		cleanStr = strings.ReplaceAll(cleanStr, ",", ".")
	}

	// If no comma is present (e.g., "76.98"), we assume the period is the
	// decimal separator and do nothing to it.

	return decimal.NewFromString(cleanStr)
}
