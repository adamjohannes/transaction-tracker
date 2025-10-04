<script setup lang="ts">
import TransactionFilter from '@/components/TransactionFilter.vue';
import TransactionForm from '@/components/TransactionForm.vue';
import TransactionList from '@/components/TransactionList.vue';
</script>

<template>
  <div class="layout-grid">
    <aside class="controls-column">
      <TransactionForm />
      <TransactionFilter />
    </aside>
    <section class="list-column">
      <TransactionList />
    </section>
  </div>
</template>

<style scoped>
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
