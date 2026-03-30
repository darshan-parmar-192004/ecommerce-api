<script setup>
definePageMeta({
  middleware: ['auth']
})

const route = useRoute()
const orderId = route.params.id

useSeoMeta({
  title: () => `Order #${orderId} - E-Commerce Store`
})

const { data: orderData, pending, error } = await useFetch(`/api/orders/${orderId}`)

const order = computed(() => orderData.value)

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusColor = (status) => {
  const colors = {
    pending: 'bg-yellow-100 text-yellow-800',
    processing: 'bg-blue-100 text-blue-800',
    shipped: 'bg-indigo-100 text-indigo-800',
    delivered: 'bg-green-100 text-green-800',
    cancelled: 'bg-red-100 text-red-800'
  }
  return colors[status] || 'bg-gray-100 text-gray-800'
}

const formatCurrency = (amount) => {
  if (amount === undefined || amount === null) return '$0.00'
  return `$${Number(amount).toFixed(2)}`
}

const orderItems = computed(() => {
  if (!order.value) return []
  if (Array.isArray(order.value.items)) return order.value.items
  if (Array.isArray(order.value.data)) return order.value.data
  return []
})

const calculateItemTotal = (item) => {
  const qty = item.quantity || 0
  const price = item.unit_price || 0
  return qty * price
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <NuxtLink to="/orders" class="inline-flex items-center text-sm text-gray-500 hover:text-indigo-600 mb-6 transition-colors">
      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Orders
    </NuxtLink>

    <div v-if="pending" class="animate-pulse space-y-6">
      <div class="h-8 bg-gray-200 rounded w-1/3"></div>
      <div class="bg-white rounded-lg shadow-sm p-6">
        <div class="h-4 bg-gray-200 rounded w-1/2 mb-4"></div>
        <div class="h-4 bg-gray-200 rounded w-1/3"></div>
      </div>
    </div>

    <div v-else-if="error" class="bg-white rounded-lg shadow-sm p-8 text-center">
      <svg class="mx-auto h-16 w-16 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <h1 class="mt-4 text-2xl font-bold text-gray-900 mb-2">Order Not Found</h1>
      <p class="text-gray-500 mb-6">The order you're looking for doesn't exist or you don't have permission to view it.</p>
      <NuxtLink to="/orders" class="btn-primary">
        View All Orders
      </NuxtLink>
    </div>

    <div v-else-if="order" class="space-y-6">
      <div class="bg-white rounded-lg shadow-sm p-6">
        <div class="flex items-center justify-between mb-6">
          <h1 class="text-2xl font-bold text-gray-900">Order #{{ order.order_id }}</h1>
          <span 
            :class="['px-3 py-1 text-sm font-medium rounded-full capitalize', getStatusColor(order.status)]"
          >
            {{ order.status }}
          </span>
        </div>

        <div class="grid grid-cols-2 gap-6 text-sm">
          <div>
            <p class="text-gray-500">Order Date</p>
            <p class="font-medium text-gray-900">{{ formatDate(order.order_date) }}</p>
          </div>
          <div>
            <p class="text-gray-500">Total Amount</p>
            <p class="font-semibold text-xl text-indigo-600">{{ formatCurrency(order.total_amount) }}</p>
          </div>
        </div>
      </div>

      <div v-if="order.shipping_address" class="bg-white rounded-lg shadow-sm p-6">
        <h2 class="text-lg font-semibold mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          Shipping Address
        </h2>
        <p class="text-gray-600 whitespace-pre-line">{{ order.shipping_address }}</p>
      </div>

      <div class="bg-white rounded-lg shadow-sm p-6">
        <h2 class="text-lg font-semibold mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          Order Items
        </h2>
        
        <div v-if="orderItems.length > 0" class="space-y-4">
          <div 
            v-for="(item, index) in orderItems" 
            :key="index"
            class="flex items-center gap-4 py-4 border-b border-gray-100 last:border-0"
          >
            <div class="w-16 h-16 bg-gray-100 rounded-lg flex-shrink-0 flex items-center justify-center">
              <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            
            <div class="flex-1">
              <p class="font-medium text-gray-900">
                {{ item.product_name || item.product_id || 'Product' }}
              </p>
              <p class="text-sm text-gray-500">
                {{ item.quantity }} × {{ formatCurrency(item.unit_price) }}
              </p>
            </div>

            <p class="font-semibold text-gray-900">
              {{ formatCurrency(calculateItemTotal(item)) }}
            </p>
          </div>
        </div>
        <div v-else class="text-center py-8 text-gray-500">
          <p>Order items not available</p>
        </div>
      </div>
    </div>
  </div>
</template>
