<script setup>
definePageMeta({
  middleware: ['auth']
})

const { orders: ordersApi } = useApi()

const { data: ordersData, pending, error } = await useAsyncData('user-orders', () => ordersApi.list())

useSeoMeta({
  title: 'Dashboard - E-Commerce Store'
})

const orders = computed(() => ordersData.value?.data || ordersData.value || [])

const stats = computed(() => ({
  totalOrders: orders.value.length,
  pending: orders.value.filter(o => o.status === 'pending' || o.status === 'processing').length,
  delivered: orders.value.filter(o => o.status === 'delivered').length,
  cancelled: orders.value.filter(o => o.status === 'cancelled').length
}))

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })
}
</script>

<template>
  <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-on_surface font-display mb-8">My Dashboard</h1>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-on_surface_variant">Total Orders</p>
            <p class="text-3xl font-bold text-on_surface">{{ stats.totalOrders }}</p>
          </div>
          <i class="pi pi-shopping-bag text-4xl text-primary/30" />
        </div>
      </div>
      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-on_surface_variant">Pending</p>
            <p class="text-3xl font-bold text-yellow-600">{{ stats.pending }}</p>
          </div>
          <i class="pi pi-clock text-4xl text-yellow-600/30" />
        </div>
      </div>
      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-on_surface_variant">Delivered</p>
            <p class="text-3xl font-bold text-green-600">{{ stats.delivered }}</p>
          </div>
          <i class="pi pi-check-circle text-4xl text-green-600/30" />
        </div>
      </div>
      <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-on_surface_variant">Cancelled</p>
            <p class="text-3xl font-bold text-red-600">{{ stats.cancelled }}</p>
          </div>
          <i class="pi pi-times-circle text-4xl text-red-600/30" />
        </div>
      </div>
    </div>

    <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
      <h2 class="text-xl font-semibold text-on_surface mb-4">Recent Orders</h2>
      <div v-if="pending" class="space-y-3">
        <div v-for="i in 3" :key="i" class="h-16 bg-surface-container rounded-lg animate-pulse" />
      </div>
      <div v-else-if="orders.length === 0" class="text-center py-8">
        <i class="pi pi-file text-5xl text-outline mb-4" />
        <p class="text-on_surface_variant">No orders yet</p>
      </div>
      <div v-else class="overflow-x-auto">
        <DataTable :value="orders" class="w-full">
          <Column field="order_id" header="Order ID" />
          <Column header="Date">
            <template #body="{ data }">{{ formatDate(data.created_at || data.order_date) }}</template>
          </Column>
          <Column header="Status">
            <template #body="{ data }">
              <Tag :value="data.status || 'Pending'" :severity="data.status === 'delivered' ? 'success' : data.status === 'cancelled' ? 'danger' : 'warn'" />
            </template>
          </Column>
          <Column header="Total">
            <template #body="{ data }">₹ {{ Number(data.total_amount || data.total || 0).toFixed(2) }}</template>
          </Column>
          <Column header="">
            <template #body="{ data }">
              <NuxtLink :to="`/orders/${data.order_id || data.id}`">
                <Button label="View" icon="pi pi-eye" text size="small" />
              </NuxtLink>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>
  </div>
</template>
