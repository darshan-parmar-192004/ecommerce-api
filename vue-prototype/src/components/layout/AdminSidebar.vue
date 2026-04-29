<script setup>
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const logout = async () => {
  await authStore.logout()
  router.push({ name: 'Login' })
}
</script>

<template>
  <aside class="w-64 bg-white border-r border-gray-200 min-h-screen">
    <div class="p-6">
      <h1 class="text-xl font-bold text-gray-900">Admin Panel</h1>
    </div>
    <nav class="px-4 space-y-1">
      <RouterLink
        to="/admin"
        class="block px-4 py-2 text-sm font-medium text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        active-class="bg-gray-100 text-gray-900"
      >
        Dashboard
      </RouterLink>
      <RouterLink
        to="/admin/products"
        class="block px-4 py-2 text-sm font-medium text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        active-class="bg-gray-100 text-gray-900"
      >
        Products
      </RouterLink>
      <RouterLink
        to="/admin/inventory"
        class="block px-4 py-2 text-sm font-medium text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
        active-class="bg-gray-100 text-gray-900"
      >
        Inventory
      </RouterLink>
      <hr class="my-4 border-gray-200" />
      <RouterLink
        to="/"
        class="block px-4 py-2 text-sm font-medium text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
      >
        Back to Store
      </RouterLink>
      <button
        @click="logout"
        class="w-full text-left px-4 py-2 text-sm font-medium text-red-600 rounded-lg hover:bg-red-50 transition-colors"
      >
        Logout
      </button>
    </nav>
  </aside>
</template>
