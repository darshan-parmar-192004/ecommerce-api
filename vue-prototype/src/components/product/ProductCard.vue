<script setup>
import { useCart } from '@/composables/useCart'
import { RouterLink } from 'vue-router'
import { Plus, Star, ChevronRight } from 'lucide-vue-next'

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const { addToCart: cartAdd } = useCart()

const addToCart = (e) => {
  e.preventDefault()
  cartAdd(props.product)
}
</script>

<template>
  <RouterLink
    :to="{ name: 'ProductDetail', params: { id: product.productId } }"
    :aria-label="`View details for ${product.name}`"
    class="group relative cursor-pointer focus:outline-none focus:ring-2 focus:ring-brand-900 focus:ring-offset-2 rounded-2xl inline-block"
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
            {{ product.categoryName || 'Uncategorized' }}
          </span>
        </div>

        <!-- Add to Cart Button - Premium -->
        <button
          @click.prevent="addToCart"
          class="absolute bottom-4 right-4 bg-gradient-to-br from-gray-900 to-gray-800 dark:from-brand-700 dark:to-brand-800 text-white p-3.5 rounded-xl shadow-xl opacity-0 translate-y-3 hover:from-gray-800 hover:to-gray-700 dark:hover:from-brand-600 dark:hover:to-brand-700 transition-all duration-300 group-hover:opacity-100 group-hover:translate-y-0 hover:scale-110 active:scale-95 border border-gray-700/30 dark:border-brand-600/30"
          :aria-label="`Add ${product.name} to cart`"
        >
          <Plus class="w-5 h-5" />
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
            <Star class="w-4 h-4 text-yellow-400" fill="currentColor" />
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
            <ChevronRight class="w-3.5 h-3.5" />
          </div>
        </div>
      </div>
    </div>

    <!-- Premium hover glow effect -->
    <div class="absolute -inset-0.5 bg-gradient-to-r from-gray-900/20 via-gray-700/20 to-gray-600/20 dark:from-brand-600/20 dark:via-brand-500/20 dark:to-brand-400/20 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-500 -z-10 blur"></div>
  </RouterLink>
</template>