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
        <th>Date</th>
        <th>Description</th>
        <th>Category</th>
        <th>Amount</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="tx in store.transactions" :key="tx.ID">
        <td>{{ tx.Date.split('T')[0] }}</td>
        <td>{{ tx.Description }}</td>
        <td>{{ tx.Category.name }} > {{ tx.SubCategory.Name }}</td>
        <td
          :class="{
              credit: tx.Type.name === 'Credit',
              debit: tx.Type.name === 'Debit',
              refund: tx.Type.name === 'Refund',
            }"
        >
          {{ tx.Currency.Code }}
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
/* Style the SVG element inside the span rendered by v-html */
.visibility-toggle span :deep(svg) {
  width: 20px;
  height: 20px;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}
th {
  background-color: #f2f2f2;
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
