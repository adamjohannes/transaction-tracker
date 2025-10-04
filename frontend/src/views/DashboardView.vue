<script setup lang="ts">
import { computed } from 'vue';
import { useTransactionStore } from '@/stores/transactions';
import { Pie } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  ArcElement,
  type ChartData,
} from 'chart.js';

// Register the necessary components for Chart.js to work
ChartJS.register(Title, Tooltip, Legend, ArcElement);

const store = useTransactionStore();

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
};

const pieColors = [
  '#FFC500', '#FF6384', '#36A2EB', '#FF9F40', '#4BC0C0',
  '#9966FF', '#C9CBCF', '#FFCD56', '#F7464A', '#46BFBD'
];

// Computed properties to format data for each chart
const debitData = computed<ChartData<'pie'>>(() => ({
  labels: store.debitChartData.labels,
  datasets: [{
    label: 'Debits',
    backgroundColor: pieColors,
    data: store.debitChartData.data,
  }],
}));

const creditData = computed<ChartData<'pie'>>(() => ({
  labels: store.creditChartData.labels,
  datasets: [{
    label: 'Credits',
    backgroundColor: pieColors,
    data: store.creditChartData.data,
  }],
}));

const refundData = computed<ChartData<'pie'>>(() => ({
  labels: store.refundChartData.labels,
  datasets: [{
    label: 'Refunds',
    backgroundColor: pieColors,
    data: store.refundChartData.data,
  }],
}));
</script>

<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h2>Category Breakdown</h2>
      <p class="subtitle">
        Visual representation of amounts per category, reflecting the currently active filters.
      </p>
    </div>
    <div class="chart-grid">
      <div class="chart-container card" v-if="debitData.datasets[0].data.length > 0">
        <h3>Debits by Category</h3>
        <Pie :data="debitData" :options="chartOptions" />
      </div>

      <div class="chart-container card" v-if="creditData.datasets[0].data.length > 0">
        <h3>Credits by Category</h3>
        <Pie :data="creditData" :options="chartOptions" />
      </div>

      <div class="chart-container card" v-if="refundData.datasets[0].data.length > 0">
        <h3>Refunds by Category</h3>
        <Pie :data="refundData" :options="chartOptions" />
      </div>
    </div>
    <div v-if="!debitData.datasets[0].data.length && !creditData.datasets[0].data.length && !refundData.datasets[0].data.length" class="no-data-message card">
      <p>No data available for the current filters.</p>
      <p>Try adjusting the filters on the Transactions page to see a breakdown.</p>
    </div>
  </div>
</template>

<style scoped>
.dashboard-header {
  margin-bottom: 2rem;
}

.dashboard h2 {
  font-size: 1.75rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.subtitle {
  margin-top: 0;
  color: var(--text-secondary);
  font-size: 1rem;
}

.chart-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
}

.chart-container {
  padding: 1.5rem;
  height: 450px;
  display: flex;
  flex-direction: column;
}

.chart-container h3 {
  text-align: center;
  margin-top: 0;
  margin-bottom: 1.5rem;
  font-weight: 600;
  font-size: 1.1rem;
  color: var(--text-primary);
}

.card {
  background-color: var(--bg-card);
  border-radius: var(--radius);
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow-md);
}

.no-data-message {
  padding: 3rem 1.5rem;
  text-align: center;
  color: var(--text-secondary);
}
.no-data-message p {
  margin: 0.5rem 0;
}
</style>
