<script setup>
import { ref, watch, onMounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'

const authStore = useAuthStore()
const cartStore = useCartStore()
const router = useRouter()
const mobileMenuOpen = ref(false)
const isDark = ref(false)

// Close mobile menu on route change
watch(() => router.currentRoute.value, () => {
  mobileMenuOpen.value = false
})

// Initialize dark mode from localStorage or system preference
onMounted(() => {
   const savedTheme = localStorage.getItem('theme')
   if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
     isDark.value = true
     document.documentElement.classList.add('dark')
   }
 })

const toggleDarkMode = () => {
   isDark.value = !isDark.value
   if (isDark.value) {
     document.documentElement.classList.add('dark')
     localStorage.setItem('theme', 'dark')
   } else {
     document.documentElement.classList.remove('dark')
     localStorage.setItem('theme', 'light')
   }
 }

const logout = async () => {
  await authStore.logout()
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
               v-if="authStore.isAdmin"
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
          <!-- Dark Mode Toggle -->
          <button
            @click="toggleDarkMode"
            class="p-2 text-gray-700 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-100 transition-colors"
            :aria-label="isDark.value ? 'Switch to light mode' : 'Switch to dark mode'"
          >
            <svg v-if="!isDark" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
          </button>

          <!-- Cart -->
           <button
             @click="cartStore.toggleDrawer()"
             class="relative p-2 text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-all duration-200 hover:scale-110"
           >
             <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
             </svg>
             <Transition name="badge">
               <span
                 v-if="cartStore.cartCount > 0"
                 class="absolute -top-1 -right-1 bg-gray-900 dark:bg-brand-700 text-white dark:text-gray-100 text-xs w-5 h-5 rounded-full flex items-center justify-center animate-bounce"
               >
                 {{ cartStore.cartCount }}
               </span>
             </Transition>
           </button>

          <!-- Auth Section -->
           <template v-if="authStore.isAuthenticated">
             <div class="relative group">
               <button
                 class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors px-3 py-2 rounded-lg hover:bg-gray-50 dark:hover:bg-brand-700"
               >
                 <div class="w-8 h-8 bg-gray-900 dark:bg-brand-700 text-white dark:text-gray-100 rounded-full flex items-center justify-center text-xs font-bold">
                   {{ (authStore.user?.name || 'U').charAt(0).toUpperCase() }}
                 </div>
                 <span class="hidden sm:inline">{{ authStore.user?.name?.split(' ')[0] || 'Account' }}</span>
                 <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                 </svg>
               </button>
               <!-- Dropdown menu - works on hover for desktop, click for mobile via group -->
               <div class="absolute right-0 mt-2 w-48 bg-white dark:bg-brand-800 rounded-xl shadow-xl border border-gray-100 dark:border-brand-700 py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
                 <RouterLink v-if="authStore.isAdmin" to="/admin" class="block px-4 py-2.5 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors rounded-lg mx-2">
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
            <Transition name="icon" mode="out-in">
              <svg v-if="!mobileMenuOpen" key="open" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
              <svg v-else key="close" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </Transition>
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
             <template v-if="authStore.isAuthenticated">
               <RouterLink v-if="authStore.isAdmin" to="/admin" class="block px-4 py-2.5 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-brand-700 rounded-lg transition-colors" @click="mobileMenuOpen = false">
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
.icon-enter-active,
.icon-leave-active {
  transition: all 0.2s ease;
}
.icon-enter-from,
.icon-leave-to {
  opacity: 0;
  transform: rotate(90deg);
}

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
