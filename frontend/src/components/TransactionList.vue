<script setup lang="ts">
import { useTransactionStore } from '@/stores/transactions';
import { onMounted } from 'vue';
import OpenEyeIcon from '@/assets/icons/open-eye.svg?raw';
import ClosedEyeIcon from '@/assets/icons/closed-eye.svg?raw';

const store = useTransactionStore();

onMounted(() => {
  store.fetchTransactions();
});
</script>

<template>
  <div class="transaction-list">
    <div class="header-controls">
      <h2>All Transactions</h2>
      <button @click="store.toggleAmountVisibility()" class="visibility-toggle" title="Toggle amount visibility">
        <span v-if="store.areAmountsVisible" v-html="OpenEyeIcon"></span>
        <span v-else v-html="ClosedEyeIcon"></span>
      </button>
    </div>
    <div v-if="store.isLoading">Loading...</div>
    <div v-else-if="store.error" class="error">{{ store.error }}</div>
    <table v-else-if="store.transactions.length > 0">
      <thead>
      <tr>
        <th>
          <button class="sort-button" @click="store.setSort('Date')">
            Date
            <span v-if="store.sortKey === 'Date'">{{ store.sortOrder === 'asc' ? '▲' : '▼' }}</span>
          </button>
        </th>
        <th>
          <button class="sort-button" @click="store.setSort('Description')">
            Description
            <span v-if="store.sortKey === 'Description'">{{ store.sortOrder === 'asc' ? '▲' : '▼' }}</span>
          </button>
        </th>
        <th>
          <button class="sort-button" @click="store.setSort('Category')">
            Category
            <span v-if="store.sortKey === 'Category'">{{ store.sortOrder === 'asc' ? '▲' : '▼' }}</span>
          </button>
        </th>
        <th>
          <button class="sort-button" @click="store.setSort('Amount')">
            Amount
            <span v-if="store.sortKey === 'Amount'">{{ store.sortOrder === 'asc' ? '▲' : '▼' }}</span>
          </button>
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="tx in store.sortedTransactions" :key="tx.ID">
        <td>{{ tx.Date?.split('T')[0] }}</td>
        <td>{{ tx.Description }}</td>
        <td>{{ tx.Category?.name }} > {{ tx.SubCategory?.Name }}</td>
        <td
          :class="{
              credit: tx.Type?.name === 'Credit',
              debit: tx.Type?.name === 'Debit',
              refund: tx.Type?.name === 'Refund',
            }"
        >
          {{ tx.Currency?.Code }}
          {{ store.areAmountsVisible ? tx.Amount : '---' }}
        </td>
      </tr>
      </tbody>
    </table>
    <div v-else>No transactions found.</div>
  </div>
</template>

<style scoped>
.header-controls {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}
.visibility-toggle {
  background: none;
  border: 1px solid #ccc;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
  padding: 0;
}
.visibility-toggle:hover {
  background-color: #f0f0f0;
}
.visibility-toggle span :deep(svg) {
  width: 20px;
  height: 20px;
}
table {
  width: 100%;
  border-collapse: collapse;
}
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}
th {
  background-color: #f2f2f2;
  border: 1px solid #ddd;
  padding: 0;
  text-align: left;
}
.sort-button {
  background: transparent;
  border: none;
  border-radius: 0;
  color: inherit;
  cursor: pointer;
  font-family: inherit;
  font-size: inherit;
  font-weight: bold;
  padding: 8px;
  text-align: left;
  width: 100%;
  transition: background-color 0.2s ease;
}
.sort-button:hover {
  background-color: #e6e6e6;
}
.sort-button span {
  margin-left: 8px;
  min-width: 12px;
  display: inline-block;
}
.credit {
  color: #28a745;
  font-weight: bold;
}
.debit {
  color: #dc3545;
}
.refund {
  color: #007bff;
}
.error {
  color: red;
  font-weight: bold;
}
</style>
