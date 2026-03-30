<script setup>
const cartStore = useCartStore()
const authStore = useAuthStore()
const route = useRoute()
const isMenuOpen = ref(false)

onMounted(() => {
  cartStore.loadCart()
  authStore.loadAuth()
})

const closeMenu = () => {
  isMenuOpen.value = false
}
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <header class="bg-white shadow-sm sticky top-0 z-40">
      <nav class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center gap-8">
            <NuxtLink to="/" class="text-xl font-bold text-indigo-600" @click="closeMenu">
              Store
            </NuxtLink>

            <div class="hidden md:flex items-center gap-6">
              <NuxtLink 
                to="/products" 
                class="text-gray-600 hover:text-indigo-600 transition-colors font-medium"
                active-class="text-indigo-600"
              >
                Products
              </NuxtLink>
              <NuxtLink 
                to="/cart" 
                class="text-gray-600 hover:text-indigo-600 transition-colors font-medium"
                active-class="text-indigo-600"
              >
                Cart
              </NuxtLink>
            </div>
          </div>

          <div class="flex items-center gap-4">
            <button 
              @click="cartStore.toggleCart"
              class="relative p-2 text-gray-600 hover:text-indigo-600 transition-colors"
              aria-label="Shopping cart"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              <span 
                v-if="cartStore.totalItems > 0"
                class="absolute -top-1 -right-1 bg-indigo-600 text-white text-xs w-5 h-5 rounded-full flex items-center justify-center font-medium"
              >
                {{ cartStore.totalItems }}
              </span>
            </button>

            <template v-if="authStore.isAuthenticated">
              <NuxtLink to="/orders" class="hidden md:block text-gray-600 hover:text-indigo-600 font-medium">
                My Orders
              </NuxtLink>
              <button 
                @click="authStore.clearAuth(); navigateTo('/')"
                class="text-gray-600 hover:text-indigo-600 font-medium"
              >
                Logout
              </button>
            </template>
            <template v-else>
              <NuxtLink to="/auth/login" class="text-gray-600 hover:text-indigo-600 font-medium hidden sm:block">
                Login
              </NuxtLink>
              <NuxtLink to="/auth/register" class="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors font-medium text-sm">
                Register
              </NuxtLink>
            </template>

            <button 
              @click="isMenuOpen = !isMenuOpen"
              class="md:hidden p-2 text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="!isMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <Transition
          enter-active-class="transition duration-200 ease-out"
          enter-from-class="opacity-0 -translate-y-2"
          enter-to-class="opacity-100 translate-y-0"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="opacity-100 translate-y-0"
          leave-to-class="opacity-0 -translate-y-2"
        >
          <div v-if="isMenuOpen" class="md:hidden py-4 border-t">
            <div class="flex flex-col gap-4">
              <NuxtLink to="/products" class="text-gray-600 hover:text-indigo-600 font-medium" @click="closeMenu">
                Products
              </NuxtLink>
              <NuxtLink to="/cart" class="text-gray-600 hover:text-indigo-600 font-medium" @click="closeMenu">
                Cart
              </NuxtLink>
              <template v-if="authStore.isAuthenticated">
                <NuxtLink to="/orders" class="text-gray-600 hover:text-indigo-600 font-medium" @click="closeMenu">
                  My Orders
                </NuxtLink>
                <button @click="authStore.clearAuth(); closeMenu(); navigateTo('/')" class="text-left text-gray-600 hover:text-indigo-600 font-medium">
                  Logout
                </button>
              </template>
              <template v-else>
                <NuxtLink to="/auth/login" class="text-gray-600 hover:text-indigo-600 font-medium" @click="closeMenu">
                  Login
                </NuxtLink>
                <NuxtLink to="/auth/register" class="text-indigo-600 font-medium" @click="closeMenu">
                  Register
                </NuxtLink>
              </template>
            </div>
          </div>
        </Transition>
      </nav>
    </header>

    <main class="flex-1">
      <slot />
    </main>

    <footer class="bg-white border-t border-gray-200 mt-auto">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <p class="text-center text-gray-500 text-sm">
          &copy; 2024 E-Commerce Store. All rights reserved.
        </p>
      </div>
    </footer>

    <CartDrawer />
  </div>
</template>