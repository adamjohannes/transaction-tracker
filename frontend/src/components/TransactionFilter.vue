<script setup lang="ts">
import { ref, watch } from 'vue';
import { useTransactionStore } from '@/stores/transactions';

const store = useTransactionStore();
const localFilters = ref({ ...store.filters });

watch(localFilters, (newFilters) => {
  store.updateFilters(newFilters);
}, { deep: true });

function clearFilters() {
  store.clearFilters();
  localFilters.value = { ...store.filters };
}
</script>

<template>
  <div class="card">
    <h3 class="card-header">Filter Transactions</h3>
    <div class="filter-content">
      <div class="date-grid">
        <div class="form-group">
          <label for="startDate">Start Date</label>
          <input id="startDate" v-model="localFilters.startDate" type="date" />
        </div>
        <div class="form-group">
          <label for="endDate">End Date</label>
          <input id="endDate" v-model="localFilters.endDate" type="date" />
        </div>
      </div>
      <div class="form-group">
        <label for="categories">Categories</label>
        <select id="categories" v-model="localFilters.categories" multiple>
          <option v-for="cat in store.categories" :key="cat.id" :value="cat.name">
            {{ cat.name }}
          </option>
        </select>
      </div>
      <div class="form-group">
        <label for="subCategories">Sub-Categories</label>
        <select id="subCategories" v-model="localFilters.subCategories" multiple :disabled="store.subCategories.length === 0">
          <option v-for="subCat in store.subCategories" :key="subCat.ID" :value="subCat.Name">
            {{ subCat.Name }}
          </option>
        </select>
      </div>
      <div class="form-group">
        <label for="types">Type</label>
        <select id="types" v-model="localFilters.types" multiple>
          <option v-for="t in store.transactionTypes" :key="t.id" :value="t.name">{{ t.name }}</option>
        </select>
      </div>

      <div class="form-group">
        <label for="statuses">Status</label>
        <select id="statuses" v-model="localFilters.statuses" multiple>
          <option v-for="s in store.statuses" :key="s.id" :value="s.name">{{ s.name }}</option>
        </select>
      </div>
      <button @click="clearFilters" class="clear-button">Clear All Filters</button>
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
  font-size: 1.25rem;
  font-weight: 600;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border-color);
  margin: 0;
}

.filter-content {
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.date-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
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

select[multiple] {
  min-height: 100px;
}

.clear-button {
  margin-top: 1rem;
  padding: 0.75rem 1.5rem;
  background-color: transparent;
  color: var(--text-secondary);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  transition: background-color 0.2s, color 0.2s;
}

.clear-button:hover {
  background-color: var(--bg-main);
  color: var(--text-primary);
}
</style>
