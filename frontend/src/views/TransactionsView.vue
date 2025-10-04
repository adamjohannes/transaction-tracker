<script setup lang="ts">
import TransactionFilter from '@/components/TransactionFilter.vue';
import TransactionForm from '@/components/TransactionForm.vue';
import TransactionList from '@/components/TransactionList.vue';
</script>

<template>
  <div class="layout-grid">
    <aside class="form-column">
      <TransactionForm />
    </aside>

    <section class="list-column">
      <TransactionList />
    </section>

    <aside class="filter-column">
      <TransactionFilter />
    </aside>
  </div>
</template>

<style scoped>
.layout-grid {
  display: grid;
  grid-template-columns: 350px 1fr 350px;
  gap: 2rem;
  align-items: flex-start;
}

.form-column {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.list-column {
  min-width: 0;
}

.filter-column {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

@media (max-width: 1200px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
