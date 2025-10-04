import {defineStore} from 'pinia';
import {
  type Category,
  type Currency,
  type NewTransactionPayload,
  type Status,
  type SubCategory,
  type Transaction,
  transactionService,
  type TransactionType
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

function processCategoryChartData(transactions: Transaction[], typeName: string) {
  const categoryTotals = transactions
    .filter(tx => tx.Type?.name.toLowerCase() === typeName)
    .reduce((accumulator, tx) => {
      const category = tx.Category?.name || 'Uncategorized';
      const amount = parseFloat(tx.Amount);
      if (!isNaN(amount)) {
        accumulator[category] = (accumulator[category] || 0) + amount;
      }
      return accumulator;
    }, {} as Record<string, number>);

  const labels = Object.keys(categoryTotals);
  const data = Object.values(categoryTotals);

  return {labels, data};
}

function processSubCategoryChartData(transactions: Transaction[], typeName: string, selectedCategory: string | null) {
  if (!selectedCategory) {
    return {labels: [], data: []};
  }

  const subCategoryTotals = transactions
    .filter(tx => tx.Type?.name.toLowerCase() === typeName && tx.Category?.name === selectedCategory)
    .reduce((accumulator, tx) => {
      const subCategory = tx.SubCategory?.Name || 'Uncategorized';
      const amount = parseFloat(tx.Amount);
      if (!isNaN(amount)) {
        accumulator[subCategory] = (accumulator[subCategory] || 0) + amount;
      }
      return accumulator;
    }, {} as Record<string, number>);

  const labels = Object.keys(subCategoryTotals);
  const data = Object.values(subCategoryTotals);

  return {labels, data};
}

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
    filters: {...initialFiltersState},
    selectedCategoryForDashboard: null as string | null,
    dailySpendingFilter: {
      startDate: '',
      endDate: '',
    },
    dailySpendingChartType: 'debit' as 'debit' | 'credit' | 'refund',
  }),

  getters: {
    // Getter to apply filters to the raw transaction list
    filteredTransactions(state): Transaction[] {
      return state.transactions.filter(tx => {
        const txDate = new Date(tx.Date?.split('T')[0]);
        if (state.filters.startDate && txDate < new Date(state.filters.startDate)) return false;
        if (state.filters.endDate && txDate > new Date(state.filters.endDate)) return false;
        if (state.filters.startDate && !state.filters.endDate && txDate.getTime() !== new Date(state.filters.startDate).getTime()) return false;
        if (state.filters.categories.length > 0 && !state.filters.categories.includes(tx.Category?.name)) return false;
        if (state.filters.subCategories.length > 0 && !state.filters.subCategories.includes(tx.SubCategory?.Name)) return false;
        if (state.filters.statuses.length > 0 && !state.filters.statuses.includes(tx.Status?.name)) return false;
        if (state.filters.types.length > 0 && !state.filters.types.includes(tx.Type?.name)) return false;
        return true;
      });
    },

    // Getter to sort the already-filtered list
    sortedTransactions(): Transaction[] {
      const transactionsToSort = this.filteredTransactions;
      if (!this.sortKey) return transactionsToSort;
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

    debitChartData: (state) => processCategoryChartData(state.transactions, 'debit'),
    creditChartData: (state) => processCategoryChartData(state.transactions, 'credit'),
    refundChartData: (state) => processCategoryChartData(state.transactions, 'refund'),
    debitSubCategoryChartData: (state) => processSubCategoryChartData(state.transactions, 'debit', state.selectedCategoryForDashboard),
    creditSubCategoryChartData: (state) => processSubCategoryChartData(state.transactions, 'credit', state.selectedCategoryForDashboard),
    refundSubCategoryChartData: (state) => processSubCategoryChartData(state.transactions, 'refund', state.selectedCategoryForDashboard),
    availableCategoriesInFiltered(): string[] {
      const categorySet = new Set<string>();
      this.filteredTransactions.forEach(tx => {
        if (tx.Category?.name) {
          categorySet.add(tx.Category.name);
        }
      });
      return Array.from(categorySet).sort();
    },

    dailySpendingChartData(state): { labels: string[], data: number[] } {
      const daysOfWeek = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
      const dailyTotals = new Array(7).fill(0);

      const transactionsToProcess = state.transactions.filter(tx => {
        if (tx.Type?.name.toLowerCase() !== state.dailySpendingChartType) return false;
        const txDate = new Date(tx.Date?.split('T')[0]);
        if (state.dailySpendingFilter.startDate && txDate < new Date(state.dailySpendingFilter.startDate)) return false;
        if (state.dailySpendingFilter.endDate && txDate > new Date(state.dailySpendingFilter.endDate)) return false;
        return true;
      });

      transactionsToProcess.forEach(tx => {
        const dayIndex = new Date(tx.Date).getUTCDay();
        const amount = parseFloat(tx.Amount);
        if (!isNaN(amount)) {
          dailyTotals[dayIndex] += amount;
        }
      });

      return {
        labels: daysOfWeek,
        data: dailyTotals,
      };
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
    setSelectedCategoryForDashboard(categoryName: string | null) {
      this.selectedCategoryForDashboard = categoryName;
    },

    updateDailySpendingFilter(dates: { startDate?: string, endDate?: string }) {
      if (dates.startDate !== undefined) this.dailySpendingFilter.startDate = dates.startDate;
      if (dates.endDate !== undefined) this.dailySpendingFilter.endDate = dates.endDate;
    },
    clearDailySpendingFilter() {
      this.dailySpendingFilter.startDate = '';
      this.dailySpendingFilter.endDate = '';
    },
    setDailySpendingChartType(type: 'debit' | 'credit' | 'refund') {
      this.dailySpendingChartType = type;
    },
  },
});
