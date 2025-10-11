<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useThemeStore } from '@/stores/theme';
import TransactionIcon from '@/assets/icons/transaction.svg?raw';

const authStore = useAuthStore();
const themeStore = useThemeStore();
themeStore.initializeTheme();
</script>

<template>
  <header class="app-header">
    <div class="logo-area">
      <div class="logo">
        <span v-html="TransactionIcon"></span>
        <h1>Transaction Tracker</h1>
      </div>
      <nav class="main-nav">
        <template v-if="authStore.isLoggedIn">
          <RouterLink to="/">Transactions</RouterLink>
          <RouterLink to="/dashboard">Dashboard</RouterLink>
          <a href="#" @click.prevent="authStore.logout()">Logout</a>
        </template>
        <template v-else>
          <RouterLink to="/login">Login</RouterLink>
          <RouterLink to="/register">Register</RouterLink>
        </template>
      </nav>
      <button @click="themeStore.toggleTheme" class="theme-toggle" :title="'Switch to ' + (themeStore.theme === 'light' ? 'Dark' : 'Light') + ' Mode'">
        {{ themeStore.theme === 'light' ? '☀️' : '🌙' }}
      </button>
    </div>
  </header>
  <main>
    <RouterView />
  </main>
</template>

<style>
/* --- Light Theme --- */
:root {
  --font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;

  /* Backgrounds */
  --bg-main: #f8f9fa;
  --bg-card: #ffffff;

  /* Text */
  --text-primary: #212529;
  --text-secondary: #6c757d;

  /* Borders & Shadows */
  --border-color: #dee2e6;
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);

  /* Primary / Accent Colors */
  --primary-gradient: linear-gradient(45deg, #4f46e5 0%, #6366f1 100%);
  --primary-color-start: #4f46e5;
  --primary-color-end: #6366f1;

  /* Status Colors */
  --color-credit: #10b981;
  --color-debit: #ef4444;
  --color-refund: #3b82f6;

  /* UI Elements */
  --radius: 8px;
}

/* --- Dark Theme --- */
html.dark-theme {
  /* Backgrounds */
  --bg-main: #060e0a;
  --bg-card: #0c1f16;

  /* Text */
  --text-primary: #e8f5e9;
  --text-secondary: #7e968a;

  /* Borders & Shadows */
  --border-color: #20402c;
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.25);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.35), 0 2px 4px -1px rgba(0, 0, 0, 0.25);

  /* Primary / Accent Colors (Vibrant Green) */
  --primary-gradient: linear-gradient(45deg, #16a34a 0%, #22c55e 100%);
  --primary-color-start: #16a34a;
  --primary-color-end: #22c55e;

  /* Status Colors */
  --color-credit: #4ade80;
  --color-debit: #f87171;
  --color-refund: #60a5fa;

  /* Browser UI hint */
  color-scheme: dark;
}

html.dark-theme .clear-button:hover {
  background-color: var(--border-color);
}
html.dark-theme select:disabled {
  background-color: var(--border-color);
  opacity: 0.5;
}
html.dark-theme .error-message {
  background-color: rgba(248, 113, 113, 0.1);
  color: var(--color-debit);
}
html.dark-theme .select-prompt,
html.dark-theme .breakdown-container,
html.dark-theme .chart-container {
  background-color: var(--bg-card);
}

/* --- Base Styles --- */
body {
  font-family: var(--font-sans);
  background-color: var(--bg-main);
  color: var(--text-primary);
  margin: 0;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.app-header {
  background-color: var(--bg-card);
  padding: 1rem 2rem;
  border-bottom: 1px solid var(--border-color);
  box-shadow: var(--shadow-sm);
  margin-bottom: 2rem;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.logo-area {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo h1 {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
}

.logo span :deep(svg) {
  width: 30px;
  height: 30px;
  color: var(--primary-color-end);
  transition: color 0.3s ease;
}

.main-nav {
  display: flex;
  gap: 0.5rem;
  margin-left: 2rem;
}

.main-nav a {
  padding: 0.5rem 1rem;
  text-decoration: none;
  font-weight: 500;
  color: var(--text-secondary);
  border-radius: 6px;
  transition: background-color 0.2s, color 0.2s;
}

.main-nav a:hover {
  background-color: var(--bg-main);
  color: var(--text-primary);
}

.main-nav a.router-link-exact-active {
  background-image: var(--primary-gradient);
  color: white;
}

.theme-toggle {
  margin-left: auto;
  background: none;
  border: 1px solid var(--border-color);
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.25rem;
  transition: background-color 0.2s, border-color 0.2s;
}

.theme-toggle:hover {
  background-color: var(--bg-main);
}
</style>
