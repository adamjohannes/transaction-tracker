<script setup lang="ts">
import { computed } from 'vue';
import { Pie } from 'vue-chartjs';
import { Chart as ChartJS, Title, Tooltip, Legend, ArcElement, type ChartData, type ChartOptions } from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, ArcElement);

const props = defineProps<{
  title: string;
  chartData: ChartData<'pie'>;
  transactionType: 'debit' | 'credit' | 'refund';
}>();

const chartOptions = computed<ChartOptions<'pie'>>(() => ({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { display: false },
    title: {
      display: true,
      text: props.title,
      font: {
        size: 16,
        weight: 'bold',
      },
      padding: { bottom: 16 }
    }
  },
}));

const pieColors = [
  '#3b82f6', '#f97316', '#ec4899', '#14b8a6', '#8b5cf6',
  '#eab308', '#ef4444', '#0ea5e9', '#10b981', '#64748b'
];

const totalAmount = computed(() => {
  return props.chartData.datasets[0]?.data.reduce((sum, value) => sum + (value as number), 0) ?? 0;
});

const categorySummary = computed(() => {
  if (totalAmount.value === 0) return [];
  const labels = props.chartData.labels ?? [];
  const data = props.chartData.datasets[0]?.data ?? [];

  const zipped = (labels as string[]).map((label, i) => ({
    label,
    value: data[i] as number,
  }));

  zipped.sort((a, b) => b.value - a.value);

  return zipped.map((item, index) => ({
    name: item.label,
    color: pieColors[index % pieColors.length],
    amount: item.value,
    percentage: ((item.value / totalAmount.value) * 100).toFixed(1),
  }));
});
</script>

<template>
  <div class="breakdown-container">
    <div v-if="totalAmount > 0" class="content-wrapper">
      <div class="chart-area">
        <Pie :data="chartData" :options="chartOptions" />
      </div>
      <div class="summary-area">
        <div class="total-display">
          <span>Total {{ transactionType }}s</span>
          <strong>{{ new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(totalAmount) }}</strong>
        </div>
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
      <h3>No Data for "{{ title }}"</h3>
      <p>There are no {{ transactionType }} transactions for this selection.</p>
    </div>
  </div>
</template>

<style scoped>
.breakdown-container {
  border: 1px solid var(--border-color);
  border-radius: var(--radius);
  padding: 1.5rem;
  background-color: #fafbfd;
  min-height: 500px;
}
.content-wrapper {
  display: grid;
  grid-template-columns: 2fr 3fr;
  gap: 2rem;
  align-items: center;
}
.chart-area {
  position: relative;
  height: 250px;
  max-width: 250px;
  margin: 0 auto;
}
.total-display {
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
  font-size: 2rem;
  font-weight: 700;
  line-height: 1.2;
}
.summary-area ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-height: 250px;
  overflow-y: auto;
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
  flex-shrink: 0;
}
.category-stats {
  display: flex;
  align-items: baseline;
  gap: 0.75rem;
}
.percentage {
  font-weight: 600;
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
  min-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: var(--text-secondary);
}
.no-data-view .icon { font-size: 3rem; margin-bottom: 1rem; }
.no-data-view h3 { font-size: 1.25rem; color: var(--text-primary); margin-bottom: 0.5rem; }
.no-data-view p { margin: 0.25rem; }

@media (max-width: 1200px) {
  .content-wrapper {
    grid-template-columns: 1fr;
  }
  .chart-area {
    height: 300px;
    max-width: 300px;
  }
}
</style>
