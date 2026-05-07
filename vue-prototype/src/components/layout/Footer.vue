<script setup>
import { X as TwitterX, Github } from 'lucide-vue-next'
</script>

<template>
  <footer class="bg-white dark:bg-brand-800 border-t border-gray-200 dark:border-brand-700 mt-auto">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8 mb-8">
        <!-- Brand -->
        <div class="md:col-span-1">
          <h3 class="text-lg font-bold text-gray-900 dark:text-gray-100 mb-4">E-Commerce</h3>
          <p class="text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
            Modern shopping experience with premium products and seamless checkout.
          </p>
        </div>

        <!-- Quick Links -->
        <div>
          <h4 class="text-sm font-semibold text-gray-900 dark:text-gray-100 mb-4">Quick Links</h4>
          <ul class="space-y-2">
            <li>
              <RouterLink to="/products" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Products
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/cart" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Cart
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/user/orders" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Orders
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/user/profile" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Profile
              </RouterLink>
            </li>
          </ul>
        </div>

        <!-- Account -->
        <div>
          <h4 class="text-sm font-semibold text-gray-900 dark:text-gray-100 mb-4">Account</h4>
          <ul class="space-y-2">
            <li>
              <RouterLink to="/auth/login" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Login
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/auth/register" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
                Register
              </RouterLink>
            </li>
          </ul>
        </div>

        <!-- Newsletter -->
        <div>
          <h4 class="text-sm font-semibold text-gray-900 dark:text-gray-100 mb-4">Stay Updated</h4>
          <p class="text-sm text-gray-600 dark:text-gray-400 mb-3">Subscribe to our newsletter for updates.</p>
          <div class="flex gap-2">
            <input
              type="email"
              placeholder="your@email.com"
              class="flex-1 px-3 py-2 text-sm border border-gray-300 dark:border-brand-600 dark:bg-brand-700 dark:text-gray-100 rounded-lg focus:outline-none focus:ring-2 focus:ring-brand-900 focus:border-transparent transition-all duration-200"
            />
            <button class="px-4 py-2 text-sm font-medium text-white bg-brand-900 dark:bg-brand-700 rounded-lg hover:bg-brand-800 dark:hover:bg-brand-600 transition-all duration-200">
              Join
            </button>
          </div>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="pt-8 border-t border-gray-200 dark:border-brand-700 flex flex-col md:flex-row items-center justify-between gap-4">
        <p class="text-sm text-gray-500 dark:text-gray-400">
          &copy; {{ new Date().getFullYear() }} E-Commerce. All rights reserved.
        </p>
        <div class="flex items-center gap-6">
          <a href="#" class="text-gray-400 hover:text-gray-600 dark:text-gray-500 dark:hover:text-gray-300 transition-colors">
            <span class="sr-only">Twitter</span>
            <TwitterX class="w-5 h-5" />
          </a>
          <a href="#" class="text-gray-400 hover:text-gray-600 dark:text-gray-500 dark:hover:text-gray-300 transition-colors">
            <span class="sr-only">GitHub</span>
            <Github class="w-5 h-5" />
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>
