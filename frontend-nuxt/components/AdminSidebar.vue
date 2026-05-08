<template>
  <nav class="w-64 bg-surface border-r border-outline-variant/20 p-4 flex flex-col">
    <div class="flex-1 space-y-1">
      <NuxtLink
        v-for="item in navigation"
        :key="item.name"
        :to="item.href"
        class="flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium transition-colors"
        :class="isActive(item.href) ? 'bg-primary-container text-primary' : 'text-on_surface_variant hover:bg-surface-container hover:text-on_surface'"
      >
        <component :is="item.icon" class="w-5 h-5" />
        {{ item.name }}
      </NuxtLink>
    </div>
    <div class="mt-4">
      <button @click="logout" class="flex items-center gap-3 w-full px-3 py-2 text-sm font-medium text-error hover:bg-error-container/30">
        <LogOut class="w-5 h-5" />
        Logout
      </button>
    </div>
  </nav>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '~/stores/auth'
import { LayoutDashboard, Package, Tags, ShoppingCart, LogOut, ClipboardList } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const navigation = [
  { name: 'Dashboard', href: '/admin', icon: LayoutDashboard },
  { name: 'Products', href: '/admin/products', icon: Package },
  { name: 'Categories', href: '/admin/categories', icon: Tags },
  { name: 'Inventory', href: '/admin/inventory', icon: ShoppingCart },
  { name: 'Orders', href: '/admin/orders', icon: ClipboardList }
]

const isActive = (href) => route.path === href

const logout = () => {
  authStore.clearAuth()
  router.push('/')
}
</script>

<style scoped>
/* Add any needed styles */
</style>