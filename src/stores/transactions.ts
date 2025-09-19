import { defineStore } from 'pinia';
import {
  transactionService,
  type Transaction,
  type Category,
  type SubCategory
} from '@/services/api';

export const useTransactionStore = defineStore('transactions', {
  state: () => ({
    // State for transactions list
    transactions: [] as Transaction[],
    isLoading: false,
    error: null as string | null,

    // State for form options
    categories: [] as Category[],
    subCategories: [] as SubCategory[],
    statuses: [] as string[],
    currencies: [] as string[],
    transactionTypes: [] as string[],
    isLoadingOptions: false,
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

    async addTransaction(newTransaction: Transaction) {
      this.isLoading = true;
      this.error = null;
      try {
        await transactionService.createTransaction(newTransaction);
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
  },
});
