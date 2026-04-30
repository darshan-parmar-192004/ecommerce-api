<script setup>
import { onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProductsStore } from '@/stores/products'
import { useCartStore } from '@/stores/cart'

const route = useRoute()
const router = useRouter()
const productsStore = useProductsStore()
const cartStore = useCartStore()

const quantity = ref(1)
const addedToCart = ref(false)
const currentImageIndex = ref(0)

onMounted(() => {
  productsStore.fetchProductById(route.params.id)
})

watch(() => route.params.id, (newId) => {
  productsStore.fetchProductById(newId)
  quantity.value = 1
  currentImageIndex.value = 0
})

const addToCart = () => {
  if (productsStore.currentProduct) {
    cartStore.addToCart(productsStore.currentProduct, quantity.value)
    addedToCart.value = true
    setTimeout(() => { addedToCart.value = false }, 2000)
  }
}

const buyNow = () => {
  addToCart()
  router.push({ name: 'Checkout' })
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <button
      @click="router.push('/')"
      class="inline-flex items-center gap-2 text-sm text-gray-600 hover:text-gray-900 mb-6 transition-colors"
    >
      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Products
    </button>

    <div v-if="productsStore.loading" class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="aspect-square bg-gray-200 rounded-xl animate-pulse" />
      <div class="space-y-4">
        <div class="h-8 bg-gray-200 rounded w-3/4 animate-pulse" />
        <div class="h-6 bg-gray-200 rounded w-1/4 animate-pulse" />
        <div class="h-24 bg-gray-200 rounded animate-pulse" />
      </div>
    </div>

    <div v-else-if="productsStore.currentProduct" class="grid grid-cols-1 md:grid-cols-2 gap-12">
      <!-- Image Gallery -->
      <div class="space-y-4">
        <div class="aspect-square rounded-xl overflow-hidden bg-gray-100">
          <Transition name="image" mode="out-in">
            <img
              :key="currentImageIndex"
              :src="productsStore.currentProduct.images?.[currentImageIndex] || '/placeholder.jpg'"
              :alt="productsStore.currentProduct.name"
              class="w-full h-full object-cover"
            />
          </Transition>
        </div>
        <!-- Thumbnail Navigation -->
        <div v-if="productsStore.currentProduct.images?.length > 1" class="flex gap-2">
          <button
            v-for="(img, idx) in productsStore.currentProduct.images"
            :key="idx"
            @click="currentImageIndex = idx"
            class="w-16 h-16 rounded-lg overflow-hidden border-2 transition-all duration-200"
            :class="currentImageIndex === idx ? 'border-gray-900' : 'border-transparent hover:border-gray-300'"
          >
            <img :src="img" class="w-full h-full object-cover" />
          </button>
        </div>
      </div>

      <!-- Product Info -->
      <div class="space-y-6">
        <div>
          <div class="flex items-center gap-2 mb-2">
            <span class="px-3 py-1 bg-gray-100 text-xs font-medium text-gray-700 rounded-full">
              {{ productsStore.currentProduct.category }}
            </span>
            <span
              v-if="productsStore.currentProduct.stock > 0"
              class="px-3 py-1 bg-green-100 text-green-700 text-xs font-medium rounded-full"
            >
              In Stock ({{ productsStore.currentProduct.stock }} available)
            </span>
            <span v-else class="px-3 py-1 bg-red-100 text-red-700 text-xs font-medium rounded-full">
              Out of Stock
            </span>
          </div>
          <h1 class="text-3xl font-bold text-gray-900">{{ productsStore.currentProduct.name }}</h1>
        </div>

        <div class="text-3xl font-bold text-gray-900">
          ₹{{ productsStore.currentProduct.price?.toFixed(2) }}
        </div>

        <p class="text-gray-700 leading-relaxed">
          {{ productsStore.currentProduct.description }}
        </p>

        <div class="flex items-center gap-4">
          <div class="flex items-center border border-gray-300 rounded-lg">
            <button
              @click="quantity = Math.max(1, quantity - 1)"
              class="px-3 py-2 hover:bg-gray-50 transition-colors"
            >
              -
            </button>
            <span class="px-4 py-2 border-x border-gray-300">{{ quantity }}</span>
            <button
              @click="quantity++"
              class="px-3 py-2 hover:bg-gray-50 transition-colors"
            >
              +
            </button>
          </div>
        </div>

        <div class="flex gap-4">
          <button
            @click="addToCart"
            :disabled="productsStore.currentProduct.stock === 0"
            :class="addedToCart ? 'bg-green-600 hover:bg-green-700 text-white' : 'btn-primary'"
            class="flex-1"
          >
            {{ addedToCart ? 'Added!' : 'Add to Cart' }}
          </button>
          <button
            @click="buyNow"
            :disabled="productsStore.currentProduct.stock === 0"
            class="btn-secondary flex-1"
          >
            Buy Now
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-enter-active,
.image-leave-active {
  transition: opacity 0.3s ease;
}
.image-enter-from,
.image-leave-to {
  opacity: 0;
}
</style>
