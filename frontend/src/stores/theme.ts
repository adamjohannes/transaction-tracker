import { defineStore } from 'pinia';

type Theme = 'light' | 'dark';

export const useThemeStore = defineStore('theme', {
  // Default theme to dark
  state: () => ({
    theme: (localStorage.getItem('theme') as Theme) || 'dark',
  }),

  actions: {
    initializeTheme() {
      if (this.theme === 'dark') {
        document.documentElement.classList.add('dark-theme');
      } else {
        document.documentElement.classList.remove('dark-theme');
      }
    },


    toggleTheme() {
      this.theme = this.theme === 'light' ? 'dark' : 'light';
      localStorage.setItem('theme', this.theme);

      if (this.theme === 'dark') {
        document.documentElement.classList.add('dark-theme');
      } else {
        document.documentElement.classList.remove('dark-theme');
      }
    },
  },
});
