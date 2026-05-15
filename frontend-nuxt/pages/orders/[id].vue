<script setup>
definePageMeta({
  middleware: ['auth']
})

const route = useRoute()
const { orders: ordersApi } = useApi()

const { data: order, pending, error, refresh } = await useAsyncData(
  `order-${route.params.id}`,
  () => ordersApi.get(route.params.id)
)

useSeoMeta({
  title: () => order.value ? `Order #${order.value.order_id || order.value.id} - E-Commerce Store` : 'Order Details'
})

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <NuxtLink to="/orders" class="inline-flex items-center text-sm text-on_surface_variant hover:text-primary mb-6">
      <i class="pi pi-arrow-left mr-1" />
      Back to Orders
    </NuxtLink>

    <div v-if="pending" class="animate-pulse space-y-4">
      <div class="h-8 bg-surface-container rounded w-1/3" />
      <div class="h-4 bg-surface-container rounded w-1/2" />
    </div>

    <div v-else-if="error" class="text-center py-12">
      <i class="pi pi-exclamation-triangle text-5xl text-error mb-4" />
      <h2 class="text-xl font-semibold text-on_surface mb-2">Failed to load order</h2>
      <Button @click="refresh" label="Try Again" icon="pi pi-refresh" />
    </div>

    <div v-else-if="order" class="space-y-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-on_surface font-display">Order #{{ order.order_id || order.id }}</h1>
          <p class="text-sm text-on_surface_variant mt-1">Placed on {{ formatDate(order.created_at || order.order_date) }}</p>
        </div>
        <Tag :value="order.status || 'Pending'" :severity="order.status === 'delivered' ? 'success' : order.status === 'cancelled' ? 'danger' : 'warn'" />
      </div>

      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <h2 class="text-lg font-semibold text-on_surface mb-4">Shipping Address</h2>
        <p class="text-on_surface_variant">{{ order.shipping_address }}</p>
      </div>

      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <h2 class="text-lg font-semibold text-on_surface mb-4">Order Items</h2>
        <div class="space-y-4">
          <div v-for="item in (order.items || order.order_items || [])" :key="item.order_item_id || item.id" class="flex items-center gap-4 py-3 border-b border-outline-variant/20 last:border-0">
            <div class="w-16 h-16 bg-surface-container rounded-lg flex items-center justify-center">
              <i class="pi pi-box text-2xl text-outline" />
            </div>
            <div class="flex-1">
              <p class="font-medium text-on_surface">{{ item.product_name || item.name }}</p>
              <p class="text-sm text-on_surface_variant">Qty: {{ item.quantity }} × ₹ {{ Number(item.unit_price || item.price).toFixed(2) }}</p>
            </div>
            <p class="font-semibold text-on_surface">₹ {{ (Number(item.unit_price || item.price) * item.quantity).toFixed(2) }}</p>
          </div>
        </div>
        <div class="mt-6 pt-4 border-t border-outline-variant/20">
          <div class="flex justify-between text-lg font-bold text-on_surface">
            <span>Total</span>
            <span class="text-primary">₹ {{ Number(order.total_amount || order.total || 0).toFixed(2) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
