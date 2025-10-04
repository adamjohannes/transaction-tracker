<script setup lang="ts">
import { computed } from 'vue';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  LinearScale,
  PointElement,
  CategoryScale,
  type ChartData,
} from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, LineElement, LinearScale, PointElement, CategoryScale);

const props = defineProps<{
  chartData: { labels: string[], data: number[] };
  transactionType: 'debit' | 'credit' | 'refund';
}>();

const chartColors = computed(() => {
  switch (props.transactionType) {
    case 'credit':
      return { bg: 'rgba(16, 185, 129, 0.2)', border: '#10b981' };
    case 'refund':
      return { bg: 'rgba(59, 130, 246, 0.2)', border: '#3b82f6' };
    default: // debit
      return { bg: 'rgba(239, 68, 68, 0.2)', border: '#ef4444' };
  }
});

const capitalizedType = computed(() =>
  props.transactionType.charAt(0).toUpperCase() + props.transactionType.slice(1)
);

const chartDataObject = computed<ChartData<'line'>>(() => ({
  labels: props.chartData.labels,
  datasets: [
    {
      label: capitalizedType.value,
      backgroundColor: chartColors.value.bg,
      borderColor: chartColors.value.border,
      pointBackgroundColor: chartColors.value.border,
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: chartColors.value.border,
      data: props.chartData.data,
      fill: true,
      tension: 0.3,
    },
  ],
}));

const chartOptions = computed(() => ({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
    title: {
      display: true,
      text: `${capitalizedType.value} Totals per Day of the Week`,
      font: { size: 16, weight: '600' },
      padding: { bottom: 16 }
    }
  },
  scales: {
    y: {
      beginAtZero: true,
    },
  },
}));
</script>

<template>
  <div class="chart-container">
    <Line :data="chartDataObject" :options="chartOptions" />
  </div>
</template>

<style scoped>
.chart-container {
  position: relative;
  height: 350px;
  padding: 1rem;
  border: 1px solid var(--border-color);
  border-radius: var(--radius);
  background-color: #fafbfd;
}
</style>
