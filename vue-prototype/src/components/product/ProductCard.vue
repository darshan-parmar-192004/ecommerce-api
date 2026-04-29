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
  router.push({ name: 'ProductDetail', params: { id: props.product.id } })
}
</script>

<template>
  <div
    @click="goToDetail"
    class="group card cursor-pointer"
  >
    <!-- Image Container -->
    <div class="relative aspect-square overflow-hidden bg-gray-100">
      <img
        :src="product.images?.[0] || '/placeholder.jpg'"
        :alt="product.name"
        class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110"
      />
      
      <!-- Add to Cart Button - appears on hover -->
      <button
        @click="addToCart"
        class="absolute bottom-4 right-4 bg-white text-gray-900 p-3 rounded-full shadow-lg opacity-0 translate-y-2 hover:bg-gray-900 hover:text-white transition-all duration-300 group-hover:opacity-100 group-hover:translate-y-0"
        :aria-label="`Add ${product.name} to cart`"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
      </button>

      <!-- Category Badge -->
      <div class="absolute top-4 left-4">
        <span class="px-3 py-1 bg-white/90 backdrop-blur-sm text-xs font-medium text-gray-900 rounded-full">
          {{ product.category }}
        </span>
      </div>
    </div>

    <!-- Content -->
    <div class="p-5">
      <div class="flex items-start justify-between gap-2">
        <h3 class="text-sm font-medium text-gray-900 line-clamp-2 leading-snug">{{ product.name }}</h3>
      </div>
      
      <div class="flex items-center justify-between mt-3">
        <span class="text-xl font-bold text-gray-900">${{ product.price?.toFixed(2) }}</span>
        
        <div class="flex items-center gap-1">
          <svg class="w-4 h-4 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
          </svg>
          <span class="text-sm text-gray-600">{{ product.rating || '0' }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
