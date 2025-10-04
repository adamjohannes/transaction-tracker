<script setup lang="ts">
import { ref, computed } from 'vue';
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

ChartJS.register(Title, Tooltip, Legend, ArcElement);

const store = useTransactionStore();
type ChartType = 'debit' | 'credit' | 'refund';
const activeType = ref<ChartType>('debit');

// Base options for the chart
const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
  },
};

const pieColors = [
  '#3b82f6', // Blue
  '#f97316', // Orange
  '#ec4899', // Pink
  '#14b8a6', // Teal
  '#8b5cf6', // Purple
  '#eab308', // Yellow
  '#ef4444', // Red
  '#0ea5e9', // Sky Blue
  '#10b981', // Green
  '#64748b', // Slate
];

function sortChartData(chartData: { labels: string[], data: number[] }) {
  const zipped = chartData.labels.map((label, i) => ({
    label,
    value: chartData.data[i]
  }));

  zipped.sort((a, b) => b.value - a.value);

  const sortedLabels = zipped.map(item => item.label);
  const sortedData = zipped.map(item => item.value);

  return { labels: sortedLabels, data: sortedData };
}

const createChartDataObject = (chartData: { labels: string[], data: number[] }): ChartData<'pie'> => ({
  labels: chartData.labels,
  datasets: [{
    backgroundColor: pieColors,
    data: chartData.data,
    borderWidth: 1,
    borderColor: 'var(--bg-card)',
  }],
});

const debitData = computed(() => {
  const sorted = sortChartData(store.debitChartData);
  return createChartDataObject(sorted);
});
const creditData = computed(() => {
  const sorted = sortChartData(store.creditChartData);
  return createChartDataObject(sorted);
});
const refundData = computed(() => {
  const sorted = sortChartData(store.refundChartData);
  return createChartDataObject(sorted);
});

const activeChartData = computed(() => {
  switch (activeType.value) {
    case 'credit': return creditData.value;
    case 'refund': return refundData.value;
    default: return debitData.value;
  }
});

const totalAmount = computed(() => {
  return activeChartData.value.datasets[0].data.reduce((sum, value) => sum + value, 0);
});

const categorySummary = computed(() => {
  const total = totalAmount.value;
  if (total === 0) return [];

  const labels = activeChartData.value.labels ?? [];
  const data = activeChartData.value.datasets[0].data ?? [];

  return labels.map((label, index) => ({
    name: label,
    color: pieColors[index % pieColors.length],
    amount: data[index],
    percentage: ((data[index] / total) * 100).toFixed(1),
  }));
});
</script>

<template>
  <div class="dashboard-card">
    <div class="card-header">
      <div class="header-text">
        <h2>Category Breakdown</h2>
        <p>An overview of your transactions by category.</p>
      </div>
      <div class="segmented-control">
        <button @click="activeType = 'debit'" :class="{ active: activeType === 'debit' }">Debits</button>
        <button @click="activeType = 'credit'" :class="{ active: activeType === 'credit' }">Credits</button>
        <button @click="activeType = 'refund'" :class="{ active: activeType === 'refund' }">Refunds</button>
      </div>
    </div>

    <div class="card-body">
      <div v-if="totalAmount > 0" class="dashboard-content">
        <div class="chart-area">
          <Transition name="fade" mode="out-in">
            <Pie :data="activeChartData" :options="chartOptions" :key="activeType" />
          </Transition>
        </div>
        <div class="summary-area">
          <div class="total-display">
            <span>Total {{ activeType }}s</span>
            <strong>{{ new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(totalAmount) }}</strong>
          </div>
          <h3>Breakdown</h3>
          <ul>
            <li v-for="item in categorySummary" :key="item.name">
              <div class="category-info">
                <span class="color-dot" :style="{ backgroundColor: item.color }"></span>
                {{ item.name }}
              </div>
              <div class="category-stats">
                <span class="percentage">{{ item.percentage }}%</span>
                <span class="amount">{{ new Intl.NumberFormat().format(item.amount) }}</span>
              </div>
            </li>
          </ul>
        </div>
      </div>
      <div v-else class="no-data-view">
        <div class="icon">📊</div>
        <h3>No Data to Display</h3>
        <p>There are no {{ activeType }} transactions for the current filters.</p>
        <p>Try adjusting the filters on the Transactions page.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-card {
  background-color: var(--bg-card);
  border-radius: var(--radius);
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow-md);
  overflow: hidden;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
  padding: 1.5rem;
  border-bottom: 1px solid var(--border-color);
}

.header-text h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.header-text p {
  margin: 0.25rem 0 0;
  color: var(--text-secondary);
}

.segmented-control {
  display: flex;
  background-color: var(--bg-main);
  border-radius: 6px;
  padding: 4px;
}

.segmented-control button {
  background: transparent;
  border: none;
  border-radius: 4px;
  padding: 0.5rem 1rem;
  font-family: inherit;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  color: var(--text-secondary);
  transition: background-color 0.2s, color 0.2s, box-shadow 0.2s;
}

.segmented-control button.active {
  background-color: var(--bg-card);
  color: var(--text-primary);
  box-shadow: var(--shadow-sm);
}

.card-body {
  padding: 1.5rem;
}

.dashboard-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  align-items: center;
  min-height: 400px;
}

.chart-area {
  position: relative;
  height: 400px;
}

.total-display {
  text-align: left;
  margin-bottom: 1.5rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid var(--border-color);
}

.total-display span {
  font-size: 1rem;
  color: var(--text-secondary);
  text-transform: capitalize;
}

.total-display strong {
  display: block;
  font-size: 2.25rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.2;
}

.summary-area h3 {
  margin-top: 0;
  margin-bottom: 1rem;
  font-weight: 600;
}

.summary-area ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.summary-area li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.9rem;
}

.category-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-weight: 500;
}

.color-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.category-stats {
  display: flex;
  align-items: baseline;
  gap: 0.75rem;
}

.percentage {
  font-weight: 600;
  color: var(--text-primary);
  background-color: var(--bg-main);
  padding: 0.1rem 0.4rem;
  border-radius: 4px;
  min-width: 45px;
  text-align: right;
}

.amount {
  font-size: 0.85rem;
  color: var(--text-secondary);
  font-variant-numeric: tabular-nums;
}

.no-data-view {
  text-align: center;
  padding: 3rem 1rem;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.no-data-view .icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.no-data-view h3 {
  font-size: 1.25rem;
  color: var(--text-primary);
  margin-bottom: 0.5rem;
}

.no-data-view p {
  color: var(--text-secondary);
  margin: 0.25rem;
  max-width: 300px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@media (max-width: 900px) {
  .dashboard-content {
    grid-template-columns: 1fr;
    gap: 3rem;
  }
}
</style>
