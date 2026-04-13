<script setup>
definePageMeta({
  middleware: ['auth']
})

useSeoMeta({
  title: 'My Orders - E-Commerce Store'
})

const { orders: ordersApi } = useApi()

const { data: ordersData, pending, error, refresh } = await useAsyncData(
  'orders',
  () => ordersApi.list()
)

const orders = computed(() => ordersData.value?.data || [])

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script>

<template>
  <div class="min-h-screen bg-surface">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <h1 class="text-3xl font-bold text-on_surface font-display mb-8">My Orders</h1>

      <div v-if="pending" class="space-y-4">
        <div v-for="i in 3" :key="i" class="bg-surface-container-lowest rounded-xl shadow-ambient p-6 animate-pulse">
          <div class="flex justify-between items-center">
            <div>
              <div class="h-5 bg-surface-container rounded w-32 mb-2"></div>
              <div class="h-4 bg-surface-container rounded w-48"></div>
            </div>
            <div class="h-6 bg-surface-container rounded w-24"></div>
          </div>
        </div>
      </div>

      <div v-else-if="error" class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center">
        <svg class="mx-auto h-16 w-16 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <h3 class="mt-4 text-lg font-medium text-on_surface">Failed to load orders</h3>
        <p class="mt-2 text-on_surface_variant">Something went wrong. Please try again.</p>
        <button @click="refresh" class="mt-4 px-4 py-2 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
          Try Again
        </button>
      </div>

      <div v-else-if="orders.length === 0" class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient">
        <svg class="mx-auto h-24 w-24 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <h2 class="mt-4 text-xl font-medium text-on_surface">No orders yet</h2>
        <p class="mt-2 text-on_surface_variant">Start shopping to see your orders here.</p>
        <NuxtLink to="/products" class="inline-block mt-6 px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
          Browse Products
        </NuxtLink>
      </div>

      <div v-else class="space-y-4">
        <NuxtLink 
          v-for="order in orders" 
          :key="order.order_id"
          :to="`/orders/${order.order_id}`"
          class="block bg-surface-container-lowest rounded-xl shadow-ambient p-6 hover:shadow-lg transition-shadow"
        >
          <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div>
              <div class="flex items-center gap-3">
                <h3 class="font-semibold text-on_surface">
                  Order #{{ order.order_id }}
                </h3>
                <span 
                  :class="[
                    'px-3 py-1 text-xs font-medium rounded-full capitalize',
                    order.status === 'delivered' ? 'bg-green-100 text-green-700' :
                    order.status === 'shipped' ? 'bg-primary-fixed text-primary' :
                    order.status === 'cancelled' ? 'bg-error-container text-error' :
                    order.status === 'processing' ? 'bg-blue-100 text-blue-700' :
                    'bg-yellow-100 text-yellow-700'
                  ]"
                >
                  {{ order.status }}
                </span>
              </div>
              <p class="text-sm text-on_surface_variant mt-1">
                {{ formatDate(order.order_date) }}
              </p>
            </div>
            
            <div class="text-left sm:text-right">
              <p class="text-lg font-semibold text-primary">
                ₹ {{ Number(order.total_amount).toFixed(2) }}
              </p>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>