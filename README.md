# Transaction Tracker

This is a full-stack web application for tracking personal financial transactions. It features a headless RESTful API
backend built with Go and a modern, reactive single-page application frontend built with Vue 3. The entire project is
containerized with Docker for easy setup and deployment.

---

<img width="1921" height="1050" alt="image" src="https://github.com/user-attachments/assets/8859f7f6-3b98-4717-bfe4-37124df4265e" />
<img width="1921" height="1075" alt="image" src="https://github.com/user-attachments/assets/85e41a97-d9f0-4511-9a07-30e2095c5fd5" />

---

<img width="1921" height="1043" alt="image" src="https://github.com/user-attachments/assets/ebdbd78f-5552-4347-a8f8-3c59b33de7d4" />
<img width="1921" height="1053" alt="image" src="https://github.com/user-attachments/assets/5832ac54-8f5e-4412-b083-3f54bd19a4e5" />

---

## ✨ Features

- **Secure User Authentication**: Full authentication flow with user registration and login. Sessions are managed using
  JSON Web Tokens (JWTs), and user passwords are securely hashed with `bcrypt`.
- **Private Transactions**: All transaction data is now user-specific. Users can only create, view, sort, and filter
  their own financial transactions.
- **Comprehensive Transaction Management**: Add, view, sort, and filter financial transactions.
- **Data Visualization Dashboard**: An interactive dashboard for visual data analysis, including:
    - A line chart for daily spending totals, filterable by date range and transaction type (Debit, Credit, or Refund).
    - Pie charts providing a breakdown of spending by category and sub-category.
- **RESTful API**: A clean, well-defined API for programmatic management of expenses and income, with protected
  endpoints for user data.
- **Dynamic Filtering & Sorting**: Filter transactions by date range, category, type, and status, and sort by various
  columns.
- **Data Persistence**: Utilizes a PostgreSQL database for reliable and structured data storage, now with multi-tenant
  support for transactions.
- **Structured Logging**: Backend application events are logged directly to a dedicated `app_logs` table in the database
  for robust monitoring and debugging.
- **Amount Privacy**: A UI toggle allows for hiding transaction amounts for enhanced privacy.

---

## 🛠️ Tech Stack

| Area                  | Technology                                                                     |
| --------------------- |--------------------------------------------------------------------------------|
| **Backend** | Go, PostgreSQL (with `pgx` driver), JWT, `bcrypt`, Docker                                     |
| **Frontend** | Vue 3, TypeScript, Pinia (for state management), Vue Router, Vite, Axios, Chart. js, Docker || * *Deployment** | Docker Compose, Nginx (for serving the frontend)                               |

---

## 🏗️ Architecture

The project is a monorepo containing two main components:

- **`backend`**: A Go application that serves a RESTful API. It follows a clean, layered architecture and now includes a
  robust authentication service using JWTs and secure password hashing. It also uses a blind index for securely querying
  encrypted usernames.
- **`frontend`**: A Vue 3 single-page application that consumes the backend API. It uses the Pinia state management
  library to handle global application state, including the user's authentication status.

---

## 🚀 Getting Started

The entire application stack can be run using Docker and Docker Compose.

### Prerequisites

* Docker
* Docker Compose
* A running PostgreSQL instance

### 1. Database Setup

This project requires an external PostgreSQL database.

1. Ensure your PostgreSQL server is running.
2. Create a new database for this application.
3. Execute the `backend/doc/postgres/initialize_db.sql` script against your new database to set up the required tables.

### 2. Environment Configuration

The backend service requires environment variables to connect to the database.

1. Create a new file named `.env` inside the `backend/` directory.
2. Add your database connection details to the file.

   ```bash
   # Example backend/.env file
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=your_db_host # Use 'host.docker.internal' if connecting from Docker to a DB on your host machine
   DB_PORT=your_db_port
   DB_NAME=your_db_name
   
   # Generate with `openssl rand -base64 32`
   ENCRYPTION_KEY="your-super-secret-encryption-key-32-chars"
   JWT_SECRET="your-super-secret-jwt-key-of-at-least-32-chars"
   SEARCH_HASH_KEY="another-different-super-secret-pepper-key"
   ```

### 3. Running the Application

With the database and environment configured, run the application using Docker Compose from the project root.

```bash
docker-compose up --build -d
```

Once the containers are up and running, the services will be available at:

- **Frontend Application**: `http://localhost:5559`
- **Backend API**: `http://localhost:8080`

---

## 🔌 API Documentation

The backend API is documented using the OpenAPI specification. New endpoints for `/register` and `/login` have been
added, and transaction-related endpoints are now protected. For detailed information, please refer to the official
documentation file:

- [OpenAPI Specification](backend/doc/openapi.yaml)
