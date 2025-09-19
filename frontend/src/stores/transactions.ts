import { defineStore } from 'pinia';
import {
  transactionService,
  type Transaction,
  type Category,
  type SubCategory,
  type Status,
  type Currency,
  type TransactionType,
  type NewTransactionPayload
} from '@/services/api';

// Helper type for the form's internal data structure (PascalCase)
type TransactionFormData = {
  Amount: string | number;
  Date: string;
  Essential: boolean;
  Description: string;
  Type: string;
  Status: string;
  Currency: string;
  Category: string;
  SubCategory: string;
};

export const useTransactionStore = defineStore('transactions', {
  state: () => ({
    transactions: [] as Transaction[],
    isLoading: false,
    error: null as string | null,
    categories: [] as Category[],
    subCategories: [] as SubCategory[],
    statuses: [] as Status[],
    currencies: [] as Currency[],
    transactionTypes: [] as TransactionType[],
    isLoadingOptions: false,
    areAmountsVisible: true,
  }),

  actions: {
    async fetchTransactions() {
      this.isLoading = true;
      this.error = null;
      try {
        this.transactions = await transactionService.getAllTransactions();
      } catch (err) {
        this.error = 'Failed to fetch transactions.';
        console.error(err);
      } finally {
        this.isLoading = false;
      }
    },

    async addTransaction(formData: TransactionFormData) {
      this.isLoading = true;
      this.error = null;
      try {
        // Transform the form data to the required API payload
        const payload: NewTransactionPayload = {
          amount: String(formData.Amount),
          date: formData.Date,
          description: formData.Description,
          category: formData.Category,
          subCategory: formData.SubCategory,
          type: formData.Type,
          status: formData.Status,
          currency: formData.Currency,
          essential: formData.Essential,
        };
        await transactionService.createTransaction(payload);
        await this.fetchTransactions(); // Refresh list after adding
      } catch (err) {
        this.error = 'Failed to add transaction.';
        console.error(err);
      } finally {
        this.isLoading = false;
      }
    },

    async fetchFormOptions() {
      this.isLoadingOptions = true;
      try {
        // Fetch all static options in parallel
        const [categories, statuses, currencies, types] = await Promise.all([
          transactionService.getCategories(),
          transactionService.getStatuses(),
          transactionService.getCurrencies(),
          transactionService.getTransactionTypes(),
        ]);
        this.categories = categories;
        this.statuses = statuses;
        this.currencies = currencies;
        this.transactionTypes = types;
      } catch (err) {
        this.error = 'Failed to load form options.';
        console.error(err);
      } finally {
        this.isLoadingOptions = false;
      }
    },

    async fetchSubCategories(categoryName: string) {
      if (!categoryName) {
        this.subCategories = [];
        return;
      }
      this.isLoadingOptions = true;
      try {
        this.subCategories = await transactionService.getSubCategoriesByCategory(categoryName);
      } catch (err) {
        this.error = `Failed to load sub-categories for ${categoryName}.`;
        console.error(err);
      } finally {
        this.isLoadingOptions = false;
      }
    },

    toggleAmountVisibility() {
      this.areAmountsVisible = !this.areAmountsVisible;
    },
  },
});
