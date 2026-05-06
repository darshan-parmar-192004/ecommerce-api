<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useToastStore } from '@/stores/toast'

const router = useRouter()
const userStore = useUserStore()
const toastStore = useToastStore()

onMounted(() => {
  userStore.fetchOrders().catch(() => {
    toastStore.error('Failed to load orders')
  })
})

const viewOrder = (orderId) => {
  router.push({ name: 'OrderDetail', params: { id: orderId } })
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Order History</h1>

    <div v-if="userStore.loading" class="space-y-4">
      <div v-for="n in 5" :key="n" class="animate-pulse h-24 bg-gray-100 dark:bg-brand-700 rounded-xl" />
    </div>

    <div v-else-if="userStore.orders.length > 0" class="space-y-4">
      <div
        v-for="order in userStore.orders"
        :key="order.order_id"
        class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700 flex items-center justify-between hover:shadow-md transition-shadow cursor-pointer"
        @click="viewOrder(order.order_id)"
      >
        <div>
          <p class="font-medium text-gray-900 dark:text-gray-100">Order #{{ order.order_id }}</p>
          <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">{{ order.items?.length || 0 }} items</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">{{ order.order_date }}</p>
        </div>
        <div class="text-right">
          <p class="text-xl font-bold text-gray-900 dark:text-gray-100">₹{{ order.total_amount?.toFixed(2) }}</p>
          <span
            class="inline-block mt-1 px-3 py-1 text-xs font-medium rounded-full"
            :class="{
              'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400': order.status === 'completed',
              'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400': order.status === 'pending',
              'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400': order.status === 'shipped',
              'bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400': order.status === 'cancelled'
            }"
          >
            {{ order.status }}
          </span>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-16">
      <p class="text-gray-500 dark:text-gray-400 mb-4">No orders found</p>
      <RouterLink to="/" class="btn-primary">
        Start Shopping
      </RouterLink>
    </div>
  </div>
</template>