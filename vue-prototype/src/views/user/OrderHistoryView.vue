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
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Order History</h1>

    <div v-if="userStore.loading" class="space-y-4">
      <div v-for="n in 5" :key="n" class="animate-pulse h-24 bg-gray-100 rounded-xl" />
    </div>

    <div v-else-if="userStore.orders.length > 0" class="space-y-4">
      <div
        v-for="order in userStore.orders"
        :key="order.order_id"
        class="bg-white p-6 rounded-xl border border-gray-200 flex items-center justify-between hover:shadow-md transition-colors cursor-pointer"
        @click="viewOrder(order.order_id)"
      >
        <div>
          <p class="font-semibold text-gray-900">Order #{{ order.order_id }}</p>
          <p class="text-sm text-gray-600 mt-1">{{ order.items?.length || 0 }} items</p>
          <p class="text-sm text-gray-500">{{ order.order_date }}</p>
        </div>
        <div class="text-right">
          <p class="text-xl font-bold text-gray-900">₹{{ order.total_amount?.toFixed(2) }}</p>
          <span
            class="inline-block mt-1 px-3 py-1 text-xs font-medium rounded-full"
            :class="{
              'bg-green-100 text-green-700': order.status === 'completed',
              'bg-yellow-100 text-yellow-700': order.status === 'pending',
              'bg-blue-100 text-blue-700': order.status === 'shipped',
              'bg-red-100 text-red-700': order.status === 'cancelled'
            }"
          >
            {{ order.status }}
          </span>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-16">
      <p class="text-gray-500 mb-4">No orders found</p>
      <RouterLink to="/" class="btn-primary">Start Shopping</RouterLink>
    </div>
  </div>
</template>