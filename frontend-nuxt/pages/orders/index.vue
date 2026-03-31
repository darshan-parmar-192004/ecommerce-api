<script setup>
definePageMeta({
  middleware: ['auth']
})

useSeoMeta({
  title: 'My Orders - E-Commerce Store'
})

const { data: ordersData, pending, error, refresh } = await useFetch('/api/orders')

const orders = computed(() => ordersData.value?.data || [])

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getStatusColor = (status) => {
  const colors = {
    pending: 'bg-yellow-100 text-yellow-800',
    processing: 'bg-blue-100 text-blue-800',
    shipped: 'bg-primary-100 text-primary-800',
    delivered: 'bg-green-100 text-green-800',
    cancelled: 'bg-red-100 text-red-800'
  }
  return colors[status] || 'bg-gray-100 text-gray-800'
}

const formatCurrency = (amount) => {
  if (amount === undefined || amount === null) return '$0.00'
  return `$${Number(amount).toFixed(2)}`
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <h1 class="text-3xl font-bold mb-8 text-gray-900">My Orders</h1>

      <div v-if="pending" class="space-y-4">
        <div v-for="i in 3" :key="i" class="bg-white rounded-lg shadow-sm p-6 animate-pulse">
          <div class="flex justify-between items-center">
            <div>
              <div class="h-5 bg-gray-200 rounded w-32 mb-2"></div>
              <div class="h-4 bg-gray-200 rounded w-48"></div>
            </div>
            <div class="h-6 bg-gray-200 rounded w-24"></div>
          </div>
        </div>
      </div>

      <div v-else-if="error" class="bg-white rounded-lg shadow-sm p-8 text-center">
        <svg class="mx-auto h-16 w-16 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <h3 class="mt-4 text-lg font-medium text-gray-900">Failed to load orders</h3>
        <p class="mt-2 text-gray-500">Something went wrong. Please try again.</p>
        <button @click="refresh" class="mt-4 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors">
          Try Again
        </button>
      </div>

      <div v-else-if="orders.length === 0" class="text-center py-16 bg-white rounded-lg shadow-sm">
        <svg class="mx-auto h-24 w-24 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <h2 class="mt-4 text-xl font-medium text-gray-900">No orders yet</h2>
        <p class="mt-2 text-gray-500">Start shopping to see your orders here.</p>
        <NuxtLink to="/products" class="inline-block mt-6 px-6 py-3 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors">
          Browse Products
        </NuxtLink>
      </div>

      <div v-else class="space-y-4">
        <NuxtLink 
          v-for="order in orders" 
          :key="order.order_id"
          :to="`/orders/${order.order_id}`"
          class="block bg-white rounded-lg shadow-sm p-6 hover:shadow-md transition-shadow"
        >
          <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div>
              <div class="flex items-center gap-3">
                <h3 class="font-semibold text-gray-900">
                  Order #{{ order.order_id }}
                </h3>
                <span 
                  :class="['px-2 py-1 text-xs font-medium rounded-full capitalize', getStatusColor(order.status)]"
                >
                  {{ order.status }}
                </span>
              </div>
              <p class="text-sm text-gray-500 mt-1">
                {{ formatDate(order.order_date) }}
              </p>
            </div>
            
            <div class="text-left sm:text-right">
              <p class="text-lg font-semibold text-primary-600">
                {{ formatCurrency(order.total_amount) }}
              </p>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>