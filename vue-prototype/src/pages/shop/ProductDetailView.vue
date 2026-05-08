<script setup>
import { onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProducts } from '@/composables/useProducts'
import { useCart } from '@/composables/useCart'
import { ArrowLeft } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const productsStore = useProducts()
const { addToCart: cartAdd } = useCart()

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
    cartAdd(productsStore.currentProduct, quantity.value)
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
      class="inline-flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 mb-6 transition-colors"
    >
      <ArrowLeft class="w-4 h-4" />
      Back to Products
    </button>

    <div v-if="productsStore.loading" class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="aspect-square bg-gray-200 dark:bg-brand-700 rounded-xl animate-pulse" />
      <div class="space-y-4">
        <div class="h-8 bg-gray-200 dark:bg-brand-700 rounded w-3/4 animate-pulse" />
        <div class="h-6 bg-gray-200 dark:bg-brand-700 rounded w-1/4 animate-pulse" />
        <div class="h-24 bg-gray-200 dark:bg-brand-700 rounded animate-pulse" />
      </div>
    </div>

    <div v-else-if="productsStore.currentProduct" class="grid grid-cols-1 md:grid-cols-2 gap-12">
      <!-- Image Gallery -->
      <div class="space-y-4">
        <div class="aspect-square rounded-xl overflow-hidden bg-gray-100 dark:bg-brand-700">
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
            :class="currentImageIndex === idx ? 'border-gray-900 dark:border-gray-100' : 'border-transparent hover:border-gray-300 dark:hover:border-brand-600'"
          >
            <img :src="img" class="w-full h-full object-cover" />
          </button>
        </div>
      </div>

      <!-- Product Info -->
      <div class="space-y-6">
        <div>
          <div class="flex items-center gap-2 mb-2">
            <span class="px-3 py-1 bg-gray-100 dark:bg-brand-700 text-xs font-medium text-gray-700 dark:text-gray-300 rounded-full">
              {{ productsStore.currentProduct.category }}
            </span>
            <span
              v-if="productsStore.currentProduct.stock > 0"
              class="px-3 py-1 bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 text-xs font-medium rounded-full"
            >
              In Stock ({{ productsStore.currentProduct.stock }} available)
            </span>
            <span v-else class="px-3 py-1 bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-400 text-xs font-medium rounded-full">
              Out of Stock
            </span>
          </div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">{{ productsStore.currentProduct.name }}</h1>
        </div>

        <div class="text-3xl font-bold text-gray-900 dark:text-gray-100">
          ₹{{ productsStore.currentProduct.price?.toFixed(2) }}
        </div>

        <p class="text-gray-700 dark:text-gray-300 leading-relaxed">
          {{ productsStore.currentProduct.description }}
        </p>

        <div class="flex items-center gap-4">
          <div class="flex items-center border border-gray-300 dark:border-brand-600 rounded-lg">
            <button
              @click="quantity = Math.max(1, quantity - 1)"
              class="px-3 py-2 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors"
              aria-label="Decrease quantity"
            >
              -
            </button>
            <span class="px-4 py-2 border-x border-gray-300 dark:border-brand-600 text-gray-900 dark:text-gray-100" aria-live="polite" aria-atomic="true">{{ quantity }}</span>
            <button
              @click="quantity++"
              class="px-3 py-2 hover:bg-gray-50 dark:hover:bg-brand-700 transition-colors"
              aria-label="Increase quantity"
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
