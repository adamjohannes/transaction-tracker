import axios from 'axios';

export interface Transaction {
  id?: number;
  amount: string;
  date: string;
  type: string;        // Will be populated by the API
  essential: boolean;
  status: string;      // Will be populated by the API
  currency: string;    // Will be populated by the API
  category: string;    // Will be populated by the API
  subCategory: string; // Will be populated by the API
  description: string;
}

export interface Category {
  id: number;
  name: string;
  description: string;
}

export interface SubCategory {
  id: number;
  parent_category: number;
  name: string;
}

const apiClient = axios.create({
  baseURL: 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const transactionService = {
  // GET /transactions
  getAllTransactions(): Promise<Transaction[]> {
    return apiClient.get('/transactions').then(res => res.data);
  },

  // POST /transactions
  createTransaction(transactionData: Transaction): Promise<Transaction> {
    return apiClient.post('/transactions', transactionData).then(res => res.data);
  },

  // GET /categories
  getCategories(): Promise<Category[]> {
    return apiClient.get('/categories').then(res => res.data);
  },

  // GET /categories/{category_name}/sub-categories
  getSubCategoriesByCategory(categoryName: string): Promise<SubCategory[]> {
    return apiClient.get(`/categories/${categoryName}/sub-categories`).then(res => res.data);
  },

  // GET /status
  getStatuses(): Promise<string[]> {
    return apiClient.get('/status').then(res => res.data);
  },

  // GET /currencies
  getCurrencies(): Promise<string[]> {
    return apiClient.get('/currencies').then(res => res.data);
  },

  // GET /types
  getTransactionTypes(): Promise<string[]> {
    return apiClient.get('/types').then(res => res.data);
  },
};
