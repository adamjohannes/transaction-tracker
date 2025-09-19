# API Section

This section of the project is a headless RESTful API server designed for managing personal financial transactions. It
provides a robust backend to create, view, and manage your expenses and income programmatically, making it an ideal
foundation for various client applications (e.g., web, mobile, or desktop).

---

## ✨ Features

- **Transaction Management**: Add new transactions with details like amount, date, type, category, and sub-category.
- **RESTful API**: Exposes clean and logical endpoints to create and list transactions, categories, and other lookup
  data programmatically.
- **Layered Architecture**: Built with a clean separation of concerns, making the codebase maintainable and easy to
  understand.
- **Data Persistence**: Uses a powerful PostgreSQL database for reliable and structured data storage.
- **CSV Data Import**: Includes a utility script to import historical transaction data from a CSV file directly into the
  database.

---

## 🛠️ Technologies Used

- **Language**: Go
- **Database**: PostgreSQL (using the `pgx` driver)

---

## 🏗️ Architecture

The project follows a clean, layered architecture to separate concerns and improve maintainability:

- **Presentation Layer (`cmd/api`)**: Handles user interactions through HTTP handlers for the API.
- **Controller Layer (`internal/controller`)**: Acts as a bridge between the API handlers and the repository, containing
  the core business logic.
- **Repository Layer (`internal/repository`)**: Manages all database operations, abstracting the SQL queries from the
  rest of the application.
- **Domain Layer (`internal/domain`)**: Defines the core data structures (entities) of the application, such as
  `Transaction` and `Category`.

---

## 🚀 Getting Started

### Prerequisites

- Go (version 1.24 or later)
- A running PostgreSQL instance

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

### 2. Run API Server

To run the headless API server, navigate to the project's root directory and run:

```bash
go run cmd/main.go
```

---

## 🔌 API Endpoints

For a detailed and complete API specification, you can view the official [OpenAPI](doc/openapi.yaml) documentation.

When running, the following endpoints are available.

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

## 📥 CSV Importer Script

The project includes a script to import transactions from a CSV file.

- Usage:

```bash
go run cmd/scripts/importer/main.go <path_to_your_csv_file>
```

An example CSV file is provided at `cmd/scripts/importer/example.csv`. The script expects the CSV to have a specific
column structure for a successful import.

---
