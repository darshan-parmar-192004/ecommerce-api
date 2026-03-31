<script setup>
const cartStore = useCartStore()
const authStore = useAuthStore()
const route = useRoute()
const isMenuOpen = ref(false)
const isUserMenuOpen = ref(false)

const { isAuthenticated, user } = storeToRefs(authStore)

onMounted(async () => {
  authStore.loadAuth()
  if (authStore.token) {
    await authStore.verifyAuth()
    cartStore.loadCart()
  }
})

const closeMenu = () => {
  isMenuOpen.value = false
  isUserMenuOpen.value = false
}

const getUserInitial = (name) => {
  if (!name) return 'U'
  return name.charAt(0).toUpperCase()
}

const handleLogout = async () => {
  try {
    await fetch('/api/auth/logout', { method: 'POST' })
  } catch (e) {
    console.error('Logout error:', e)
  }
  authStore.clearAuth()
  cartStore.clearCart()
  closeMenu()
  navigateTo('/')
}
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <AnimatedGrid variant="minimal" />
    
    <header class="bg-white/80 backdrop-blur-md shadow-sm sticky top-0 z-40 border-b border-gray-100">
      <nav class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center gap-8">
            <NuxtLink to="/" class="text-xl font-bold bg-gradient-to-r from-primary-600 to-accent-600 bg-clip-text text-transparent" @click="closeMenu">
              Store
            </NuxtLink>

            <div class="hidden md:flex items-center gap-6">
              <NuxtLink 
                to="/products" 
                class="text-gray-600 hover:text-primary-600 transition-colors font-medium"
                active-class="text-primary-600"
              >
                Products
              </NuxtLink>
              <NuxtLink 
                to="/cart" 
                class="text-gray-600 hover:text-primary-600 transition-colors font-medium"
                active-class="text-primary-600"
              >
                Cart
              </NuxtLink>
            </div>
          </div>

          <div class="flex items-center gap-4">
            <button 
              @click="cartStore.toggleCart"
              class="relative p-2 text-gray-600 hover:text-primary-600 transition-colors"
              aria-label="Shopping cart"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              <span 
                v-if="cartStore.totalItems > 0"
                class="absolute -top-1 -right-1 bg-primary-600 text-white text-xs w-5 h-5 rounded-full flex items-center justify-center font-medium animate-bounce-subtle"
              >
                {{ cartStore.totalItems }}
              </span>
            </button>

            <template v-if="isAuthenticated">
              <!-- User Dropdown -->
              <div class="relative">
                <button 
                  @click="isUserMenuOpen = !isUserMenuOpen"
                  class="flex items-center gap-2 p-1.5 rounded-full hover:bg-gray-100 transition-colors"
                >
                  <div class="w-8 h-8 bg-gradient-to-r from-primary-500 to-accent-500 rounded-full flex items-center justify-center text-white font-semibold text-sm">
                    {{ getUserInitial(user?.name) }}
                  </div>
                  <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>

                <Transition
                  enter-active-class="transition ease-out duration-200"
                  enter-from-class="opacity-0 scale-95"
                  enter-to-class="opacity-100 scale-100"
                  leave-active-class="transition ease-in duration-150"
                  leave-from-class="opacity-100 scale-100"
                  leave-to-class="opacity-0 scale-95"
                >
                  <div v-if="isUserMenuOpen" class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-lg border border-gray-100 py-2 z-50">
                    <div class="px-4 py-3 border-b border-gray-100">
                      <p class="text-sm font-medium text-gray-900">{{ user?.name || 'User' }}</p>
                      <p class="text-xs text-gray-500">{{ user?.email || '' }}</p>
                    </div>
                    <NuxtLink 
                      to="/orders" 
                      class="flex items-center gap-3 px-4 py-2 text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
                      @click="closeMenu"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                      </svg>
                      My Orders
                    </NuxtLink>
                    <button 
                      @click="handleLogout"
                      class="w-full flex items-center gap-3 px-4 py-2 text-red-600 hover:bg-red-50 transition-colors"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                      </svg>
                      Logout
                    </button>
                  </div>
                </Transition>
              </div>
            </template>
            <template v-else>
              <NuxtLink to="/auth/login" class="text-gray-600 hover:text-primary-600 font-medium hidden sm:block">
                Login
              </NuxtLink>
              <NuxtLink to="/auth/register" class="px-4 py-2 bg-gradient-to-r from-primary-600 to-accent-600 text-white rounded-lg hover:from-primary-700 hover:to-accent-700 transition-all font-medium text-sm shadow-md shadow-primary-500/25 hover:shadow-lg">
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
              <NuxtLink to="/products" class="text-gray-600 hover:text-primary-600 font-medium" @click="closeMenu">
                Products
              </NuxtLink>
              <NuxtLink to="/cart" class="text-gray-600 hover:text-primary-600 font-medium" @click="closeMenu">
                Cart
              </NuxtLink>
              <template v-if="isAuthenticated">
                <NuxtLink to="/orders" class="text-gray-600 hover:text-primary-600 font-medium" @click="closeMenu">
                  My Orders
                </NuxtLink>
                <button @click="handleLogout" class="text-left text-red-600 hover:text-red-700 font-medium">
                  Logout
                </button>
              </template>
              <template v-else>
                <NuxtLink to="/auth/login" class="text-gray-600 hover:text-primary-600 font-medium" @click="closeMenu">
                  Login
                </NuxtLink>
                <NuxtLink to="/auth/register" class="text-primary-600 font-medium" @click="closeMenu">
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

<style scoped>
.animate-bounce-subtle {
  animation: bounceSubtle 0.5s ease-out;
}

@keyframes bounceSubtle {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}
</style>
