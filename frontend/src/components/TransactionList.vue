<script setup lang="ts">
import { useTransactionStore } from '@/stores/transactions';
import { onMounted } from 'vue';
import OpenEyeIcon from '@/assets/icons/open-eye.svg?raw';
import ClosedEyeIcon from '@/assets/icons/closed-eye.svg?raw';

const store = useTransactionStore();

onMounted(() => {
  store.fetchTransactions();
});

type SortableKeys = 'Date' | 'Description' | 'Category' | 'Amount';

const sortOptions: { key: SortableKeys, label: string }[] = [
  { key: 'Date', label: 'Date' },
  { key: 'Description', label: 'Description' },
  { key: 'Category', label: 'Category' },
  { key: 'Amount', label: 'Amount' },
];
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h2>Transactions</h2>
      <div class="controls">
        <div class="sort-controls">
          <span class="sort-label">Sort by:</span>
          <button
            v-for="option in sortOptions"
            :key="option.key"
            class="sort-button"
            :class="{ active: store.sortKey === option.key }"
            @click="store.setSort(option.key)"
          >
            {{ option.label }}
            <span v-if="store.sortKey === option.key">{{ store.sortOrder === 'asc' ? '▲' : '▼' }}</span>
          </button>
        </div>
        <button @click="store.toggleAmountVisibility()" class="visibility-toggle" title="Toggle amount visibility">
          <span v-if="store.areAmountsVisible" v-html="OpenEyeIcon"></span>
          <span v-else v-html="ClosedEyeIcon"></span>
        </button>
      </div>
    </div>
    <div class="list-container">
      <div v-if="store.isLoading" class="feedback-message">Loading transactions...</div>
      <div v-else-if="store.error" class="feedback-message error">{{ store.error }}</div>
      <div v-else-if="store.sortedTransactions.length > 0" class="transaction-list">
        <div
          v-for="tx in store.sortedTransactions"
          :key="tx.ID"
          class="transaction-card"
          :class="tx.Type?.name.toLowerCase()"
        >
          <div class="tx-info">
            <p class="tx-description">{{ tx.Description }}</p>
            <p class="tx-details">
              <span>{{ tx.Date?.split('T')[0] }}</span>
              <span>•</span>
              <span>{{ tx.Category?.name }} > {{ tx.SubCategory?.Name }}</span>
            </p>
          </div>
          <div class="tx-amount">
            {{ tx.Currency?.Code }}
            {{ store.areAmountsVisible ? tx.Amount : '–––' }}
          </div>
        </div>
      </div>
      <div v-else class="feedback-message">No transactions found. Try adjusting your filters.</div>
    </div>
  </div>
</template>

<style scoped>
.card {
  background-color: var(--bg-card);
  border-radius: var(--radius);
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow-md);
  overflow: hidden;
}

.card-header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.card-header h2 {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
}

.controls {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.sort-controls {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: var(--bg-main);
  padding: 4px;
  border-radius: 6px;
}
.sort-label {
  font-size: 0.875rem;
  color: var(--text-secondary);
  margin-left: 0.5rem;
}
.sort-button {
  background: transparent;
  border: none;
  border-radius: 4px;
  padding: 0.5rem 0.75rem;
  font-family: inherit;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  color: var(--text-secondary);
  transition: background-color 0.2s, color 0.2s;
}
.sort-button.active {
  background-color: var(--bg-card);
  color: var(--text-primary);
  box-shadow: var(--shadow-sm);
}
.sort-button span {
  margin-left: 4px;
}

.visibility-toggle {
  background: none;
  border: 1px solid var(--border-color);
  width: 36px;
  height: 36px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s, border-color 0.2s;
}
.visibility-toggle:hover {
  background-color: var(--bg-main);
  border-color: #cdd5e0;
}
.visibility-toggle span :deep(svg) {
  width: 18px;
  height: 18px;
  color: var(--text-secondary);
}

.list-container {
  padding: 0.5rem 1.5rem 1.5rem;
}

.transaction-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  margin-top: 1rem;
}

.transaction-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  border-left-width: 4px;
  transition: box-shadow 0.2s;
}
.transaction-card:hover {
  box-shadow: var(--shadow-sm);
}

.transaction-card.credit { border-left-color: var(--color-credit); }
.transaction-card.debit { border-left-color: var(--color-debit); }
.transaction-card.refund { border-left-color: var(--color-refund); }

.tx-info {
  line-height: 1.4;
}
.tx-description {
  font-weight: 600;
  margin: 0 0 0.25rem 0;
  color: var(--text-primary);
}
.tx-details {
  font-size: 0.875rem;
  color: var(--text-secondary);
  margin: 0;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.tx-amount {
  font-size: 1.125rem;
  font-weight: 600;
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
}

.credit .tx-amount { color: var(--color-credit); }
.debit .tx-amount { color: var(--color-debit); }
.refund .tx-amount { color: var(--color-refund); }

.feedback-message {
  text-align: center;
  padding: 3rem 1rem;
  color: var(--text-secondary);
  font-size: 1rem;
}

.feedback-message.error {
  color: var(--color-debit);
  font-weight: 500;
}
</style>
