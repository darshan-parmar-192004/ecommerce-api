<script setup>
import { ref, watch } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import { useCart } from '@/composables/useCart'
import ThemeToggle from '@/components/common/ThemeToggle.vue'
import { ShoppingCart, ChevronDown, Menu, X } from 'lucide-vue-next'

const { user, isAuthenticated, isAdmin, logout: authLogout } = useAuth()
const { cartCount, toggleDrawer } = useCart()
const router = useRouter()
const mobileMenuOpen = ref(false)

// Close mobile menu on route change
watch(() => router.currentRoute.value, () => {
  mobileMenuOpen.value = false
})

const logout = async () => {
  await authLogout()
  router.push({ name: 'Login' })
  mobileMenuOpen.value = false
}
</script>

<template>
    <nav class="sticky top-0 z-40 bg-white/80 dark:bg-brand-800/80 backdrop-blur-md border-b border-gray-200 dark:border-brand-700 shadow-sm">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <!-- Logo -->
        <div class="flex items-center gap-8">
          <RouterLink
             to="/"
             class="text-xl font-bold text-gray-900 dark:text-gray-100 hover:text-gray-700 dark:hover:text-gray-300 transition-colors"
             @click="mobileMenuOpen = false"
           >
             E-Commerce
           </RouterLink>
           <div class="hidden md:flex items-center gap-6">
             <RouterLink
               to="/products"
               class="text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors relative group"
             >
               Products
               <span class="absolute bottom-0 left-0 w-0 h-0.5 bg-gray-900 dark:bg-gray-100 group-hover:w-full transition-all duration-300" />
             </RouterLink>
              <RouterLink
                v-if="isAdmin"
                to="/admin"
               class="text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors relative group"
             >
               Admin Panel
               <span class="absolute bottom-0 left-0 w-0 h-0.5 bg-gray-900 dark:bg-gray-100 group-hover:w-full transition-all duration-300" />
             </RouterLink>
           </div>
        </div>

        <!-- Right side -->
        <div class="flex items-center gap-4">
          <!-- Theme Toggle -->
          <ThemeToggle />

          <!-- Cart -->
            <button
              @click="toggleDrawer()"
              class="relative p-2 text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-all duration-200 hover:scale-110"
            >
              <ShoppingCart class="w-6 h-6" />
              <Transition name="badge">
                <span
                  v-if="cartCount > 0"
                  class="absolute -top-1 -right-1 bg-gray-900 dark:bg-brand-700 text-white dark:text-gray-100 text-xs w-5 h-5 rounded-full flex items-center justify-center animate-bounce"
                >
                  {{ cartCount }}
                </span>
             </Transition>
           </button>

          <!-- Auth Section -->
            <template v-if="isAuthenticated">
              <div class="relative group">
                <button
                  class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors px-3 py-2 rounded-lg hover:bg-gray-50 dark:hover:bg-brand-700"
                >
                   <div class="w-8 h-8 bg-gray-900 dark:bg-brand-700 text-white dark:text-gray-100 rounded-full flex items-center justify-center text-xs font-bold">
                     {{ (user?.name || 'U').charAt(0).toUpperCase() }}
                   </div>
                   <span class="hidden sm:inline">{{ user?.name?.split(' ')[0] || 'Account' }}</span>
                  <ChevronDown class="w-4 h-4" />
               </button>
               <!-- Dropdown menu - works on hover for desktop, click for mobile via group -->
               <div class="absolute right-0 mt-2 w-48 bg-white dark:bg-brand-800 rounded-xl shadow-xl border border-gray-100 dark:border-brand-700 py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
                  <RouterLink v-if="isAdmin" to="/admin" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors rounded-lg mx-2">
                   Admin Panel
                 </RouterLink>
                 <RouterLink to="/user/dashboard" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors rounded-lg mx-2">
                   Dashboard
                 </RouterLink>
                 <RouterLink to="/user/orders" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors rounded-lg mx-2">
                   Orders
                 </RouterLink>
                 <RouterLink to="/user/profile" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors rounded-lg mx-2">
                   Profile
                 </RouterLink>
                 <hr class="my-2 border-gray-100 dark:border-brand-700" />
                 <button @click="logout" class="block w-full text-left px-4 py-2.5 text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors rounded-lg mx-2">
                   Logout
                 </button>
               </div>
             </div>
           </template>
           <template v-else>
             <RouterLink to="/auth/login" class="text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
               Login
             </RouterLink>
             <RouterLink to="/auth/register" class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-gray-900 dark:bg-brand-700 rounded-lg hover:bg-gray-800 dark:hover:bg-brand-600 transition-all duration-200 hover:shadow-lg hover:-translate-y-0.5">
               Sign Up
             </RouterLink>
           </template>

           <!-- Mobile menu button -->
           <button
             @click="mobileMenuOpen = !mobileMenuOpen"
             class="md:hidden p-2 text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
             :aria-label="mobileMenuOpen ? 'Close menu' : 'Open menu'"
           >
             <Menu v-if="!mobileMenuOpen" class="w-6 h-6" />
             <X v-else class="w-6 h-6" />
           </button>
         </div>
       </div>

      <!-- Mobile Menu -->
      <Transition
        name="mobile-menu"
        enter-active-class="transition-all duration-300 ease-out"
        leave-active-class="transition-all duration-200 ease-in"
        enter-from-class="opacity-0 max-h-0"
        enter-to-class="opacity-100 max-h-96"
        leave-from-class="opacity-100 max-h-96"
        leave-to-class="opacity-0 max-h-0"
      >
        <div v-if="mobileMenuOpen" class="md:hidden overflow-hidden">
           <div class="py-4 space-y-1 border-t border-gray-100 dark:border-brand-700">
               <RouterLink
                 to="/products"
                 class="block px-4 py-2.5 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors"
                 @click="mobileMenuOpen = false"
               >
                 Products
               </RouterLink>
             <template v-if="isAuthenticated">
                <RouterLink v-if="isAdmin" to="/admin" class="block px-4 py-2.5 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors" @click="mobileMenuOpen = false">
                 Admin Panel
               </RouterLink>
               <RouterLink to="/user/dashboard" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors" @click="mobileMenuOpen = false">
                 Dashboard
               </RouterLink>
               <RouterLink to="/user/orders" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors" @click="mobileMenuOpen = false">
                 Orders
               </RouterLink>
               <button @click="logout" class="block w-full text-left px-4 py-2.5 text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors">
                 Logout
               </button>
             </template>
             <template v-else>
               <RouterLink to="/auth/login" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors" @click="mobileMenuOpen = false">
                 Login
               </RouterLink>
             </template>
           </div>
         </div>
      </Transition>
    </div>
  </nav>
</template>

<style scoped>
.badge-enter-active,
.badge-leave-active {
  transition: all 0.3s ease;
}
.badge-enter-from,
.badge-leave-to {
  opacity: 0;
  transform: scale(0);
}

.mobile-menu-enter-active,
.mobile-menu-leave-active {
  transition: all 0.3s ease;
}
.mobile-menu-enter-from,
.mobile-menu-leave-to {
  opacity: 0;
  max-height: 0;
}
.mobile-menu-enter-to,
.mobile-menu-leave-from {
  opacity: 1;
  max-height: 400px;
}
</style>
