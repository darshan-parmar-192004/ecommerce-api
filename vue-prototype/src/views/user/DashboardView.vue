<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { useMotion } from '@vueuse/motion'

const userStore = useUserStore()
const router = useRouter()

onMounted(() => {
  userStore.fetchOrders({ limit: 5 })
  userStore.fetchProfile()
})
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Dashboard</h1>

    <div useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0 } }" class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Total Orders</p>
        <p class="text-3xl font-bold text-gray-900 mt-2">{{ userStore.orders.length }}</p>
      </div>
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Account Status</p>
        <p class="text-lg font-semibold text-green-600 mt-2">Active</p>
      </div>
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <p class="text-sm text-gray-600">Member Since</p>
        <p class="text-lg font-semibold text-gray-900 mt-2">{{ new Date().getFullYear() }}</p>
      </div>
    </div>

    <div class="bg-white rounded-xl border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-xl font-semibold text-gray-900">Recent Orders</h2>
        <button @click="router.push('/user/orders')" class="text-sm text-gray-600 hover:text-gray-900">
          View All →
        </button>
      </div>

      <div v-if="userStore.loading" class="space-y-3">
        <div v-for="n in 3" :key="n" class="animate-pulse h-16 bg-gray-100 rounded-lg" />
      </div>

      <div v-else-if="userStore.orders.length > 0" class="space-y-3">
        <div v-for="order in userStore.orders" :key="order.id" class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
          <div>
            <p class="font-medium text-gray-900">#{{ order.id }}</p>
            <p class="text-sm text-gray-600">{{ order.date }}</p>
          </div>
          <div class="text-right">
            <p class="font-semibold text-gray-900">${{ order.total?.toFixed(2) }}</p>
            <span class="text-xs px-2 py-1 rounded-full" :class="order.status === 'completed' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'">
              {{ order.status }}
            </span>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-8 text-gray-500">
        No orders yet. <RouterLink to="/" class="text-gray-900 underline">Start shopping</RouterLink>
      </div>
    </div>
  </div>
</template>
