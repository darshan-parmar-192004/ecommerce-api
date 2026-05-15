<script setup>
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

import { statusMessages } from '~/constants/errorMessages'

const { admin: adminApi } = useApi()
const { success: showSuccess, error: showError } = useAppToast()

const orders = ref([])
const loading = ref(true)

const fetchOrders = async () => {
  loading.value = true
  try {
    const response = await adminApi.orders.list()
    orders.value = response.data || response || []
  } catch (err) {
    console.error('Failed to fetch orders', err)
  } finally {
    loading.value = false
  }
}

const statusSeverity = (status) => {
  const map = {
    pending: 'warn',
    confirmed: 'info',
    processing: 'info',
    shipped: 'success',
    delivered: 'success',
    cancelled: 'danger',
    refunded: 'danger'
  }
  return map[status?.toLowerCase()] || 'info'
}

const updateStatus = async (order, newStatus) => {
  try {
    await adminApi.orders.updateStatus(order.order_id || order.id, newStatus)
    showSuccess(`Order ${order.order_id || order.id} → ${newStatus}`)
    await fetchOrders()
  } catch (err) {
    showError(err.message || statusMessages.orderUpdateFailed)
  }
}

onMounted(fetchOrders)
</script>

<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-on_surface">Manage Orders</h1>
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-surface-container rounded-lg animate-pulse" />
    </div>

    <div v-else-if="orders.length === 0" class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-12 text-center">
      <i class="pi pi-clipboard text-5xl text-outline mb-4" />
      <h2 class="text-xl font-medium text-on_surface mb-2">No orders yet</h2>
      <p class="text-on_surface_variant">Orders will appear here once customers start purchasing.</p>
    </div>

    <div v-else class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 overflow-hidden">
      <DataTable :value="orders" stripedRows class="w-full">
        <Column field="order_id" header="Order ID" sortable />
        <Column header="Customer">
          <template #body="{ data }">
            {{ data.customer_name || data.customer?.name || data.customer_id || '—' }}
          </template>
        </Column>
        <Column header="Total">
          <template #body="{ data }">
            ${{ (data.total_amount || data.total || 0).toFixed(2) }}
          </template>
        </Column>
        <Column field="status" header="Status" sortable>
          <template #body="{ data }">
            <Tag :value="data.status" :severity="statusSeverity(data.status)" />
          </template>
        </Column>
        <Column field="created_at" header="Date" sortable>
          <template #body="{ data }">
            {{ data.created_at ? new Date(data.created_at).toLocaleDateString() : '—' }}
          </template>
        </Column>
        <Column header="Actions">
          <template #body="{ data }">
            <div class="flex gap-2">
              <Button
                v-if="data.status !== 'shipped' && data.status !== 'delivered' && data.status !== 'cancelled'"
                @click="updateStatus(data, 'shipped')"
                icon="pi pi-truck"
                text
                rounded
                size="small"
                v-tooltip.top="'Mark as Shipped'"
              />
              <Button
                v-if="data.status === 'shipped'"
                @click="updateStatus(data, 'delivered')"
                icon="pi pi-check-circle"
                text
                rounded
                size="small"
                severity="success"
                v-tooltip.top="'Mark as Delivered'"
              />
              <Button
                v-if="data.status !== 'cancelled' && data.status !== 'delivered'"
                @click="updateStatus(data, 'cancelled')"
                icon="pi pi-times-circle"
                text
                rounded
                size="small"
                severity="danger"
                v-tooltip.top="'Cancel Order'"
              />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>
