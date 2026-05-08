<script setup>
import { useAuthStore } from '~/stores/auth'
import { useUserStore } from '~/stores/user'
import { Package, Clock, CheckCircle, XCircle, FileText } from 'lucide-vue-next'

definePageMeta({
  middleware: ['auth']
})

const userStore = useUserStore()
const authStore = useAuthStore()

onMounted(async () => {
  await Promise.all([
    userStore.fetchOrders(),
    userStore.fetchProfile()
  ])
})

const recentOrders = computed(() => userStore.orders.slice(0, 5))

const stats = computed(() => ({
  totalOrders: userStore.orders.length,
  pending: userStore.orders.filter(o => o.status === 'pending' || o.status === 'processing').length,
  completed: userStore.orders.filter(o => o.status === 'delivered' || o.status === 'completed').length,
  cancelled: userStore.orders.filter(o => o.status === 'cancelled').length
}))

const getStatusIcon = (status) => {
  if (status === 'delivered' || status === 'completed') return CheckCircle
  if (status === 'cancelled') return XCircle
  return Clock
}

const getStatusClass = (status) => {
  if (status === 'delivered' || status === 'completed') return 'text-green-600 bg-green-100'
  if (status === 'cancelled') return 'text-red-600 bg-red-100'
  return 'text-yellow-600 bg-yellow-100'
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('en-IN', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const formatPrice = (price) => {
  return `₹${Number(price).toFixed(2)}`
}
</script>

<template>
  <div class="space-y-8">
    <div>
      <h1 class="text-3xl font-bold text-on_surface font-display">My Dashboard</h1>
      <p class="text-on_surface_variant mt-1">Welcome back, {{ userStore.profile?.name || authStore.user?.name || 'Customer' }}</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-primary/10 rounded-lg">
            <FileText class="w-6 h-6 text-primary" />
          </div>
          <div>
            <p class="text-2xl font-bold text-on_surface">{{ stats.totalOrders }}</p>
            <p class="text-sm text-on_surface_variant">Total Orders</p>
          </div>
        </div>
      </div>

      <div class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-yellow-100 rounded-lg">
            <Clock class="w-6 h-6 text-yellow-600" />
          </div>
          <div>
            <p class="text-2xl font-bold text-on_surface">{{ stats.pending }}</p>
            <p class="text-sm text-on_surface_variant">Pending</p>
          </div>
        </div>
      </div>

      <div class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-green-100 rounded-lg">
            <CheckCircle class="w-6 h-6 text-green-600" />
          </div>
          <div>
            <p class="text-2xl font-bold text-on_surface">{{ stats.completed }}</p>
            <p class="text-sm text-on_surface_variant">Completed</p>
          </div>
        </div>
      </div>

      <div class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-red-100 rounded-lg">
            <XCircle class="w-6 h-6 text-red-600" />
          </div>
          <div>
            <p class="text-2xl font-bold text-on_surface">{{ stats.cancelled }}</p>
            <p class="text-sm text-on_surface_variant">Cancelled</p>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20">
      <div class="p-6 border-b border-outline-variant/20 flex items-center justify-between">
        <h2 class="text-xl font-semibold text-on_surface">Recent Orders</h2>
        <NuxtLink to="/orders" class="text-primary hover:text-primary/80 font-medium text-sm">
          View All
        </NuxtLink>
      </div>

      <div v-if="recentOrders.length === 0" class="p-12 text-center">
        <Package class="w-12 h-12 text-outline/40 mx-auto mb-4" />
        <p class="text-on_surface_variant">No orders yet</p>
        <NuxtLink to="/products" class="inline-block mt-4 px-6 py-2 bg-primary text-white rounded-lg font-medium hover:bg-primary/90">
          Start Shopping
        </NuxtLink>
      </div>

      <div v-else class="divide-y divide-outline-variant/20">
        <NuxtLink
          v-for="order in recentOrders"
          :key="order.order_id"
          :to="`/orders/${order.order_id}`"
          class="p-6 flex items-center justify-between hover:bg-surface-container transition-colors"
        >
          <div class="flex items-center gap-4">
            <div class="p-2 bg-surface-container rounded-lg">
              <Package class="w-5 h-5 text-on_surface_variant" />
            </div>
            <div>
              <p class="font-medium text-on_surface">Order #{{ order.order_id }}</p>
              <p class="text-sm text-on_surface_variant">{{ formatDate(order.created_at || order.createdAt) }}</p>
            </div>
          </div>
          <div class="flex items-center gap-4">
            <span :class="getStatusClass(order.status)" class="px-3 py-1 rounded-full text-xs font-medium capitalize">
              {{ order.status }}
            </span>
            <span class="font-bold text-on_surface">{{ formatPrice(order.total) }}</span>
          </div>
        </NuxtLink>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <NuxtLink to="/user/profile" class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20 hover:border-primary/30 transition-colors">
        <h3 class="font-semibold text-on_surface mb-2">Profile Settings</h3>
        <p class="text-sm text-on_surface_variant">Update your name, email, phone, and address</p>
      </NuxtLink>

      <NuxtLink to="/orders" class="bg-surface-container-lowest rounded-xl p-6 border border-outline-variant/20 hover:border-primary/30 transition-colors">
        <h3 class="font-semibold text-on_surface mb-2">Order History</h3>
        <p class="text-sm text-on_surface_variant">View all your past orders</p>
      </NuxtLink>
    </div>
  </div>
</template>