# Transaction Tracker

Transaction Tracker is a cross-platform desktop application designed for managing personal financial transactions. It
provides a clean graphical user interface to add, view, filter, and visualize your expenses and income.

---

## ✨ Features

* **Transaction Management**: Add new transactions with details like amount, date, type, category, and sub-category.
* **Comprehensive Listing**: View all transactions in a sortable list. Sort data by any column, such as Amount, Date, or
  Category, in both ascending and descending order.
* **Advanced Filtering**: A dedicated screen to filter transactions by date range, category, sub-category, transaction
  type, and more.
* **Data Visualization**: Analyze your financial data through a dedicated charts view with pie charts that break down
  your spending and income by category and sub-category.
* **Cross-Platform**: Built with Fyne, the application can be compiled for various desktop platforms and Android.
* **CSV Data Import**: Includes a utility script to import historical transaction data from a CSV file directly into the
  database.

---

## 📸 Screenshots

*(Placeholder for Home Screen Screenshot)*
**Caption**: The main screen provides easy access to all features.

*(Placeholder for List View Screenshot)*
**Caption**: The list view showing sortable transactions.

*(Placeholder for Data/Charts View Screenshot)*
**Caption**: Pie charts visualizing spending by category.

---

## 🛠️ Technologies Used

* **Language**: Go
* **GUI Framework**: Fyne UI Toolkit
* **Database**: PostgreSQL (using the `pgx` driver)
* **Charting**: `fyne-charts` library
* **Changelog Automation**: `git-cliff`

---

## 🏗️ Architecture

The project follows a clean, layered architecture to separate concerns and improve maintainability:

* **Presentation Layer (`cmd/gui`)**: Handles the user interface and user interactions using the Fyne toolkit.
* **Controller Layer (`internal/controller`)**: Acts as a bridge between the GUI and the repository, containing the core
  business logic.
* **Repository Layer (`internal/repository`)**: Manages all database operations, abstracting the SQL queries from the
  rest of the application.
* **Domain Layer (`internal/domain`)**: Defines the core data structures (entities) of the application, such as
  `Transaction` and `Category`.

---

## 🚀 Getting Started

### Prerequisites

* Go (version 1.24 or later)
* A running PostgreSQL instance
* Fyne dependencies. Please follow the [official Fyne guide](https://developer.fyne.io/started/) to install them.

### 1. Database Setup

The database connection details are configured using environment variables. The application will look for the following:

* `DB_USER`
* `DB_PASSWORD`
* `DB_HOST`
* `DB_PORT`
* `DB_NAME`

Once your PostgreSQL server is running, create a database and run
the [initialization script](doc/postgres/initialize_db.sql) to set up the required
tables.

### 2. Run the Application

Navigate to the project's root directory and run the following command:

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
