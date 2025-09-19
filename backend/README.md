# Transaction Tracker

Transaction Tracker is a cross-platform desktop application designed for managing personal financial transactions. It 
provides a clean graphical user interface to add, view, filter, and visualize your expenses and income. It can also run 
as a headless API server to manage transactions programmatically.

---

## ✨ Features

- **Transaction Management**: Add new transactions with details like amount, date, type, category, and sub-category.
- **Comprehensive Listing**: View all transactions in a sortable list. Sort data by any column, such as Amount, Date, or 
  Category, in both ascending and descending order.
- **Advanced Filtering**: A dedicated screen to filter transactions by date range, category, sub-category, transaction 
  type, and more.
- **Data Visualization**: Analyze your financial data through a dedicated charts view with pie charts that break down 
  your spending and income by category and sub-category.
- **Headless API**: Exposes RESTful endpoints to create and list transactions programmatically.
- **Cross-Platform**: Built with Fyne, the application can be compiled for various desktop platforms and Android.
- **CSV Data Import**: Includes a utility script to import historical transaction data from a CSV file directly into the 
  database.

---

## 🛠️ Technologies Used

- **Language**: Go
- **GUI Framework**: Fyne UI Toolkit
- **Database**: PostgreSQL (using the `pgx` driver)
- **Charting**: `fyne-charts` library
- **Changelog Automation**: `git-cliff`

---

## 🏗️ Architecture

The project follows a clean, layered architecture to separate concerns and improve maintainability:

- **Presentation Layer (`cmd/gui`, `cmd/api`)**: Handles user interactions, either through the Fyne toolkit for the GUI 
  or through HTTP handlers for the API.
- **Controller Layer (`internal/controller`)**: Acts as a bridge between the GUI and the repository, containing the core
  business logic.
- **Repository Layer (`internal/repository`)**: Manages all database operations, abstracting the SQL queries from the
  rest of the application.
- **Domain Layer (`internal/domain`)**: Defines the core data structures (entities) of the application, such as
  `Transaction` and `Category`.

---

## 🚀 Getting Started

### Prerequisites

- Go (version 1.24 or later)
- A running PostgreSQL instance
- Fyne dependencies. Please follow the [official Fyne guide](https://developer.fyne.io/started/) to install them.

### 1. Database Setup

The database connection details are configured using environment variables. The application will look for the following:

```bash
export DB_USER=[DB_USER]
export DB_PASSWORD=[DB_PASSWORD]
export DB_HOST=[DB_HOST]
export DB_PORT=[DB_PORT]
export DB_NAME=[DB_NAME]
```

Once your PostgreSQL server is running, create a database and run
the [initialization script](doc/postgres/initialize_db.sql) to set up the required
tables.

### 2. Run the Application

The application can run in two modes: GUI (default) or as a headless API server.

#### GUI Mode (Default)

To run the desktop application, navigate to the project's root directory and run:

```bash
go run cmd/main.go
```

#### API Server Mode

To run the headless API server, use the `-s` or `--server` flag:

```bash
# Using the long flag
go run cmd/main.go --server

# Using the shorthand flag
go run cmd/main.go -s
```

---

## 🔌 Headless API Endpoints

When running in server mode, the following endpoints are available.

### Create a Transaction

- Endpoint: `POST /transactions`
- Description: Creates a new transaction record in the database.
- Example Request:

```bash 
curl -X POST http://localhost:8080/transactions \
-H "Content-Type: application/json" \
-d '{
      "amount": "125.50",
      "date": "2025-09-19",
      "type": "Debit",
      "essential": true,
      "status": "Completed",
      "currency": "BRL",
      "category": "Supermercado",
      "subCategory": "Feira",
      "description": "Weekly groceries"
    }'
```

### List All Transactions

- Endpoint: `GET /transactions`
- Description: Retrieves a list of all transactions from the database.
- Example Request:

```bash 
curl http://localhost:8080/transactions
```

### List All Categories

- Endpoint: `GET /categories`
- Description: Retrieves a list of all transaction categories.
- Example Request:

```bash
curl http://localhost:8080/categories
```

### List All Sub-Categories

- Endpoint: `GET /sub-categories`
- Description: Retrieves a list of all transaction sub-categories.
- Example Request:

```bash
curl http://localhost:8080/sub-categories
```

### List Sub-Categories by Parent Category

- Endpoint: `GET /categories/{category_name}/sub-categories`
- Description: Retrieves sub-categories belonging to a specific parent category.
- Example Request:

```bash
curl http://localhost:8080/categories/Supermercado/sub-categories
```

### List All Statuses

- Endpoint: `GET /status`
- Description: Retrieves all possible transaction statuses.
- Example Request:

```bash
curl http://localhost:8080/status
```

### List All Currencies

- Endpoint: `GET /currencies`
- Description: Retrieves all supported currencies.
- Example Request:

```bash
curl http://localhost:8080/currencies
```

### List All Transaction Types

- Endpoint: `GET /types`
- Description: Retrieves all possible transaction types.
- Example Request:

```bash
curl http://localhost:8080/types
```

---

## 📦 Building and Packaging

### Desktop (Windows, macOS, Linux)

You can create a standalone executable using the `fyne` command:

### Android

An Android APK can be built using the provided script. This requires having the Android SDK and NDK set up as per the
Fyne documentation.

```bash
./android-compile.sh
```

---

## 📥 CSV Importer Script

The project includes a script to import transactions from a CSV file.

- Usage:

```bash
go run cmd/scripts/importer/main.go <path_to_your_csv_file>
```

An example CSV file is provided at `cmd/scripts/importer/example.csv`. The script expects the CSV to have a specific
column structure for a successful import.

---

## 📜 Changelog

All notable changes to this project are documented in the `CHANGELOG.md` file.
