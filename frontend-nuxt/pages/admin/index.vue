<script setup>
definePageMeta({
  layout: 'admin',
  middleware: 'admin',
  keepalive: true
})

const stats = ref({ totalOrders: 0, totalProducts: 0, totalCustomers: 0, revenue: 0 })
const loading = ref(false)

const fetchStats = async () => {
  loading.value = true
  try {
    const productsRes = await fetch('/api/admin/products')
    const productsData = await productsRes.json()
    const ordersRes = await fetch('/api/admin/orders')
    const ordersData = await ordersRes.json()
    const orders = ordersData.data || ordersData || []
    const products = productsData.data || productsData || []
    stats.value = {
      totalOrders: orders.length,
      totalProducts: products.length,
      totalCustomers: new Set(orders.map(o => o.customer_id || o.customer?.customer_id)).size,
      revenue: orders.reduce((sum, o) => sum + (o.total_amount || o.total || 0), 0)
    }
  } catch (error) {
    console.error('Failed to fetch dashboard stats:', error)
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
      <div class="bg-surface-container-lowest rounded-lg p-6">
        <p class="text-sm text-on_surface_variant mb-1">Total Orders</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalOrders }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-lg p-6">
        <p class="text-sm text-on_surface_variant mb-1">Total Products</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalProducts }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-lg p-6">
        <p class="text-sm text-on_surface_variant mb-1">Total Customers</p>
        <p class="text-3xl font-bold text-on_surface">{{ loading ? '—' : stats.totalCustomers }}</p>
      </div>
      <div class="bg-surface-container-lowest rounded-lg p-6">
        <p class="text-sm text-on_surface_variant mb-1">Revenue</p>
        <p class="text-3xl font-bold text-on_surface">${{ loading ? '—' : stats.revenue }}</p>
      </div>
    </div>
  </div>
</template>
