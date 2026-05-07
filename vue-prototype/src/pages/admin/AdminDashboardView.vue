<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/lib/productService'
import { useErrorHandler } from '@/composables/useErrorHandler'

const { showError } = useErrorHandler()

const stats = ref({
  totalProducts: 0,
  totalOrders: 0,
  revenue: 0,
  activeUsers: 0
})
const topSellers = ref([])
const lifetimeValues = ref([])
const loading = ref(true)

const fetchStats = async () => {
  try {
    const productsRes = await productService.getProducts({ limit: 1 })
    const productsData = productsRes.data || productsRes
    stats.value.totalProducts = productsData.pagination?.total_items || 0
    // TODO: Replace with actual API calls when backend endpoints are ready
    // stats.value.totalOrders = await orderService.getTotalOrders()
    // stats.value.revenue = await orderService.getTotalRevenue()
    // stats.value.activeUsers = await customerService.getActiveUsers()
    stats.value.totalOrders = 0
    stats.value.revenue = 0
    stats.value.activeUsers = 0
  } catch (err) {
    showError(err, 'Failed to fetch stats')
  }
}

const fetchTopSellers = async () => {
  try {
    const response = await productService.getTopSellers()
    const responseData = response.data || response
    topSellers.value = responseData.data || responseData || []
  } catch (err) {
    showError(err, 'Failed to fetch top sellers')
  }
}

const fetchLifetimeValues = async () => {
  try {
    const response = await productService.getCustomerLifetimeValue()
    const responseData = response.data || response
    lifetimeValues.value = responseData.data || responseData || []
  } catch (err) {
    showError(err, 'Failed to fetch lifetime values')
  }
}

onMounted(async () => {
  await Promise.all([
    fetchStats(),
    fetchTopSellers(),
    fetchLifetimeValues()
  ])
  loading.value = false
})
</script>

<template>
  <div class="admin-dashboard-page">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Admin Dashboard</h1>

    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div v-for="n in 4" :key="n" class="h-32 bg-gray-100 dark:bg-brand-700 rounded-xl animate-pulse" />
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Total Products</p>
        <p class="text-3xl font-bold text-gray-900 dark:text-gray-100 mt-2">{{ stats.totalProducts }}</p>
      </div>
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Total Orders</p>
        <p class="text-3xl font-bold text-gray-900 dark:text-gray-100 mt-2">{{ stats.totalOrders }}</p>
      </div>
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Revenue</p>
        <p class="text-3xl font-bold text-gray-900 dark:text-gray-100 mt-2">₹{{ stats.revenue.toFixed(2) }}</p>
      </div>
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Active Users</p>
        <p class="text-3xl font-bold text-gray-900 dark:text-gray-100 mt-2">{{ stats.activeUsers }}</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Top Sellers -->
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">Top Selling Products</h2>
        <div v-if="topSellers.length === 0" class="text-gray-500 dark:text-gray-400 text-center py-8">
          No sales data available
        </div>
        <div v-else class="space-y-3">
          <div v-for="(item, index) in topSellers" :key="item.product || 'top-seller-' + index" class="flex items-center justify-between p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
            <div class="flex items-center gap-3">
              <span class="w-6 h-6 bg-indigo-100 dark:bg-indigo-900/30 text-indigo-700 dark:text-indigo-400 rounded-full flex items-center justify-center text-sm font-bold">
                {{ index + 1 }}
              </span>
              <span class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ item.product || 'Product' }}</span>
            </div>
            <span class="text-sm font-bold text-gray-900 dark:text-gray-100">{{ item.units_sold || 0 }} units</span>
          </div>
        </div>
      </div>

      <!-- Customer Lifetime Value -->
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">Customer Lifetime Value</h2>
        <div v-if="lifetimeValues.length === 0" class="text-gray-500 dark:text-gray-400 text-center py-8">
          No customer data available
        </div>
        <div v-else class="space-y-3">
          <div v-for="(item, index) in lifetimeValues" :key="item.customerId" class="flex items-center justify-between p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
            <div class="flex items-center gap-3">
              <span class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ item.customerId }}</span>
            </div>
            <div class="text-right">
              <p class="text-sm font-bold text-gray-900 dark:text-gray-100">₹{{ item.lifetimeValue?.toFixed(2) }}</p>
              <p class="text-xs text-gray-500 dark:text-gray-400">{{ item.totalOrders || 0 }} orders</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
      <!-- Quick Actions -->
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">Quick Actions</h2>
        <div class="space-y-3">
          <RouterLink to="/admin/products" class="block p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg hover:bg-gray-100 dark:hover:bg-brand-700 transition-colors">
            Manage Products
          </RouterLink>
          <RouterLink to="/admin/categories" class="block p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg hover:bg-gray-100 dark:hover:bg-brand-700 transition-colors">
            Manage Categories
          </RouterLink>
          <RouterLink to="/admin/inventory" class="block p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg hover:bg-gray-100 dark:hover:bg-brand-700 transition-colors">
            Inventory Management
          </RouterLink>
        </div>
      </div>

      <!-- System Status -->
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">System Status</h2>
        <div class="space-y-3">
          <div class="flex justify-between items-center p-3 bg-green-50 dark:bg-green-900/20 rounded-lg">
            <span class="text-sm text-green-700 dark:text-green-400">Database</span>
            <span class="text-sm font-medium text-green-700 dark:text-green-400">Connected</span>
          </div>
          <div class="flex justify-between items-center p-3 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
            <span class="text-sm text-blue-700 dark:text-blue-400">Cache</span>
            <span class="text-sm font-medium text-blue-700 dark:text-blue-400">Active</span>
          </div>
          <div class="flex justify-between items-center p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
            <span class="text-sm text-gray-700 dark:text-gray-300">API Version</span>
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">v1.0</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-dashboard-page {
  max-width: 1400px;
  margin: 0 auto;
  padding: 1rem;
}
</style>
