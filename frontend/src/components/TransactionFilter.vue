<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useTransactionStore } from '@/stores/transactions';
import BaseMultiCombobox from './BaseMultiCombobox.vue';

type Category = { name: string };
type SubCategory = { ID?: string | number; Name: string };
type TxType = { name: string };
type Status = { name: string };

const store = useTransactionStore();

const localFilters = ref({ ...store.filters });
const localSubCategories = ref<SubCategory[]>([]);

onMounted(() => {
  if (!store.categories.length || !store.statuses.length || !store.transactionTypes.length) {
    store.fetchFormOptions();
  }
});

watch(localFilters, (next) => store.updateFilters(next), { deep: true });

watch(
  () => localFilters.value.categories as string[],
  async (names) => {
    if (!names?.length) {
      localSubCategories.value = [];
      localFilters.value.subCategories = [];
      return;
    }
    const seen = new Map<string, SubCategory>();
    for (const n of names) {
      await store.fetchSubCategories(n);
      for (const sc of store.subCategories as SubCategory[]) {
        const key = String(sc.ID ?? sc.Name);
        if (!seen.has(key)) seen.set(key, sc);
      }
    }
    localSubCategories.value = Array.from(seen.values());
    const valid = new Set(localSubCategories.value.map(sc => sc.Name));
    localFilters.value.subCategories = (localFilters.value.subCategories as string[]).filter(n => valid.has(n));
  },
  { deep: false }
);

const categoryOptions = computed(() =>
  (store.categories as Category[]).map(c => ({ label: c.name, value: c.name }))
);

const subCategoryOptions = computed(() =>
  (localSubCategories.value.length ? localSubCategories.value : (store.subCategories as SubCategory[]))
    .map(sc => ({ label: sc.Name, value: sc.Name }))
);

const typeOptions = computed(() =>
  (store.transactionTypes as TxType[]).map(t => ({ label: t.name, value: t.name }))
);

const statusOptions = computed(() =>
  (store.statuses as Status[]).map(s => ({ label: s.name, value: s.name }))
);

function clearFilters() {
  store.clearFilters();
  localFilters.value = { ...store.filters };
  localSubCategories.value = [];
}
</script>

<template>
  <div class="card">
    <h3 class="card-header">Filter Transactions</h3>
    <div class="filter-content">
      <div class="date-grid">
        <div class="form-group">
          <label for="startDate">Start Date</label>
          <input id="startDate" v-model="(localFilters as any).startDate" type="date" />
        </div>
        <div class="form-group">
          <label for="endDate">End Date</label>
          <input id="endDate" v-model="(localFilters as any).endDate" type="date" />
        </div>
      </div>

      <div class="form-group">
        <label>Categories</label>
        <BaseMultiCombobox
          v-model="(localFilters as any).categories"
          :options="categoryOptions"
          placeholder="Pick categories"
          search-placeholder="Search categories..."
          :disabled="store.isLoadingOptions"
        />
      </div>

      <div class="form-group">
        <label>Sub-Categories</label>
        <BaseMultiCombobox
          v-model="(localFilters as any).subCategories"
          :options="subCategoryOptions"
          placeholder="Pick sub-categories"
          search-placeholder="Search sub-categories..."
          :disabled="store.isLoadingOptions || !((localFilters as any).categories?.length)"
        />
      </div>

      <div class="form-group">
        <label>Type</label>
        <BaseMultiCombobox
          v-model="(localFilters as any).types"
          :options="typeOptions"
          placeholder="Pick types"
          :disabled="store.isLoadingOptions"
        />
      </div>

      <div class="form-group">
        <label>Status</label>
        <BaseMultiCombobox
          v-model="(localFilters as any).statuses"
          :options="statusOptions"
          placeholder="Pick statuses"
          :disabled="store.isLoadingOptions"
        />
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
  overflow: visible;
}

.card-header {
  font-size: 1.25rem;
  font-weight: 600;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border-color);
  margin: 0;
}

.filter-content { padding: 1.5rem; display: flex; flex-direction: column; gap: 1rem; }

.date-grid { display: grid; grid-template-columns: minmax(0,1fr) minmax(0,1fr); gap: 1rem; }
.form-group { min-width: 0; }

*,
*::before,
*::after { box-sizing: border-box; }

.form-group label {
  display: block; font-size: 0.875rem; font-weight: 500; color: var(--text-secondary); margin-bottom: 0.5rem;
}

.date-grid input[type="date"] {
  display: block;
  width: 100%;
  min-width: 0;
}

input {
  width: 100%;
  padding: 0.75rem;
  font-size: 1rem;
  font-family: inherit;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  background-color: var(--bg-main);
}

input:focus {
  outline: none;
  border-color: var(--primary-color-end);
  box-shadow: 0 0 0 3px rgba(37,117,252,.2);
}

.clear-button {
  margin-top: 1rem;
  padding: .75rem 1.5rem;
  background: transparent;
  color: var(--text-secondary);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
}
.clear-button:hover { background-color: var(--bg-main); color: var(--text-primary); }

@media (max-width: 900px) { .date-grid { grid-template-columns: 1fr; } }
</style>
