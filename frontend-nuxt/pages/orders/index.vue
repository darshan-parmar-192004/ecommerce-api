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
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-on_surface mb-8">My Orders</h1>

    <div v-if="pending" class="space-y-4">
      <div v-for="i in 3" :key="i" class="bg-surface-container-lowest rounded-lg p-6 animate-pulse">
        <div class="h-6 bg-surface-container rounded w-1/3 mb-4" />
        <div class="h-4 bg-surface-container rounded w-1/2" />
      </div>
    </div>

    <div v-else-if="error" class="text-center py-12">
      <i class="pi pi-exclamation-triangle text-5xl text-error mb-4" />
      <h2 class="text-xl font-semibold text-on_surface mb-2">Failed to load orders</h2>
      <p class="text-on_surface_variant mb-4">{{ error.message }}</p>
      <Button @click="refresh" label="Try Again" icon="pi pi-refresh" />
    </div>

    <div v-else-if="orders.length === 0" class="text-center py-12">
      <i class="pi pi-box text-5xl text-outline mb-4" />
      <h2 class="text-xl font-semibold text-on_surface mb-2">No orders yet</h2>
      <p class="text-on_surface_variant mb-6">Start shopping to see your orders here.</p>
      <NuxtLink to="/products">
        <Button label="Browse Products" icon="pi pi-shopping-cart" />
      </NuxtLink>
    </div>

    <div v-else class="space-y-4">
      <NuxtLink
        v-for="order in orders"
        :key="order.order_id || order.id"
        :to="`/orders/${order.order_id || order.id}`"
        class="block bg-surface-container-lowest rounded-lg p-6 hover:shadow-lg transition-shadow"
      >
        <div class="flex items-center justify-between mb-3">
          <span class="font-semibold text-on_surface">
            Order #{{ order.order_id || order.id }}
          </span>
          <Tag
            :value="order.status || 'Pending'"
            :severity="order.status === 'delivered' ? 'success' : order.status === 'cancelled' ? 'danger' : 'warn'"
          />
        </div>
        <div class="text-sm text-on_surface_variant space-y-1">
          <p>{{ formatDate(order.created_at || order.order_date) }}</p>
          <p class="font-medium text-primary">
            ₹ {{ Number(order.total_amount || order.total || 0).toFixed(2) }}
          </p>
          <p v-if="order.items_count">{{ order.items_count }} item(s)</p>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
