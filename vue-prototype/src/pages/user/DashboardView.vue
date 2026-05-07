<script setup>
import { onMounted } from 'vue'
import { useUser } from '@/composables/useUser'
import { useRouter } from 'vue-router'
import { useMotion } from '@vueuse/motion'

const { orders, loading, fetchOrders, fetchProfile } = useUser()
const router = useRouter()

onMounted(() => {
  fetchOrders({ limit: 5 })
  fetchProfile()
})
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Dashboard</h1>

    <div useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0 } }" class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Total Orders</p>
        <p class="text-3xl font-bold text-gray-900 dark:text-gray-100 mt-2">{{ orders.length }}</p>
      </div>
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Account Status</p>
        <p class="text-lg font-semibold text-green-600 dark:text-green-400 mt-2">Active</p>
      </div>
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <p class="text-sm text-gray-600 dark:text-gray-400">Member Since</p>
        <p class="text-lg font-semibold text-gray-900 dark:text-gray-100 mt-2">{{ new Date().getFullYear() }}</p>
      </div>
    </div>

    <div class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 p-6">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-100">Recent Orders</h2>
        <button @click="router.push('/user/orders')" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100">
          View All →
        </button>
      </div>

      <div v-if="loading" class="space-y-3">
        <div v-for="n in 3" :key="n" class="animate-pulse h-16 bg-gray-100 dark:bg-brand-700 rounded-lg" />
      </div>

      <div v-else-if="orders.length > 0" class="space-y-3">
        <div v-for="order in orders" :key="order.id" class="flex items-center justify-between p-4 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
          <div>
            <p class="font-medium text-gray-900 dark:text-gray-100">#{{ order.id }}</p>
            <p class="text-sm text-gray-600 dark:text-gray-400">{{ order.date }}</p>
          </div>
          <div class="text-right">
            <p class="font-semibold text-gray-900 dark:text-gray-100">₹{{ order.total?.toFixed(2) }}</p>
            <span class="text-xs px-2 py-1 rounded-full" :class="order.status === 'completed' ? 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400' : 'bg-yellow-100 dark:bg-yellow-900/30 text-yellow-700 dark:text-yellow-400'">
              {{ order.status }}
            </span>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-8 text-gray-500 dark:text-gray-400">
        No orders yet. <RouterLink to="/" class="text-gray-900 dark:text-gray-100 underline">Start shopping</RouterLink>
      </div>
    </div>
  </div>
</template>
