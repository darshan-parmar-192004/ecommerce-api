<script setup>
import ToggleSwitch from 'primevue/toggleswitch'
import { useCartStore } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'
import { useThemeStore } from '~/stores/theme'

const CartDrawer = defineAsyncComponent(() => import('~/components/cart/CartDrawer.vue'))

const cartStore = useCartStore()
const authStore = useAuthStore()
const themeStore = useThemeStore()
const route = useRoute()
const isMenuOpen = ref(false)
const userMenu = ref(null)

const { isAuthenticated, user } = storeToRefs(authStore)

const closeMenu = () => {
  isMenuOpen.value = false
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

const { isDark } = storeToRefs(themeStore)
</script>

<template>
  <div class="min-h-screen flex flex-col relative overflow-hidden bg-surface font-body text-on_surface">
    <header class="sticky top-0 z-40 backdrop-blur-xl bg-surface/70 border-b border-white/10">
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

          <div class="flex items-center gap-3">
            <!-- Cart Button -->
            <button
              @click="cartStore.toggleCart"
              class="relative p-2.5 rounded-full transition-colors duration-300 hover:bg-surface-container text-on_surface_variant"
              aria-label="Shopping cart"
            >
              <i class="pi pi-shopping-cart text-xl" />
              <span
                v-if="cartStore.totalItems > 0"
                class="absolute -top-0.5 -right-0.5 text-[10px] w-5 h-5 rounded-full flex items-center justify-center font-bold animate-bounce-subtle bg-gradient-to-r from-primary to-primary-container text-white"
              >
                {{ cartStore.totalItems }}
              </span>
            </button>

            <!-- Theme Toggle Switch -->
            <ToggleSwitch
              :modelValue="isDark"
              @update:modelValue="themeStore.toggleTheme()"
            >
              <template #handle="{ checked }">
                <i :class="['!text-xs pi', checked ? 'pi-moon' : 'pi-sun']" />
              </template>
            </ToggleSwitch>

            <!-- Auth Links -->
            <template v-if="isAuthenticated">
              <div class="relative">
                <button
                  @click="userMenu.toggle($event)"
                  class="flex items-center gap-2 p-1 rounded-full transition-colors duration-300 hover:bg-surface-container"
                >
                  <Avatar
                    :label="getUserInitial(user?.name)"
                    shape="circle"
                    class="bg-gradient-to-r from-primary to-primary-container text-white font-semibold"
                  />
                </button>

                <OverlayPanel ref="userMenu" :style="{ width: '16rem' }" class="!bg-surface-container-lowest">
                  <div class="p-4 border-b border-outline-variant/20">
                    <p class="text-sm font-semibold text-on_surface">{{ user?.name || 'User' }}</p>
                    <p class="text-xs text-outline">{{ user?.email || '' }}</p>
                  </div>
                  <div class="py-2">
                    <NuxtLink
                      to="/orders"
                      class="flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                    >
                      <i class="pi pi-shopping-cart text-on_surface_variant" />
                      <span class="text-sm font-medium text-on_surface_variant">My Orders</span>
                    </NuxtLink>
                    <NuxtLink
                      to="/user/profile"
                      class="flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                    >
                      <i class="pi pi-user text-on_surface_variant" />
                      <span class="text-sm font-medium text-on_surface_variant">Profile</span>
                    </NuxtLink>
                    <NuxtLink
                      v-if="authStore.isAdmin"
                      to="/admin"
                      class="flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors text-primary"
                    >
                      <i class="pi pi-shield" />
                      <span class="text-sm font-medium">Admin Panel</span>
                    </NuxtLink>
                    <hr class="mx-3 my-1 border-outline-variant/20" />
                    <NuxtLink
                      to="/"
                      class="flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                    >
                      <i class="pi pi-home text-on_surface_variant" />
                      <span class="text-sm font-medium text-on_surface_variant">Back to Website</span>
                    </NuxtLink>
                    <button
                      @click="handleLogout"
                      class="w-full flex items-center gap-3 px-4 py-3 hover:bg-error-container/30 transition-colors text-error"
                    >
                      <i class="pi pi-sign-out" />
                      <span class="text-sm font-medium">Logout</span>
                    </button>
                  </div>
                </OverlayPanel>
              </div>
            </template>
            <template v-else>
              <NuxtLink
                to="/auth/login"
                class="px-5 py-2.5 text-white rounded-md font-medium text-sm transition-all duration-300 bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl hover:scale-[1.02] active:scale-[0.98]"
              >
                Login
              </NuxtLink>
              <NuxtLink
                v-if="authStore.isAdmin"
                to="/admin"
                class="text-sm font-medium text-primary"
              >
                Admin Panel
              </NuxtLink>
            </template>
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
