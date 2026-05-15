<script setup>
definePageMeta({
  layout: 'admin',
  middleware: 'admin',
  keepalive: true
})

const { admin: adminApi } = useApi()
const stats = ref({ totalOrders: 0, totalProducts: 0, totalCustomers: 0, revenue: 0 })
const { showError } = useErrorHandler()
const loading = ref(false)

const fetchStats = async () => {
  loading.value = true
  try {
    const [productsRes, ordersRes] = await Promise.all([
      adminApi.products.list(),
      adminApi.orders.list()
    ])
    const orders = ordersRes.data || ordersRes || []
    const products = productsRes.data || productsRes || []
    stats.value = {
      totalOrders: orders.length,
      totalProducts: products.length,
      totalCustomers: new Set(orders.map(o => o.customer_id || o.customer?.customer_id)).size,
      revenue: orders.reduce((sum, o) => sum + (o.total_amount || o.total || 0), 0)
    }
  } catch (error) {
    showError(error)
  } finally {
    loading.value = false
  }
}

onMounted(() => { fetchStats() })
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold text-on_surface mb-6">Dashboard</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-6 shadow-ambient">
        <p class="text-sm text-on_surface_variant mb-1">Total Orders</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalOrders }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-6 shadow-ambient">
        <p class="text-sm text-on_surface_variant mb-1">Total Products</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalProducts }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-6 shadow-ambient">
        <p class="text-sm text-on_surface_variant mb-1">Total Customers</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalCustomers }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-6 shadow-ambient">
        <p class="text-sm text-on_surface_variant mb-1">Revenue</p>
        <p class="text-3xl font-bold text-on_surface">${{ loading ? '—' : stats.revenue }}</p>
      </div>
    </div>
  </div>
</template>
