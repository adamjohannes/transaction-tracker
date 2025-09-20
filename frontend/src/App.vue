<script setup lang="ts">
import TransactionList from '@/components/TransactionList.vue';
import TransactionForm from '@/components/TransactionForm.vue';
import TransactionFilter from '@/components/TransactionFilter.vue';
</script>

<template>
  <header class="app-header">
    <div class="logo">
      <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"/><path d="M3 5v14a2 2 0 0 0 2 2h16v-5"/><path d="M18 12a2 2 0 0 0 0 4h4v-4h-4Z"/></svg>
      <h1>Transaction Tracker</h1>
    </div>
  </header>
  <main class="container">
    <div class="layout-grid">
      <aside class="controls-column">
        <TransactionForm />
        <TransactionFilter />
      </aside>
      <section class="list-column">
        <TransactionList />
      </section>
    </div>
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
  --primary-gradient: linear-gradient(45deg, #6a11cb 0%, #2575fc 100%);
  --primary-color-start: #6a11cb;
  --primary-color-end: #2575fc;
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

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo h1 {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  background: var(--primary-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 2rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: 350px 1fr;
  gap: 2rem;
  align-items: flex-start;
}

.controls-column {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.list-column {
  min-width: 0;
}

@media (max-width: 900px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
