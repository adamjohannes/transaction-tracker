# API Section

This section of the project is a headless RESTful API server designed for managing personal financial transactions. It
provides a robust backend to create, view, and manage your expenses and income programmatically.

---

## ✨ Features

- **Secure User Authentication & Authorization**: Full auth flow with `POST /register` and `POST /login` endpoints.
  Manages user sessions with secure, signed JSON Web Tokens (JWTs).
- **Password Hashing**: User passwords are never stored in plaintext. They are securely hashed using the `bcrypt`
  algorithm.
- **Transaction Management**: Add new transactions with details like amount, date, type, category, and sub-category. All
  transaction endpoints are protected and user-specific.
- **RESTful API**: Exposes clean and logical endpoints to manage transactions and related data.
- **Layered Architecture**: Built with a clean separation of concerns, making the codebase maintainable and easy to
  understand.
- **Data Persistence**: Uses a powerful PostgreSQL database for reliable and structured data storage.
- **CSV Data Import**: Includes a utility script to import historical transaction data from a CSV file directly into the
  database.

---

## 🛠️ Technologies Used

- **Language**: Go
- **Database**: PostgreSQL (using the `pgx` driver)
- **Authentication**: JWT (`github.com/golang-jwt/jwt/v5`), `bcrypt`

---

## 🏗️ Architecture

The project follows a clean, layered architecture to separate concerns and improve maintainability:

- **Presentation Layer (`cmd/api`)**: Handles user interactions through HTTP handlers for the API.
- **Controller Layer (`internal/controller`)**: Acts as a bridge between the API handlers and the repository, containing
  the core business logic. Includes new controllers for `auth` and `user`.
- **Repository Layer (`internal/repository`)**: Manages all database operations. The `transaction` repository is now
  user-aware, and a new `user` repository handles user data.
- **Domain Layer (`internal/domain`)**: Defines the core data structures, including a new `User` entity.
- **Auth Service (`internal/auth`)**: A dedicated service for handling password hashing, JWT generation, and validation.

### 🗄️ Database Schema

The database is designed with a core transactions table linked to several lookup tables for data consistency and
normalization. This structure makes it easy to manage and query financial data.

- `users`: Stores user ID, a securely **encrypted** username, and a hashed password. To allow for fast and secure
  lookups, it also contains a `username_search_hash` column, which serves as a **blind index**.
- `transactions`: This is the central table that stores individual financial records. It includes the amount, date, and
  foreign keys that link to the various lookup tables below.
- `transaction_categories` & `transaction_sub_categories`: These tables define the classification of transactions. Each
  sub-category (e.g., "Groceries") belongs to a parent category (e.g., "Supermarket").
- `transaction_types`: A simple lookup table that defines the nature of the transaction (e.g., 'Debit', 'Credit').
- `transaction_status`: A lookup table for the state of a transaction (e.g., 'Completed', 'Pending').
- `currencies`: Stores the supported currency codes (e.g., 'BRL', 'USD').
- `app_logs`: Contains structured application logs, including the log level, message, and attributes, written directly
  from the application's logger.

---

## 🚀 Getting Started

### Prerequisites & Configuration

- Go (version 1.24 or later)
- A running PostgreSQL instance

This project uses environment variables for configuration. A template is provided in the `.env.example` file. To get
started, copy it to a new `.env` file and update it with your database credentials. Most tools that work with Go will
load
this file automatically.

```bash
cp .env.example .env
```

**Note**: If these environment variables are not set, the application will fall back to default values suitable for a
local Postgres instance.

### 1. Database Initialization

Once your PostgreSQL server is running and configured, create a database and run
the [initialization script](doc/postgres/initialize_db.sql) to set up the required tables.

### 2. Run API Server

To run the headless API server, navigate to the project's root directory and run:

```bash
go run cmd/main.go
```

---

## 🔌 API Endpoints

For a detailed and complete API specification, you can view the official [OpenAPI](doc/openapi.yaml) documentation.

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

---

## 📥 CSV Importer Script

The project includes a script to import transactions from a CSV file.

### Usage:

```bash
go run cmd/scripts/importer/main.go <path_to_your_csv_file>
```

An example CSV file is provided at `cmd/scripts/importer/example.csv`. The script expects the CSV to have a specific
column structure for a successful import.

---
