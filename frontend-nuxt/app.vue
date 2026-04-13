<script setup>
import { useCartStore } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'

const cartStore = useCartStore()
const authStore = useAuthStore()

onMounted(() => {
  // Load auth state from localStorage
  authStore.loadAuth()

  // Initialize cart from localStorage
  cartStore.init()

  // Verify auth token if exists
  if (authStore.token) {
    authStore.verifyAuth()
  }
})
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  <Toast />
</template>

<style>
/* Hide scrollbar but allow scrolling */
* {
  scrollbar-width: none;
  -ms-overflow-style: none;
}

*::-webkit-scrollbar {
  display: none;
}

.page-enter-active,
.page-leave-active {
  transition: opacity 400ms ease-out, transform 400ms ease-out;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(12px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
