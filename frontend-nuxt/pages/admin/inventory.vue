<script setup>
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

const { admin: adminApi } = useApi()
const inventory = ref([])
const loading = ref(true)

const fetchInventory = async () => {
  loading.value = true
  try {
    const response = await adminApi.products.list()
    const products = response.data || response || []
    inventory.value = products.map(product => ({
      product_name: product.name,
      product_id: product.product_id,
      warehouse_id: product.warehouse_id || 'N/A',
      quantity: product.stock || 0,
      last_updated: null
    }))
  } catch (err) {
    console.error('Failed to fetch inventory', err)
  } finally {
    loading.value = false
  }
}

onMounted(fetchInventory)
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold text-on_surface mb-8">Inventory Management</h1>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-surface-container rounded-xl animate-pulse" />
    </div>

    <div v-else class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 overflow-hidden">
      <DataTable :value="inventory" stripedRows class="w-full">
        <Column field="product_name" header="Product Name" />
        <Column field="product_id" header="Product ID" />
        <Column field="warehouse_id" header="Warehouse" />
        <Column field="quantity" header="Current Stock" />
        <Column header="Status">
          <template #body="{ data }">
            <Tag 
              :value="data.quantity > 10 ? 'In Stock' : data.quantity > 0 ? 'Low Stock' : 'Out of Stock'"
              :severity="data.quantity > 10 ? 'success' : data.quantity > 0 ? 'warn' : 'danger'"
            />
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>
