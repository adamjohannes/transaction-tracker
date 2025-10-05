import { defineStore } from 'pinia';
import { authService, type AuthPayload } from '@/services/api';
import router from '@/router';
import { useTransactionStore } from './transactions';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    error: null as string | null,
    isLoading: false,
  }),

  getters: {
    isLoggedIn(): boolean {
      return !!this.token;
    },
  },

  actions: {
    setToken(token: string | null) {
      this.token = token;
      if (token) {
        localStorage.setItem('token', token);
      } else {
        localStorage.removeItem('token');
      }
    },

    async login(payload: AuthPayload) {
      this.isLoading = true;
      this.error = null;
      try {
        const { token } = await authService.login(payload);
        this.setToken(token);

        // After login, fetch user-specific data
        const transactionStore = useTransactionStore();
        await transactionStore.fetchTransactions();
        await transactionStore.fetchFormOptions();
        await router.push('/');
      } catch (err: any) {
        this.error = err.response?.data?.error || 'Failed to log in.';
        console.error(err);
      } finally {
        this.isLoading = false;
      }
    },

    async register(payload: AuthPayload) {
      this.isLoading = true;
      this.error = null;
      try {
        const { token } = await authService.register(payload);
        this.setToken(token);
        const transactionStore = useTransactionStore();
        await transactionStore.fetchTransactions();
        await transactionStore.fetchFormOptions();
        await router.push('/');
      } catch (err: any) {
        this.error = err.response?.data?.error || 'Failed to register.';
        console.error(err);
      } finally {
        this.isLoading = false;
      }
    },

    logout() {
      const transactionStore = useTransactionStore();
      this.setToken(null);
      transactionStore.$reset(); // Reset transaction store state
      router.push('/login');
    },
  },
});
