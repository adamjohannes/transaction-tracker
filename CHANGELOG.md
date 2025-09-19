# Changelog

All notable changes to this project will be documented in this file.

## [0.1.0] - 2025-09-19

### 🚀 Features

- *(gui)* Add sorting methods to the header buttons
- Implement a charts view
- *(gui)* Implement a separation by tabs for viewing Debit, Credit and Refund transactions
- Add a combo box to 'data' view to alternate between listing the categories and sub categories
- Add a combo box to cycle between categories and sub categories in the 'data' view
- *(api)* Add headless API mode

### 🐛 Bug Fixes

- *(repository/transaction)* Adjust sub query for getting sub category name when inserting new transactions to the DB
- Adjust 'data' view min width

### 🚜 Refactor

- *(gui)* Make the sorting buttons align with the table's content
- *(gui)* Insert the charts in a scrollable container
- Create a button inside the 'list' view to open the 'charts' view
- Adjust icon path
- *(api)* Initialize API with CLI options

### 📚 Documentation

- Add README.md file
- Add API docs

### ⚙️ Miscellaneous Tasks

- Add README.md to .gitignore

## [0.0.5] - 2025-09-07

### 🚀 Features

- *(gui)* Add a refresh button to 'list' view
- *(gui)* Add a filter button
- *(scripts)* Add a script to import transactions from a CSV file
- Add script to compile to android
- Implement transaction filtering in the 'list' view

### 🐛 Bug Fixes

- *(scripts/importer)* Add check for period
- Set the value for transaction ID as a big int

### 🚜 Refactor

- *(transaction)* Remove sub category from transaction fields
- *(gui)* Adjust width for sub category field in list view
- *(gui)* Adjust width for description field in list view

### 📚 Documentation

- *(scripts/importer)* Add example CSV file for importer

### ⚙️ Miscellaneous Tasks

- Add android-compile.sh to .gitignore

## [0.0.4] - 2025-09-07

### 🚀 Features

- *(gui)* Start to implement the list view
- Dynamically fetch transactions data from the database in the 'list' view

### 🐛 Bug Fixes

- *(gui)* Refresh the table when opening the 'list' view in order to see the new transactions

### 🚜 Refactor

- *(gui)* Make column headers follow focus

## [0.0.3] - 2025-09-07

### 🚀 Features

- *(internal)* Add controller section
- *(controller/transaction)* Add a controller for Transaction
- *(repository/transaction)* Add a repository layer for Transaction
- *(repository/transaction)* Implement method to add transactions to the DB
- *(repository/transaction)* Implement method to add transactions to the DB
- *(gui)* Add transaction controller to GUI
- Connect transaction creation with the database
- Dynamically fetch categories from the database
- Dynamically fetch sub categories from the database

### 🐛 Bug Fixes

- *(database)* Adjust logic for interacting with the database

### 🚜 Refactor

- *(gui)* Rename 'home.go' to 'homeGUI.go'
- *(gui)* Refactor 'add' screen
- Update transaction handling

### ⚙️ Miscellaneous Tasks

- Add repostory/ to .gitignore
- Add category/ to .gitignore

## [0.0.2] - 2025-09-06

### 🚀 Features

- *(doc/postgres)* Add script to initialize the database
- *(domain/transaction)* Add Essential field to transaction
- *(gui)* Add Essential field
- *(gui)* Add a temp map to collect the values for the transaction

### 🚜 Refactor

- *(gui)* Refactor category section to be a radio group
- *(gui)* Set category as required and refactor fields order in UI
- *(gui)* Replace VBox container by Form
- *(gui)* Refactor status to radio group
- *(gui)* Add a generic method to build the radio groups inside an accordion
- *(gui)* Set category as required and refactor fields order in UI
- *(gui)* Replace calls to fmt.Println to log.Println
- *(gui)* Update cancel and save buttons
- *(doc/postgres)* Update DB SQL init query
- *(gui)* Add a controller and finish implementation of navigation between the 'home' and 'add' screen

### ⚙️ Miscellaneous Tasks

- *(gui)* Fix typo in 'Amount' label
- *(gui)* Refactor text data
- Add doc/ to .gitignore
- Add inner .gitignore to doc/

## [0.0.1] - 2025-09-05

### 🚀 Features

- Add .gitignore
- Add go.mod and go.sum
- Add main.go
- Add cmd/ and internal/ to .gitignore
- Add cmd/ and internal/ to .gitignore
- *(internal)* Add inner .gitignore
- *(internal/domain)* Add inner .gitignore
- *(internal/domain)* Add basic domain structure
- *(cmd)* Add inner .gitignore
- *(internal/domain/currency)* Implemented Currency domain
- *(internal/domain/currency)* Implemented Category domain
- *(internal/domain/currency)* Implemented Status domain
- *(internal/domain/sub_category)* Implemented SubCategory domain
- *(internal/domain/sub_category)* Implemented TransactionType domain
- *(internal/domain/transaction)* Implemented Transaction domain
- *(internal/config)* Add postgres config struct
- *(internal/database)* Add postgres db handler
- *(cmd/gui)* Started to implement the GUI
- *(.github)* Add inner .gitignore
- *(.github/workflows)* Add cd pipeline
- Add git cliff to project
- *(.github/workflows)* Add CI pipeline
- *(internal/config)* Add config struct for GUI
- *(cmd/gui)* Implement initial screen to add new transactions
- *(cmd/gui)* Implement initial home screen

### 🐛 Bug Fixes

- *(workflows)* Install fyne before trying to build
- *(workflows)* Add write permission to CI pipeline
- *(workflows)* Refactor logic for collecting release notes

### 🚜 Refactor

- *(cmd)* Move main.go to cmd/api/main.go
- *(internal/domain/currency)* Expose Code attribute
- *(internal/domain/category)* Adjust return type of Build()
- *(internal/domain/status)* Adjust return type of Build()
- *(internal/domain/category)* Adjust return type of Build()
- *(cmd)* Move main.go to cmd/
- *(cmd)* Start testing new GUI
- *(cmd/gui)* Rename gui.go to home.go
- *(internal/config)* Rename New to NewPostgresConfig
- *(internal/config)* Change width and height types to float32

### 📚 Documentation

- Add CHANGELOG.md

### ⚙️ Miscellaneous Tasks

- Fix typo
- Fix typo
- Go mod tidy
- Add config/ to .gitignore
- Added database/ to .gitignore
- Add main.go to .gitignore
- Add fyne to project
- Add gui/ to .gitignore
- Go mod tidy
- Updated .gitignore
- Updated .gitignore
- Updated .gitignore
- Import date picker widget

<!-- generated by git-cliff -->
