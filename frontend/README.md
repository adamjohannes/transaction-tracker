# UI Section

A modern, single-page application for tracking personal financial transactions. This project is built with Vue 3 and
TypeScript, offering a fast, responsive, and type-safe user experience.

## ✨ Key Features

- **Secure Authentication Flow**: Complete user authentication with dedicated pages for Login and Registration. User
  sessions are persisted in local storage.
- **Protected Views**: The core application views (Transactions and Dashboard) are protected. Unauthenticated users are
  automatically redirected to the login page.
- **Add Transactions**: Easily add new transactions with details like description, amount, date, and category.
- **Comprehensive List View**: View all your personal transactions in a clear, sortable table.
- **Dynamic Filtering**: Filter transactions by a date range, categories, sub-categories, types, and statuses.
- **Interactive Sorting**: Sort the transaction list by date, description, category, or amount.
- **Data Visualization Dashboard**: An interactive dashboard displays your personal transaction data visually.
- **Amount Privacy**: Toggle the visibility of transaction amounts for enhanced privacy.

## 🛠️ Tech Stack

This project leverages a modern frontend technology stack for a robust and maintainable application.

- **Framework**: Vue 3
- **Language**: TypeScript
- **State Management**: Pinia
- **Routing**: Vue Router
- **Build Tool**: Vite
- **HTTP Client**: Axios
- **Charting**: Chart.js with `vue-chartjs`
- **Routing**: Vue Router
- **Linting & Formatting**: ESLint + Prettier
- **Deployment**: Docker & Nginx

## 🏗️ Project Structure

The project follows a standard Vue project structure with a clear separation of concerns:

- `src/components/`: Reusable Vue components.
- `src/views/`: Top-level components for each application route, including new `LoginView.vue` and `RegisterView.vue`.
- `src/stores/`: Pinia store modules. A new `auth.ts` store manages authentication state, while `transactions.ts`
  handles all transaction-related data.
- `src/services/`: API communication layer. `api.ts` now includes an Axios interceptor to automatically attach the
  authentication token to requests.
- `src/router/`: Vue Router configuration, now with a `beforeEach` navigation guard to protect routes.
- `src/assets/`: Static assets like icons and styles.

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- Node.js (`^20.19.0` or `>=22.12.0`)
- npm (comes with Node.js)

```
cd transaction-tracker/frontend
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