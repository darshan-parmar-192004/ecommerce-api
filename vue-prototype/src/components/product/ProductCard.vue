<script setup>
import { useCartStore } from '@/stores/cart'
import { useRouter } from 'vue-router'

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()
const router = useRouter()

const addToCart = (e) => {
  e.stopPropagation()
  cartStore.addToCart(props.product)
}

const goToDetail = () => {
  router.push({ name: 'ProductDetail', params: { id: props.product.product_id } })
}

const handleKeydown = (e) => {
  if (e.key === 'Enter' || e.key === ' ') {
    e.preventDefault()
    goToDetail()
  }
}
</script>

<template>
  <div
    @click="goToDetail"
    @keydown="handleKeydown"
    role="button"
    tabindex="0"
    :aria-label="`View details for ${product.name}`"
    class="group relative cursor-pointer focus:outline-none focus:ring-2 focus:ring-brand-900 focus:ring-offset-2 rounded-2xl"
  >
    <!-- Premium Card Container -->
    <div class="relative bg-gradient-to-br from-white/95 to-white/85 dark:from-brand-800/95 dark:to-brand-700/85 backdrop-blur-xl rounded-2xl border border-white/50 dark:border-brand-600/50 shadow-lg shadow-gray-200/50 dark:shadow-brand-900/50 overflow-hidden transition-all duration-500 hover:shadow-2xl hover:shadow-gray-900/10 dark:hover:shadow-brand-900/20 hover:scale-[1.02] hover:border-gray-200/60 dark:hover:border-brand-600/60">
      <!-- Image Container -->
      <div class="relative aspect-square overflow-hidden bg-gradient-to-br from-gray-50 to-gray-100/50 dark:from-brand-700 dark:to-brand-800/50">
        <img
          :src="product.images?.[0] || '/placeholder.jpg'"
          :alt="product.name"
          loading="lazy"
          class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110"
        />

        <!-- Gradient overlay on hover -->
        <div class="absolute inset-0 bg-gradient-to-t from-black/20 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>

        <!-- Category Badge - Premium -->
        <div class="absolute top-4 left-4">
          <span class="px-3 py-1.5 bg-white/95 dark:bg-brand-800/95 backdrop-blur-xl text-xs font-bold text-gray-900 dark:text-gray-100 rounded-xl border border-white/50 dark:border-brand-600/50 shadow-md">
            {{ product.category_name || 'Uncategorized' }}
          </span>
        </div>

        <!-- Add to Cart Button - Premium -->
        <button
          @click="addToCart"
          class="absolute bottom-4 right-4 bg-gradient-to-br from-gray-900 to-gray-800 dark:from-brand-700 dark:to-brand-800 text-white p-3.5 rounded-xl shadow-xl opacity-0 translate-y-3 hover:from-gray-800 hover:to-gray-700 dark:hover:from-brand-600 dark:hover:to-brand-700 transition-all duration-300 group-hover:opacity-100 group-hover:translate-y-0 hover:scale-110 active:scale-95 border border-gray-700/30 dark:border-brand-600/30"
          :aria-label="`Add ${product.name} to cart`"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4" />
          </svg>
        </button>
      </div>

      <!-- Content -->
      <div class="p-5 space-y-3">
        <div class="flex items-start justify-between gap-2">
          <h3 class="text-base font-bold text-gray-900 dark:text-gray-100 line-clamp-2 leading-tight group-hover:text-gray-700 dark:group-hover:text-gray-300 transition-colors duration-300">
            {{ product.name }}
          </h3>
        </div>

        <!-- Rating Stars -->
        <div v-if="product.rating" class="flex items-center gap-1.5">
          <div class="flex items-center">
            <svg class="w-4 h-4 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
            <span class="text-sm font-semibold text-gray-800 dark:text-gray-300 ml-1">{{ product.rating }}</span>
          </div>
          <span class="text-xs text-gray-400 dark:text-gray-500">•</span>
          <span class="text-xs text-gray-500 dark:text-gray-400">Verified</span>
        </div>

        <!-- Price & CTA Row -->
        <div class="flex items-center justify-between pt-2 border-t border-gray-100/80 dark:border-brand-700/80">
          <div class="flex items-baseline gap-1">
            <span class="text-2xl font-bold bg-gradient-to-r from-gray-900 to-gray-700 dark:from-gray-100 dark:to-gray-300 bg-clip-text text-transparent">
              ₹{{ product.price?.toFixed(2) }}
            </span>
          </div>

          <!-- Quick view indicator -->
          <div class="flex items-center gap-1.5 text-xs font-medium text-gray-600 dark:text-gray-400 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
            <span>View details</span>
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 5l7 7-7 7"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Premium hover glow effect -->
    <div class="absolute -inset-0.5 bg-gradient-to-r from-gray-900/20 via-gray-700/20 to-gray-600/20 dark:from-brand-600/20 dark:via-brand-500/20 dark:to-brand-400/20 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-500 -z-10 blur"></div>
  </div>
</template>