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
    console.error('Failed to fetch order details', err)
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
    <button @click="goBack" class="mb-6 text-indigo-600 hover:text-indigo-700 flex items-center gap-2">
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
      </svg>
      Back to Orders
    </button>

    <div v-if="loading" class="space-y-4">
      <div class="h-64 bg-gray-100 rounded-xl animate-pulse" />
    </div>

    <div v-else-if="order" class="space-y-6">
      <!-- Order Header -->
      <div class="bg-white p-6 rounded-xl border border-gray-200">
        <div class="flex justify-between items-start mb-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">Order #{{ order.order_id }}</h1>
            <p class="text-sm text-gray-500 mt-1">Placed on {{ new Date(order.order_date).toLocaleDateString() }}</p>
          </div>
          <span
            class="px-3 py-1 text-sm font-medium rounded-full"
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
        
        <div class="border-t border-gray-200 pt-4">
          <h3 class="text-sm font-medium text-gray-700 mb-2">Shipping Address</h3>
          <p class="text-gray-900">{{ order.shipping_address }}</p>
        </div>
      </div>

      <!-- Order Items -->
      <div class="bg-white rounded-xl border border-gray-200 overflow-hidden">
        <div class="p-6">
          <h2 class="text-lg font-semibold text-gray-900 mb-4">Order Items</h2>
          <div v-if="!order.items || order.items.length === 0" class="text-gray-500 text-center py-4">
            <p>No items in this order</p>
          </div>
          <div v-else class="space-y-3">
            <div v-for="item in order.items" :key="item.order_item_id" class="flex justify-between items-center py-3 border-b border-gray-100 last:border-b-0">
              <div>
                <p class="text-sm font-medium text-gray-900">{{ item.product_id }}</p>
                <p class="text-xs text-gray-500">Quantity: {{ item.quantity }}</p>
                <p class="text-xs text-gray-500">Unit Price: ₹{{ item.unit_price?.toFixed(2) }}</p>
              </div>
              <p class="text-sm font-semibold text-gray-900">₹{{ (item.unit_price * item.quantity).toFixed(2) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-6 py-4 border-t border-gray-200">
          <div class="flex justify-between items-center">
            <span class="text-lg font-semibold text-gray-900">Total</span>
            <span class="text-2xl font-bold text-gray-900">₹{{ order.total_amount?.toFixed(2) }}</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-16">
      <p class="text-gray-500 mb-4">Order not found</p>
      <button @click="goBack" class="btn-primary">Back to Orders</button>
    </div>
  </div>
</template>