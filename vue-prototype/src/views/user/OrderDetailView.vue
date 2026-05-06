<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useToastStore } from '@/stores/toast'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const toastStore = useToastStore()

const order = ref(null)
const items = ref([])
const loading = ref(true)

onMounted(async () => {
  const orderId = route.params.id
  try {
    await userStore.fetchOrderDetail(orderId)
    order.value = userStore.currentOrder
  } catch (err) {
    toastStore.error('Failed to load order')
  } finally {
    loading.value = false
  }
})

const goBack = () => {
  router.push({ name: 'OrderHistory' })
}
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <button
      @click="goBack"
      class="inline-flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 mb-6 transition-colors"
    >
      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Orders
    </button>

    <div v-if="loading" class="space-y-4">
      <div class="h-64 bg-gray-100 dark:bg-brand-700 rounded-xl animate-pulse" />
    </div>

    <div v-else-if="order" class="space-y-6">
      <!-- Order Header -->
      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700">
        <div class="flex justify-between items-start mb-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-gray-100">Order #{{ order.order_id }}</h1>
            <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">Placed on {{ new Date(order.order_date).toLocaleDateString() }}</p>
          </div>
          <span
            class="px-3 py-1 text-sm font-medium rounded-full"
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

        <div class="border-t border-gray-200 dark:border-brand-700 pt-4">
          <h3 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Shipping Address</h3>
          <p class="text-gray-900 dark:text-gray-100">{{ order.shipping_address }}</p>
        </div>
      </div>

      <!-- Order Items -->
      <div class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 overflow-hidden">
        <div class="p-6">
          <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">Order Items</h2>
          <div v-if="!order.items || order.items.length === 0" class="text-gray-500 dark:text-gray-400 text-center py-4">
            <p>No items in this order</p>
          </div>
          <div v-else class="space-y-3">
            <div v-for="item in order.items" :key="item.order_item_id" class="flex justify-between items-center py-3 border-b border-gray-100 dark:border-brand-700 last:border-b-0">
              <div>
                <p class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ item.product_id }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Quantity: {{ item.quantity }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Unit Price: ₹{{ item.unit_price?.toFixed(2) }}</p>
              </div>
              <p class="text-sm font-semibold text-gray-900 dark:text-gray-100">₹{{ (item.unit_price * item.quantity).toFixed(2) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 dark:bg-brand-700/50 px-6 py-4 border-t border-gray-200 dark:border-brand-700">
          <div class="flex justify-between items-center">
            <span class="text-lg font-semibold text-gray-900 dark:text-gray-100">Total</span>
            <span class="text-2xl font-bold text-gray-900 dark:text-gray-100">₹{{ order.total_amount?.toFixed(2) }}</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-16">
      <p class="text-gray-500 dark:text-gray-400 mb-4">Order not found</p>
      <button @click="goBack" class="btn-primary">Back to Orders</button>
    </div>
  </div>
</template>