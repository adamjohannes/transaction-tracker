<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';
import { useTransactionStore } from '@/stores/transactions';
import type { Transaction } from '@/services/api';

const store = useTransactionStore();

const newTransaction = ref<Transaction>({
  amount: '',
  date: new Date().toISOString().split('T')[0],
  type: '',
  essential: true,
  status: '',
  currency: '',
  category: '',
  subCategory: '',
  description: '',
});

watch(
  () => newTransaction.value.category,
  (newCategoryName) => {
    // When category changes, reset sub-category and fetch new options
    newTransaction.value.subCategory = '';
    store.fetchSubCategories(newCategoryName);
  }
);

onMounted(() => {
  store.fetchFormOptions(); // Fetch all static options when the component is mounted
});

async function handleSubmit() {
  if (!newTransaction.value.amount || !newTransaction.value.category) {
    alert('Please fill in all required fields.');
    return;
  }
  await store.addTransaction({ ...newTransaction.value });

  // Reset fields after submission
  newTransaction.value.amount = '';
  newTransaction.value.description = '';
  newTransaction.value.category = '';
  newTransaction.value.subCategory = '';
}
</script>

<template>
  <div class="transaction-form">
    <h2>Add New Transaction</h2>
    <div v-if="store.isLoadingOptions" class="loading">Loading form data...</div>
    <form v-else @submit.prevent="handleSubmit">
      <div class="form-grid">
        <input v-model="newTransaction.description" placeholder="Description" required />
        <input v-model.number="newTransaction.amount" type="number" step="0.01" placeholder="Amount" required />
        <input v-model="newTransaction.date" type="date" required />

        <select v-model="newTransaction.category" required>
          <option disabled value="">Select a Category</option>
          <option v-for="cat in store.categories" :key="cat.id" :value="cat.name">
            {{ cat.name }}
          </option>
        </select>

        <select v-model="newTransaction.subCategory" :disabled="!newTransaction.category || store.subCategories.length === 0">
          <option disabled value="">Select a Sub-Category</option>
          <option v-for="subCat in store.subCategories" :key="subCat.ID" :value="subCat.Name">
            {{ subCat.Name }}
          </option>
        </select>

        <select v-model="newTransaction.type" required>
          <option disabled value="">Select a Type</option>
          <option v-for="t in store.transactionTypes" :key="t.id" :value="t.name">{{ t.name }}</option>
        </select>

        <select v-model="newTransaction.status" required>
          <option disabled value="">Select a Status</option>
          <option v-for="s in store.statuses" :key="s.id" :value="s.name">{{ s.name }}</option>
        </select>

        <select v-model="newTransaction.currency" required>
          <option disabled value="">Select a Currency</option>
          <option v-for="c in store.currencies" :key="c.Code" :value="c.Code">{{ c.Code }}</option>
        </select>
      </div>

      <button type="submit" :disabled="store.isLoading">
        {{ store.isLoading ? 'Adding...' : 'Add Transaction' }}
      </button>
    </form>
    <div v-if="store.error" class="error-message">{{ store.error }}</div>
  </div>
</template>

<style scoped>
.transaction-form {
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
}
.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 15px;
  margin-bottom: 20px;
}
input, select {
  padding: 10px;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}
select:disabled {
  background-color: #e9ecef;
  cursor: not-allowed;
}
button {
  padding: 10px 20px;
  font-size: 1rem;
  color: white;
  background-color: #42b983;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}
button:disabled {
  background-color: #a5d6b8;
  cursor: wait;
}
button:hover:not(:disabled) {
  background-color: #369a6e;
}
.loading, .error-message {
  text-align: center;
  padding: 10px;
  margin-top: 10px;
}
.error-message {
  color: #d9534f;
  background-color: #f2dede;
  border: 1px solid #ebccd1;
  border-radius: 4px;
}
</style>
