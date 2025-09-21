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

type SortableKeys = 'Date' | 'Description' | 'Category' | 'Amount';

const initialFiltersState = {
  startDate: '',
  endDate: '',
  categories: [] as string[],
  subCategories: [] as string[],
  statuses: [] as string[],
  types: [] as string[],
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
    areAmountsVisible: false,
    sortKey: null as SortableKeys | null,
    sortOrder: 'asc' as 'asc' | 'desc',
    filters: { ...initialFiltersState },
  }),

  getters: {
    // Getter to apply filters to the raw transaction list
    filteredTransactions(state): Transaction[] {
      return state.transactions.filter(tx => {
        const txDate = new Date(tx.Date?.split('T')[0]);

        // Date Range Filter
        if (state.filters.startDate && txDate < new Date(state.filters.startDate)) {
          return false;
        }
        if (state.filters.endDate && txDate > new Date(state.filters.endDate)) {
          return false;
        }
        // Specific Date: if start date is set but no end date, it works as a specific date filter
        if (state.filters.startDate && !state.filters.endDate && txDate.getTime() !== new Date(state.filters.startDate).getTime()) {
          return false;
        }

        // Multi-select Filters
        if (state.filters.categories.length > 0 && !state.filters.categories.includes(tx.Category?.name)) {
          return false;
        }
        if (state.filters.subCategories.length > 0 && !state.filters.subCategories.includes(tx.SubCategory?.Name)) {
          return false;
        }
        if (state.filters.statuses.length > 0 && !state.filters.statuses.includes(tx.Status?.name)) {
          return false;
        }
        if (state.filters.types.length > 0 && !state.filters.types.includes(tx.Type?.name)) {
          return false;
        }

        return true; // Include transaction if all checks pass
      });
    },

    // Getter to sort the already-filtered list
    sortedTransactions(): Transaction[] {
      const transactionsToSort = this.filteredTransactions;

      if (!this.sortKey) {
        return transactionsToSort;
      }

      return [...transactionsToSort].sort((a, b) => {
        let valA, valB;
        switch (this.sortKey) {
          case 'Date':
            valA = new Date(a.Date || 0).getTime();
            valB = new Date(b.Date || 0).getTime();
            break;
          case 'Description':
            valA = a.Description.toLowerCase();
            valB = b.Description.toLowerCase();
            return this.sortOrder === 'asc' ? valA.localeCompare(valB) : valB.localeCompare(valA);
          case 'Category':
            valA = `${a.Category?.name}${a.SubCategory?.Name}`.toLowerCase();
            valB = `${b.Category?.name}${b.SubCategory?.Name}`.toLowerCase();
            return this.sortOrder === 'asc' ? valA.localeCompare(valB) : valB.localeCompare(valA);
          case 'Amount':
            valA = parseFloat(a.Amount);
            valB = parseFloat(b.Amount);
            break;
          default:
            return 0;
        }
        return this.sortOrder === 'asc' ? valA - valB : valB - valA;
      });
    },
  },

  actions: {
    // Filter actions
    updateFilters(newFilters: Partial<typeof initialFiltersState>) {
      this.filters = { ...this.filters, ...newFilters };
      if (newFilters.categories && newFilters.categories.length > 0) {
        this.fetchSubCategories(newFilters.categories[0]);
      } else if (newFilters.categories && newFilters.categories.length === 0) {
        this.subCategories = [];
      }
    },
    clearFilters() {
      this.filters = { ...initialFiltersState };
      this.subCategories = [];
    },

    // Data and UI actions
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

    setSort(key: SortableKeys) {
      if (this.sortKey === key) {
        // Same key, cycle through [asc -> desc -> inactive]
        if (this.sortOrder === 'asc') {
          this.sortOrder = 'desc';
        } else {
          this.sortKey = null; // Go to inactive state
        }
      } else {
        // New key, start with ascending
        this.sortKey = key;
        this.sortOrder = 'asc';
      }
    },
  },
});
