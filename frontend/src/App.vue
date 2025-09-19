<script setup lang="ts">
import TransactionList from '@/components/TransactionList.vue';
import TransactionForm from '@/components/TransactionForm.vue';
</script>

<template>
  <main>
    <h1>Transaction Tracker</h1>
    <TransactionForm />
    <hr />
    <TransactionList />
  </main>
</template>

<style scoped>
main {
  font-family: sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}
hr { margin: 20px 0; }
</style>
