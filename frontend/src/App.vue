<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import TransactionIcon from '@/assets/icons/transaction.svg?raw';
</script>

<template>
  <header class="app-header">
    <div class="logo-area">
      <div class="logo">
        <span v-html="TransactionIcon"></span>
        <h1>Transaction Tracker</h1>
      </div>
      <nav class="main-nav">
        <RouterLink to="/">Transactions</RouterLink>
        <RouterLink to="/dashboard">Dashboard</RouterLink>
      </nav>
    </div>
  </header>
  <main class="container">
    <RouterView />
  </main>
</template>

<style>
:root {
  --font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
  --bg-main: #f3f4f7;
  --bg-card: #ffffff;
  --text-primary: #1a202c;
  --text-secondary: #5a6a7a;
  --border-color: #e2e8f0;
  --primary-gradient: linear-gradient(45deg, #8f5300 0%, #ffc500 100%);
  --primary-color-start: #8f5300;
  --primary-color-end: #ffc500;
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  --color-credit: #10b981;
  --color-debit: #ef4444;
  --color-refund: #3b82f6;
  --radius: 8px;
}

body {
  font-family: var(--font-sans);
  background-color: var(--bg-main);
  color: var(--text-primary);
  margin: 0;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.app-header {
  background-color: var(--bg-card);
  padding: 1rem 2rem;
  border-bottom: 1px solid var(--border-color);
  box-shadow: var(--shadow-sm);
  margin-bottom: 2rem;
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

.logo span svg {
  width: 30px;
  height: 30px;
  color: #17B978;
}

.container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 2rem;
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
</style>
