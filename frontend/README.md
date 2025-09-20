# Transaction Tracker

A modern, single-page application for tracking personal financial transactions. This project is built with Vue 3 and
TypeScript, offering a fast, responsive, and type-safe user experience. It features a clean interface for adding,
viewing, sorting, and filtering transactions.

## ✨ Key Features

- **Add Transactions**: Easily add new transactions with details like description, amount, date, category, sub-category,
  type, and status.
- **Comprehensive List View**: View all transactions in a clear, sortable table.
- **Dynamic Filtering**: Filter transactions by a date range, categories, sub-categories, types, and statuses.
- **Interactive Sorting**: Sort the transaction list by date, description, category, or amount by clicking the table
  headers. The sorting cycles through ascending, descending, and unsorted states.
- **Amount Privacy**: Toggle the visibility of transaction amounts for enhanced privacy.
- **User Feedback**: Integrated loading and error states to keep the user informed during data fetching and submission.

## 🛠️ Tech Stack

This project leverages a modern frontend technology stack for a robust and maintainable application.

- **Framework**: Vue 3 (using Composition API with `<script setup>`)
- **Language**: TypeScript
- **State Management**: Pinia
- **Build Tool**: Vite
- **HTTP Client**: Axios
- **Routing**: Vue Router
- **Linting & Formatting**: ESLint + Prettier
- **Deployment**: Docker & Nginx

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- Node.js (`^20.19.0` or `>=22.12.0`)
- npm (comes with Node.js)

### Installation

1. Clone the repository to your local machine:
   ```sh
   git clone <repository-url>
   cd transaction-tracker-v2
   ```
2. Install the project dependencies:
   ```sh
   npm install
   ```

### Running the Application

1. **Development Mode**
   Start the Vite development server with hot-reloading:
   ```sh
   npm run dev
   ```
   The application will be available at `http://localhost:5173` (or the next available port).

2. **Production Build**
   To compile and minify the application for production:
   ```sh
   npm run build
   ```
   This command generates a `dist` folder with the static assets.

3. **Linting**
   To run ESLint and automatically fix issues:
   ```sh
   npm run lint
   ```

### Backend API Requirement

This is a frontend-only application. For full functionality, it requires a backend API server running and accessible at
`http://localhost:8080`. The API must provide the following endpoints as defined in `src/services/api.ts`:

- `GET /transactions`
- `POST /transactions`
- `GET /categories`
- `GET /categories/{category_name}/sub-categories`
- `GET /status`
- `GET /currencies`
- `GET /types`

## 🐳 Deployment with Docker

The project includes a multi-stage `Dockerfile` for creating an optimized, production-ready Nginx container.

1. **Build the Docker Image**
   From the root of the project, run:
   ```sh
   docker build -t transaction-tracker .
   ```

2. **Run the Docker Container**
   Run the newly created image in a container:
   ```sh
   docker run -p 8080:80 transaction-tracker
   ```
   The application will now be served by Nginx and accessible at `http://localhost:8080`.