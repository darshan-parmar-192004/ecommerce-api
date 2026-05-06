<script setup>
import { useCartStore } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'

const cartStore = useCartStore()
const authStore = useAuthStore()
const route = useRoute()
const isMenuOpen = ref(false)
const isUserMenuOpen = ref(false)
const layoutEl = ref(null)

const { isAuthenticated, user } = storeToRefs(authStore)

onMounted(() => {
  if (layoutEl.value) {
    layoutEl.value.addEventListener('mousemove', (e) => {
      const rect = layoutEl.value.getBoundingClientRect()
      layoutEl.value.style.setProperty('--x', `${e.clientX - rect.left}px`)
      layoutEl.value.style.setProperty('--y', `${e.clientY - rect.top}px`)
    })
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

const { auth } = useApi()

const handleLogout = async () => {
  try {
    await auth.logout()
  } catch (e) {
    console.error('Logout error:', e)
  } finally {
    authStore.clearAuth()
    cartStore.clearCart()
    closeMenu()
    navigateTo('/')
  }
}
</script>

<template>
  <div ref="layoutEl" class="min-h-screen flex flex-col relative overflow-hidden bg-surface font-body text-on_surface">
    <!-- Debug info -->
    <div v-if="isAuthenticated" class="fixed top-4 right-4 bg-red-500 text-white p-2 rounded text-xs z-50">
      Debug: Role: {{ user?.role }}, IsAdmin: {{ authStore.isAdmin }}
    </div>
    <div class="pointer-events-none fixed inset-0 z-0" style="background: radial-gradient(600px circle at var(--x, 50%) var(--y, 50%), rgba(62, 81, 251, 0.06), transparent 40%);"></div>
    <AnimatedGrid variant="minimal" />
    
    <header class="sticky top-0 z-40 glass backdrop-blur-xl bg-surface/70 border-b border-white/10">
      <nav class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="flex justify-between items-center h-20">
          <div class="flex items-center gap-12">
            <NuxtLink to="/" class="text-2xl font-bold tracking-tight font-display text-primary" @click="closeMenu">
              The Curator
            </NuxtLink>

            <div class="hidden md:flex items-center gap-8">
              <NuxtLink 
                to="/products" 
                class="text-xs font-semibold uppercase tracking-widest transition-colors duration-300 text-on_surface_variant"
                active-class="text-primary"
              >
                Products
              </NuxtLink>
              <NuxtLink 
                to="/cart" 
                class="text-xs font-semibold uppercase tracking-widest transition-colors duration-300 text-on_surface_variant"
                active-class="text-primary"
              >
                Cart
              </NuxtLink>
            </div>
          </div>

          <div class="flex items-center gap-5">
            <button 
              @click="cartStore.toggleCart"
              class="relative p-2.5 rounded-full transition-colors duration-300 hover:bg-surface-container text-on_surface_variant"
              aria-label="Shopping cart"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
              </svg>
              <span 
                v-if="cartStore.totalItems > 0"
                class="absolute -top-0.5 -right-0.5 text-[10px] w-5 h-5 rounded-full flex items-center justify-center font-bold animate-bounce-subtle bg-gradient-to-r from-primary to-primary-container text-white"
              >
                {{ cartStore.totalItems }}
              </span>
            </button>

            <template v-if="isAuthenticated">
              <div class="relative">
                <button 
                  @click="isUserMenuOpen = !isUserMenuOpen"
                  class="flex items-center gap-2 p-1 rounded-full transition-colors duration-300 hover:bg-surface-container"
                >
                  <div class="w-9 h-9 rounded-full flex items-center justify-center text-white font-semibold text-sm bg-gradient-to-r from-primary to-primary-container">
                    {{ getUserInitial(user?.name) }}
                  </div>
                  <svg class="w-4 h-4 transition-transform duration-300 text-outline" :class="{ 'rotate-180': isUserMenuOpen }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>

                <Transition
                  enter-active-class="transition ease-out duration-300"
                  enter-from-class="opacity-0 scale-95 -translate-y-1"
                  enter-to-class="opacity-100 scale-100 translate-y-0"
                  leave-active-class="transition ease-in duration-200"
                  leave-from-class="opacity-100 scale-100 translate-y-0"
                  leave-to-class="opacity-0 scale-95 -translate-y-1"
                >
                  <div v-if="isUserMenuOpen" class="absolute right-0 mt-3 w-64 rounded-2xl py-3 z-50 bg-surface-container-lowest shadow-ambient">
                    <div class="px-5 py-4 border-b border-outline-variant/20">
                      <p class="text-sm font-semibold text-on_surface">{{ user?.name || 'User' }}</p>
                      <p class="text-xs text-outline">{{ user?.email || '' }}</p>
                    </div>
                    <NuxtLink 
                      to="/orders" 
                      class="flex items-center gap-3 px-5 py-3 transition-colors duration-200 hover:bg-surface-container-low text-on_surface_variant"
                      @click="closeMenu"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
                      </svg>
                      <span class="text-sm font-medium">My Orders</span>
                    </NuxtLink>
                    <NuxtLink 
                      to="/user/profile" 
                      class="flex items-center gap-3 px-5 py-3 transition-colors duration-200 hover:bg-surface-container-low text-on_surface_variant"
                      @click="closeMenu"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                      </svg>
                      <span class="text-sm font-medium">Profile</span>
                    </NuxtLink>
                    <NuxtLink 
                      v-if="authStore.isAdmin"
                      to="/admin" 
                      class="flex items-center gap-3 px-5 py-3 transition-colors duration-200 hover:bg-surface-container-low text-primary"
                      @click="closeMenu"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.321 3.35.108z" />
                      </svg>
                      <span class="text-sm font-medium">Admin Panel</span>
                    </NuxtLink>
                    <button 
                      @click="handleLogout"
                      class="w-full flex items-center gap-3 px-5 py-3 transition-colors duration-200 hover:bg-error-container/30 text-error"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15m3 0l3-3m0 0l-3-3m3 3H9" />
                      </svg>
                      <span class="text-sm font-medium">Logout</span>
                    </button>
                  </div>
                </Transition>
              </div>
            </template>
            <template v-else>
              <NuxtLink to="/auth/login" class="text-sm font-medium transition-colors duration-300 hidden sm:block hover:text-primary text-on_surface_variant">
                Login
              </NuxtLink>
              <NuxtLink to="/auth/register" class="px-5 py-2.5 text-white rounded-md font-medium text-sm transition-all duration-300 bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl hover:scale-[1.02] active:scale-[0.98]">
                Register
              </NuxtLink>
<NuxtLink 
                      to="/admin" 
                      v-if="authStore.isAdmin"
                      class="text-sm font-medium text-primary"
                    >
                      Admin Panel
                    </NuxtLink>
                  </template>

            <button 
              @click="isMenuOpen = !isMenuOpen"
              class="md:hidden p-2 rounded-full transition-colors duration-300 hover:bg-surface-container text-on_surface_variant"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path v-if="!isMenuOpen" stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </nav>
    </header>

    <main class="flex-1">
      <slot />
    </main>

    <footer class="mt-auto bg-surface-container-low">
      <div class="max-w-7xl mx-auto px-6 lg:px-8 py-16">
        <div class="text-center">
          <p class="text-xs uppercase tracking-widest text-outline font-body">
            &copy; 2024 The Curator. All rights reserved.
          </p>
        </div>
      </div>
    </footer>

    <CartDrawer />
  </div>
</template>

<style scoped>
@keyframes bounceSubtle {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

.animate-bounce-subtle {
  animation: bounceSubtle 0.5s ease-out;
}
</style>
