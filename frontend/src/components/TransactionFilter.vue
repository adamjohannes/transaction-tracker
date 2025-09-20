<script setup lang="ts">
import { ref, watch } from 'vue';
import { useTransactionStore } from '@/stores/transactions';

const store = useTransactionStore();

// Local state for filter inputs, initialized from the store's state
const localFilters = ref({ ...store.filters });

// Watch for changes in localFilters and update the store
watch(localFilters, (newFilters) => {
  store.updateFilters(newFilters);
}, { deep: true });

// Function to reset filters
function clearFilters() {
  store.clearFilters();
  localFilters.value = { ...store.filters };
}
</script>

<template>
  <div class="filter-panel">
    <h3>Filter Transactions</h3>
    <div class="filter-grid">
      <div class="filter-group">
        <label for="startDate">Start Date</label>
        <input id="startDate" v-model="localFilters.startDate" type="date" />
      </div>
      <div class="filter-group">
        <label for="endDate">End Date</label>
        <input id="endDate" v-model="localFilters.endDate" type="date" />
      </div>

      <div class="filter-group">
        <label for="categories">Categories</label>
        <select id="categories" v-model="localFilters.categories" multiple>
          <option v-for="cat in store.categories" :key="cat.id" :value="cat.name">
            {{ cat.name }}
          </option>
        </select>
      </div>

      <div class="filter-group">
        <label for="subCategories">Sub-Categories</label>
        <select id="subCategories" v-model="localFilters.subCategories" multiple :disabled="store.subCategories.length === 0">
          <option v-for="subCat in store.subCategories" :key="subCat.ID" :value="subCat.Name">
            {{ subCat.Name }}
          </option>
        </select>
      </div>

      <div class="filter-group">
        <label for="types">Type</label>
        <select id="types" v-model="localFilters.types" multiple>
          <option v-for="t in store.transactionTypes" :key="t.id" :value="t.name">{{ t.name }}</option>
        </select>
      </div>

      <div class="filter-group">
        <label for="statuses">Status</label>
        <select id="statuses" v-model="localFilters.statuses" multiple>
          <option v-for="s in store.statuses" :key="s.id" :value="s.name">{{ s.name }}</option>
        </select>
      </div>
    </div>
    <button @click="clearFilters" class="clear-button">Clear Filters</button>
  </div>
</template>

<style scoped>
.filter-panel {
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
  margin-bottom: 20px;
}
.filter-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 15px;
  margin-top: 10px;
}
.filter-group {
  display: flex;
  flex-direction: column;
}
.filter-group label {
  margin-bottom: 5px;
  font-weight: bold;
  font-size: 0.9rem;
}
input, select {
  padding: 10px;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}
select[multiple] {
  height: 100px;
  background-color: white;
}
.clear-button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}
.clear-button:hover {
  background-color: #c82333;
}
</style>
