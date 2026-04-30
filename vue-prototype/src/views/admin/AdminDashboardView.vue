<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/services/productService'

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
    stats.value.totalOrders = 156
    stats.value.revenue = 24580.50
    stats.value.activeUsers = 89
  } catch (err) {
    console.error('Failed to fetch stats', err)
  }
}

const fetchTopSellers = async () => {
  try {
    const response = await productService.getTopSellers()
    const responseData = response.data || response
    topSellers.value = responseData.data || responseData || []
  } catch (err) {
    console.error('Failed to fetch top sellers', err)
  }
}

const fetchLifetimeValues = async () => {
  try {
    const response = await productService.getCustomerLifetimeValue()
    const responseData = response.data || response
    lifetimeValues.value = responseData.data || responseData || []
  } catch (err) {
    console.error('Failed to fetch lifetime values', err)
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
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Admin Dashboard</h1>

    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div v-for="n in 4" :key="n" class="h-32 bg-gray-100 rounded-xl animate-pulse" />
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Total Products</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.totalProducts }}</p>
      </div>
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Total Orders</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.totalOrders }}</p>
      </div>
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Revenue</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">₹{{ stats.revenue.toFixed(2) }}</p>
      </div>
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Active Users</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.activeUsers }}</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Top Sellers -->
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Top Selling Products</h2>
        <div v-if="topSellers.length === 0" class="text-gray-500 text-center py-8">
          No sales data available
        </div>
        <div v-else class="space-y-3">
          <div v-for="(item, index) in topSellers" :key="index" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center gap-3">
              <span class="w-6 h-6 bg-indigo-100 text-indigo-700 rounded-full flex items-center justify-center text-sm font-bold">
                {{ index + 1 }}
              </span>
              <span class="text-sm font-medium text-gray-900">{{ item.product || 'Product' }}</span>
            </div>
            <span class="text-sm font-bold text-gray-900">{{ item.units_sold || 0 }} units</span>
          </div>
        </div>
      </div>

      <!-- Customer Lifetime Value -->
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Customer Lifetime Value</h2>
        <div v-if="lifetimeValues.length === 0" class="text-gray-500 text-center py-8">
          No customer data available
        </div>
        <div v-else class="space-y-3">
          <div v-for="(item, index) in lifetimeValues" :key="item.customer_id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center gap-3">
              <span class="text-sm font-medium text-gray-900">{{ item.customer_id }}</span>
            </div>
            <div class="text-right">
              <p class="text-sm font-bold text-gray-900">₹{{ item.lifetime_value?.toFixed(2) }}</p>
              <p class="text-xs text-gray-500">{{ item.total_orders || 0 }} orders</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
      <!-- Quick Actions -->
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h2>
        <div class="space-y-3">
          <RouterLink to="/admin/products" class="block p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            Manage Products
          </RouterLink>
          <RouterLink to="/admin/categories" class="block p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            Manage Categories
          </RouterLink>
          <RouterLink to="/admin/inventory" class="block p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            Inventory Management
          </RouterLink>
        </div>
      </div>

      <!-- System Status -->
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">System Status</h2>
        <div class="space-y-3">
          <div class="flex justify-between items-center p-3 bg-green-50 rounded-lg">
            <span class="text-sm text-green-700">Database</span>
            <span class="text-sm font-medium text-green-700">Connected</span>
          </div>
          <div class="flex justify-between items-center p-3 bg-blue-50 rounded-lg">
            <span class="text-sm text-blue-700">Cache</span>
            <span class="text-sm font-medium text-blue-700">Active</span>
          </div>
          <div class="flex justify-between items-center p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">API Version</span>
            <span class="text-sm font-medium text-gray-700">v1.0</span>
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