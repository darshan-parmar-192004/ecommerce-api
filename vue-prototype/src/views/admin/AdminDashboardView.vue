<script setup>
import { ref, onMounted } from 'vue'
import { useMotion } from '@vueuse/motion'
import productService from '@/services/productService'

const stats = ref({
  totalProducts: 0,
  totalOrders: 0,
  revenue: 0,
  activeUsers: 0
})
const loading = ref(true)

onMounted(async () => {
  try {
    const { data } = await productService.getProducts({ limit: 1 })
    stats.value.totalProducts = data.total || 0
  } catch (err) {
    console.error('Failed to fetch stats', err)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="admin-dashboard-page">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Admin Dashboard</h1>

    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div v-for="n in 4" :key="n" class="h-32 bg-gray-100 rounded-xl animate-pulse" />
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div
        useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0, delay: 100 } }"
        class="bg-white p-6 rounded-xl border border-gray-200"
      >
        <p class="text-sm text-gray-600">Total Products</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.totalProducts }}</p>
      </div>
      <div
        useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0, delay: 200 } }"
        class="bg-white p-6 rounded-xl border border-gray-200"
      >
        <p class="text-sm text-gray-600">Total Orders</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.totalOrders }}</p>
      </div>
      <div
        useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0, delay: 300 } }"
        class="bg-white p-6 rounded-xl border border-gray-200"
      >
        <p class="text-sm text-gray-600">Revenue</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">${{ stats.revenue.toFixed(2) }}</p>
      </div>
      <div
        useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0, delay: 400 } }"
        class="bg-white p-6 rounded-xl border border-gray-200"
      >
        <p class="text-sm text-gray-600">Active Users</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.activeUsers }}</p>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h2>
        <div class="space-y-3">
          <RouterLink to="/admin/products" class="block p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            Manage Products
          </RouterLink>
          <RouterLink to="/admin/inventory" class="block p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            Inventory Management
          </RouterLink>
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
