<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import { useTransactionStore } from '@/stores/transactions';
import CategoryBreakdownChart from '@/components/CategoryBreakdownChart.vue';
import DailySpendingChart from '@/components/DailySpendingChart.vue';
import type { ChartData } from 'chart.js';

const store = useTransactionStore();
type ChartType = 'debit' | 'credit' | 'refund';

// State for Pie Charts
const mainChartType = ref<ChartType>('debit');
const subChartType = ref<ChartType>('debit');

// State for Line Chart Filter
const dailyFilterDates = ref({
  startDate: store.dailySpendingFilter.startDate,
  endDate: store.dailySpendingFilter.endDate,
});

watch(dailyFilterDates, (newDates) => {
  store.updateDailySpendingFilter(newDates);
}, { deep: true });

function clearDailyFilter() {
  store.clearDailySpendingFilter();
  dailyFilterDates.value.startDate = '';
  dailyFilterDates.value.endDate = '';
}

// Use a read-only computed for the value
const dailyChartType = computed(() => store.dailySpendingChartType);

// Create an explicit method to call the store action
function setDailySpendingChartType(type: ChartType) {
  store.setDailySpendingChartType(type);
}

const pieColors = [
  '#3b82f6', '#f97316', '#ec4899', '#14b8a6', '#8b5cf6',
  '#eab308', '#ef4444', '#0ea5e9', '#10b981', '#64748b'
];

const createChartDataObject = (chartData: { labels: string[], data: number[] }): ChartData<'pie'> => ({
  labels: chartData.labels,
  datasets: [{
    backgroundColor: pieColors,
    data: chartData.data,
    borderWidth: 2,
    borderColor: '#fafbfd',
  }],
});

// Computed properties for Pie Charts
const mainChartData = computed(() => createChartDataObject(
  mainChartType.value === 'credit' ? store.creditChartData :
    mainChartType.value === 'refund' ? store.refundChartData :
      store.debitChartData
));
const subCategoryChartData = computed(() => createChartDataObject(
  subChartType.value === 'credit' ? store.creditSubCategoryChartData :
    subChartType.value === 'refund' ? store.refundSubCategoryChartData :
      store.debitSubCategoryChartData
));
const selectedCategory = computed({
  get: () => store.selectedCategoryForDashboard,
  set: (value) => store.setSelectedCategoryForDashboard(value)
});
watch(() => store.availableCategoriesInFiltered, (newCategories) => {
  if (newCategories.length > 0 && !store.selectedCategoryForDashboard) {
    store.setSelectedCategoryForDashboard(newCategories[0]);
  } else if (newCategories.length === 0) {
    store.setSelectedCategoryForDashboard(null);
  }
}, { immediate: true });

// Computed property for Line Chart
const dailySpendingData = computed(() => store.dailySpendingChartData);

</script>

<template>
  <div class="dashboard-card">
    <div class="card-header">
      <div class="header-text">
        <h2>Breakdown Dashboards</h2>
        <p>An overview of your transactions by category and sub-category.</p>
      </div>
    </div>

    <div class="card-body">
      <div class="charts-grid">
        <div class="chart-wrapper full-width">
          <div class="chart-controls space-between">
            <div class="date-filters">
              <div class="form-group">
                <label for="line-start-date">Start Date</label>
                <input id="line-start-date" type="date" v-model="dailyFilterDates.startDate" />
              </div>
              <div class="form-group">
                <label for="line-end-date">End Date</label>
                <input id="line-end-date" type="date" v-model="dailyFilterDates.endDate" />
              </div>
              <button @click="clearDailyFilter" class="clear-button">Clear</button>
            </div>
            <div class="segmented-control">
              <button @click="setDailySpendingChartType('debit')" :class="{ active: dailyChartType === 'debit' }">Debits</button>
              <button @click="setDailySpendingChartType('credit')" :class="{ active: dailyChartType === 'credit' }">Credits</button>
              <button @click="setDailySpendingChartType('refund')" :class="{ active: dailyChartType === 'refund' }">Refunds</button>
            </div>
          </div>
          <DailySpendingChart :chart-data="dailySpendingData" :transaction-type="dailyChartType" />
        </div>

        <div class="chart-wrapper">
          <div class="chart-controls">
            <div class="segmented-control">
              <button @click="mainChartType = 'debit'" :class="{ active: mainChartType === 'debit' }">Debits</button>
              <button @click="mainChartType = 'credit'" :class="{ active: mainChartType === 'credit' }">Credits</button>
              <button @click="mainChartType = 'refund'" :class="{ active: mainChartType === 'refund' }">Refunds</button>
            </div>
          </div>
          <CategoryBreakdownChart
            title="Overall Category Breakdown"
            :chart-data="mainChartData"
            :transaction-type="mainChartType"
            :key="`main-${mainChartType}`"
          />
        </div>

        <div class="chart-wrapper">
          <div class="chart-controls sub-category-header">
            <div class="sub-cat-select-group">
              <label for="category-select">Sub-Category Breakdown For:</label>
              <select id="category-select" v-model="selectedCategory">
                <option :value="null" disabled>-- Select a Category --</option>
                <option v-for="cat in store.availableCategoriesInFiltered" :key="cat" :value="cat">
                  {{ cat }}
                </option>
              </select>
            </div>
            <div class="segmented-control">
              <button @click="subChartType = 'debit'" :class="{ active: subChartType === 'debit' }">Debits</button>
              <button @click="subChartType = 'credit'" :class="{ active: subChartType === 'credit' }">Credits</button>
              <button @click="subChartType = 'refund'" :class="{ active: subChartType === 'refund' }">Refunds</button>
            </div>
          </div>
          <CategoryBreakdownChart
            v-if="store.selectedCategoryForDashboard"
            :title="`'${store.selectedCategoryForDashboard}' Breakdown`"
            :chart-data="subCategoryChartData"
            :transaction-type="subChartType"
            :key="`sub-${store.selectedCategoryForDashboard}-${subChartType}`"
          />
          <div v-else class="select-prompt">
            <p>Please select a category above to see its breakdown.</p>
          </div>
        </div>
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
  padding: 1.5rem;
  border-bottom: 1px solid var(--border-color);
}
.header-text h2 { margin: 0; font-size: 1.5rem; font-weight: 600; }
.header-text p { margin: 0.25rem 0 0; color: var(--text-secondary); }

.segmented-control {
  display: flex;
  background-color: var(--bg-main);
  border-radius: 6px;
  padding: 4px;
  width: fit-content;
}
.segmented-control button {
  background: transparent; border: none; border-radius: 4px;
  padding: 0.5rem 1rem; font-family: inherit; font-size: 0.875rem;
  font-weight: 600; cursor: pointer; color: var(--text-secondary);
  transition: background-color 0.2s, color 0.2s, box-shadow 0.2s;
}
.segmented-control button.active {
  background-color: var(--bg-card);
  color: var(--text-primary);
  box-shadow: var(--shadow-sm);
}

.card-body {
  padding: 2rem;
  background-color: var(--bg-main);
}
.charts-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}
.chart-wrapper.full-width {
  grid-column: 1 / -1;
}
.chart-controls {
  margin-bottom: 1rem;
  display: flex;
  justify-content: flex-end;
}
/* New helper class for space-between layout */
.chart-controls.space-between {
  justify-content: space-between;
  align-items: flex-end;
  flex-wrap: wrap;
  gap: 1rem;
}
.sub-category-header {
  justify-content: space-between;
  align-items: flex-end;
  flex-wrap: wrap;
  gap: 1rem;
}
.sub-cat-select-group {
  display: flex; flex-direction: column; gap: 0.5rem; flex-grow: 1;
}
.sub-cat-select-group label { font-weight: 600; font-size: 1rem; }

#category-select {
  width: 100%; max-width: 400px; padding: 0.75rem; font-size: 1rem;
  font-family: inherit; border: 1px solid var(--border-color);
  border-radius: 6px; background-color: var(--bg-card);
}
.select-prompt {
  border: 1px solid var(--border-color); border-radius: var(--radius);
  background-color: #fafbfd; min-height: 500px; display: flex;
  align-items: center; justify-content: center;
  color: var(--text-secondary); font-weight: 500;
}
.date-filters {
  display: flex;
  align-items: flex-end;
  gap: 1rem;
  justify-content: flex-start;
}
.form-group label {
  display: block; font-size: 0.875rem; font-weight: 500;
  color: var(--text-secondary); margin-bottom: 0.5rem;
}
.form-group input[type="date"] {
  padding: 0.5rem; font-family: inherit; border-radius: 6px;
  border: 1px solid var(--border-color); background-color: var(--bg-card);
}
.clear-button {
  padding: 0.5rem 1rem; font-family: inherit; font-weight: 500;
  border-radius: 6px; border: 1px solid var(--border-color);
  background-color: var(--bg-card); cursor: pointer;
  transition: background-color 0.2s;
  height: fit-content;
}
.clear-button:hover { background-color: #f1f5f9; }
@media (max-width: 900px) {
  .charts-grid {
    grid-template-columns: 1fr;
  }
}
</style>
