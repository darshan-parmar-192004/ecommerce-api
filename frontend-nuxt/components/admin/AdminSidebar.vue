<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '~/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const navigation = [
  { name: 'Dashboard', href: '/admin', icon: 'pi-th-large' },
  { name: 'Products', href: '/admin/products', icon: 'pi-box' },
  { name: 'Categories', href: '/admin/categories', icon: 'pi-tags' },
  { name: 'Inventory', href: '/admin/inventory', icon: 'pi-shopping-cart' },
  { name: 'Orders', href: '/admin/orders', icon: 'pi-clipboard' },
]

const isActive = (href) => route.path === href

const logout = () => {
  authStore.clearAuth()
  router.push('/')
}
</script>

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
        <i :class="'pi ' + item.icon" />
        {{ item.name }}
      </NuxtLink>
    </div>
    <div class="mt-4">
      <Button
        @click="logout"
        severity="danger"
        text
        class="w-full justify-start gap-3 px-3 py-2"
      >
        <i class="pi pi-sign-out" />
        Logout
      </Button>
    </div>
  </nav>
</template>
