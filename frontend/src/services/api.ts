import axios from 'axios';

export interface Transaction {
  ID: number;
  Amount: string;
  Date: string;
  Type: TransactionType;
  Essential: boolean;
  Status: Status;
  Currency: Currency;
  Category: Category;
  SubCategory: SubCategory;
  Description: string;
}

export interface Category {
  id: number;
  name: string;
  description: string;
}

export interface SubCategory {
  ID: number;
  ParentID: number;
  Name: string;
}

export interface Status {
  id: number;
  name: string;
}

export interface Currency {
  Code: string;
}

export interface TransactionType {
  id: number;
  name: string;
}

// Interface for the transaction creation payload from the form
export interface NewTransactionPayload {
  amount: string;
  date: string;
  essential: boolean;
  description: string;
  type: string;
  status: string;
  currency: string;
  category: string;
  subCategory: string;
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
  createTransaction(transactionData: NewTransactionPayload): Promise<Transaction> {
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
  getStatuses(): Promise<Status[]> {
    return apiClient.get('/status').then(res => res.data);
  },

  // GET /currencies
  getCurrencies(): Promise<Currency[]> {
    return apiClient.get('/currencies').then(res => res.data);
  },

  // GET /types
  getTransactionTypes(): Promise<TransactionType[]> {
    return apiClient.get('/types').then(res => res.data);
  },
};
