<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';
import { useTransactionStore } from '@/stores/transactions';

const store = useTransactionStore();

const newTransaction = ref({
  Amount: '',
  Date: new Date().toISOString().split('T')[0],
  Type: '',
  Essential: true,
  Status: '',
  Currency: '',
  Category: '',
  SubCategory: '',
  Description: '',
});

watch(
  () => newTransaction.value.Category,
  (newCategoryName) => {
    newTransaction.value.SubCategory = '';
    store.fetchSubCategories(newCategoryName);
  }
);

onMounted(() => {
  store.fetchFormOptions();
});

async function handleSubmit() {
  if (!newTransaction.value.Amount || !newTransaction.value.Category) {
    alert('Please fill in all required fields.');
    return;
  }
  await store.addTransaction({ ...newTransaction.value });

  newTransaction.value.Amount = '';
  newTransaction.value.Description = '';
  newTransaction.value.Category = '';
  newTransaction.value.SubCategory = '';
}
</script>

<template>
  <div class="card">
    <h2 class="card-header">Add Transaction</h2>
    <div v-if="store.isLoadingOptions" class="loading">Loading form data...</div>
    <form v-else @submit.prevent="handleSubmit" class="form-content">
      <div class="form-group">
        <label for="description">Description</label>
        <input id="description" v-model="newTransaction.Description" placeholder="e.g., Coffee with friends" required />
      </div>

      <div class="form-group">
        <label for="amount">Amount</label>
        <input id="amount" v-model="newTransaction.Amount" type="number" step="0.01" placeholder="0.00" required />
      </div>

      <div class="form-group">
        <label for="date">Date</label>
        <input id="date" v-model="newTransaction.Date" type="date" required />
      </div>

      <div class="form-group">
        <label for="category">Category</label>
        <select id="category" v-model="newTransaction.Category" required>
          <option disabled value="">Select a Category</option>
          <option v-for="cat in store.categories" :key="cat.id" :value="cat.name">
            {{ cat.name }}
          </option>
        </select>
      </div>

      <div class="form-group">
        <label for="subcategory">Sub-Category</label>
        <select id="subcategory" v-model="newTransaction.SubCategory" :disabled="!newTransaction.Category || store.subCategories.length === 0">
          <option disabled value="">Select a Sub-Category</option>
          <option v-for="subCat in store.subCategories" :key="subCat.ID" :value="subCat.Name">
            {{ subCat.Name }}
          </option>
        </select>
      </div>

      <div class="form-group">
        <label for="type">Type</label>
        <select id="type" v-model="newTransaction.Type" required>
          <option disabled value="">Select a Type</option>
          <option v-for="t in store.transactionTypes" :key="t.id" :value="t.name">{{ t.name }}</option>
        </select>
      </div>

      <div class="form-group">
        <label for="status">Status</label>
        <select id="status" v-model="newTransaction.Status" required>
          <option disabled value="">Select a Status</option>
          <option v-for="s in store.statuses" :key="s.id" :value="s.name">{{ s.name }}</option>
        </select>
      </div>

      <div class="form-group">
        <label for="currency">Currency</label>
        <select id="currency" v-model="newTransaction.Currency" required>
          <option disabled value="">Select a Currency</option>
          <option v-for="c in store.currencies" :key="c.Code" :value="c.Code">{{ c.Code }}</option>
        </select>
      </div>


      <button type="submit" class="submit-button" :disabled="store.isLoading">
        {{ store.isLoading ? 'Adding...' : 'Add Transaction' }}
      </button>
    </form>
    <div v-if="store.error" class="error-message">{{ store.error }}</div>
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
  font-size: 1.25rem;
  font-weight: 600;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border-color);
  margin: 0;
}

.form-content {
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-secondary);
  margin-bottom: 0.5rem;
}

input, select {
  width: 100%;
  padding: 0.75rem;
  font-size: 1rem;
  font-family: inherit;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  background-color: var(--bg-main);
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}

input:focus, select:focus {
  outline: none;
  border-color: var(--primary-color-end);
  box-shadow: 0 0 0 3px rgba(37, 117, 252, 0.2);
}

select:disabled {
  background-color: #e9ecef;
  cursor: not-allowed;
  opacity: 0.7;
}

.submit-button {
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  color: white;
  background: var(--primary-gradient);
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
  margin-top: 1rem;
}

.submit-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.submit-button:disabled {
  opacity: 0.6;
  cursor: wait;
}

.loading, .error-message {
  text-align: center;
  padding: 1rem 1.5rem;
}

.error-message {
  color: var(--color-debit);
  background-color: #fef2f2;
  font-weight: 500;
}
</style>
