<script setup>
import { useCartStore } from '~/stores/cart'

const route = useRoute()
const cartStore = useCartStore()

const { data: product, pending, error } = await useFetch(`/api/products/${route.params.id}`)

useSeoMeta({
  title: () => product.value ? `${product.value.name} - E-Commerce Store` : 'Product',
  description: () => product.value?.description || 'Product details'
})

const quantity = ref(1)
const added = ref(false)
const isZoomed = ref(false)
const mousePosition = ref({ x: 0, y: 0 })

const reviews = ref([
  { id: 1, name: 'John D.', rating: 5, comment: 'Excellent product! Very satisfied with the quality.', date: '2024-01-15' },
  { id: 2, name: 'Sarah M.', rating: 4, comment: 'Good value for money. Fast delivery.', date: '2024-01-10' },
  { id: 3, name: 'Alex K.', rating: 5, comment: 'Highly recommend! Exceeded expectations.', date: '2024-01-05' }
])

const incrementQuantity = () => {
  quantity.value++
}

const decrementQuantity = () => {
  if (quantity.value > 1) {
    quantity.value--
  }
}

const addToCart = () => {
  if (product.value) {
    for (let i = 0; i < quantity.value; i++) {
      cartStore.addItem(product.value)
    }
    added.value = true
    cartStore.openCart()
    setTimeout(() => {
      added.value = false
    }, 2000)
  }
}

const handleMouseMove = (e) => {
  if (!isZoomed.value) return
  const rect = e.target.getBoundingClientRect()
  mousePosition.value = {
    x: ((e.clientX - rect.left) / rect.width) * 100,
    y: ((e.clientY - rect.top) / rect.height) * 100
  }
}

const toggleWishlist = () => {
  console.log('Wishlist toggle - implement wishlist store')
}

const averageRating = computed(() => {
  if (!reviews.value.length) return 0
  return reviews.value.reduce((sum, r) => sum + r.rating, 0) / reviews.value.length
})
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface">
    <NuxtLink to="/products" class="inline-flex items-center text-sm text-on_surface_variant hover:text-primary mb-6 transition-colors">
      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Products
    </NuxtLink>

    <div v-if="pending" class="animate-pulse">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="bg-surface-container aspect-square rounded-lg" />
        <div>
          <div class="h-8 bg-surface-container rounded w-3/4 mb-4" />
          <div class="h-6 bg-surface-container rounded w-1/4 mb-6" />
          <div class="h-4 bg-surface-container rounded w-full mb-2" />
          <div class="h-4 bg-surface-container rounded w-2/3" />
        </div>
      </div>
    </div>

    <div v-else-if="error" class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center">
      <svg class="mx-auto h-16 w-16 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <h1 class="mt-4 text-2xl font-bold text-on_surface mb-2">Product Not Found</h1>
      <p class="text-on_surface_variant mb-6">The product you're looking for doesn't exist.</p>
      <NuxtLink to="/products" class="inline-block px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
        Browse Products
      </NuxtLink>
    </div>

    <div v-else-if="product" class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="relative group">
        <div 
          class="bg-surface-container aspect-square rounded-lg flex items-center justify-center overflow-hidden cursor-zoom-in"
          @mouseenter="isZoomed = true"
          @mouseleave="isZoomed = false"
          @mousemove="handleMouseMove"
        >
          <div 
            class="w-full h-full flex items-center justify-center transition-transform duration-200"
            :style="{
              transform: isZoomed ? 'scale(1.5)' : 'scale(1)',
              transformOrigin: `${mousePosition.value.x}% ${mousePosition.value.y}%`
            }"
          >
            <svg class="w-32 h-32 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          
          <div v-if="isZoomed" class="absolute inset-0 pointer-events-none border-2 border-primary/50 rounded-lg"></div>
        </div>
        
        <button 
          @click="toggleWishlist"
          class="absolute top-4 right-4 p-3 bg-surface/80 backdrop-blur-sm rounded-full shadow-lg hover:scale-110 transition-transform"
        >
          <svg class="w-6 h-6 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
          </svg>
        </button>
        
        <div v-if="product.stock === 0" class="absolute top-4 left-4 bg-error text-on_error text-xs font-bold uppercase tracking-wider px-3 py-1.5 rounded">
          Out of Stock
        </div>
      </div>

      <div>
        <h1 class="text-3xl font-bold text-on_surface font-display mb-2">
          {{ product.name }}
        </h1>
        
        <div class="flex items-center gap-4 mb-4">
          <div class="flex items-center gap-1">
            <svg v-for="i in 5" :key="i" class="w-5 h-5" :class="i <= Math.round(averageRating) ? 'text-yellow-400' : 'text-outline'" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
            <span class="ml-2 text-sm text-on_surface_variant">{{ averageRating.toFixed(1) }} ({{ reviews.length }} reviews)</span>
          </div>
        </div>

        <p class="text-3xl font-semibold text-primary mb-4">
          ₹ {{ Number(product.price).toFixed(2) }}
        </p>

        <div v-if="product.description" class="prose prose-gray mb-6">
          <p class="text-on_surface_variant">{{ product.description }}</p>
        </div>

        <div class="mb-6">
          <p class="text-sm text-on_surface_variant mb-2">
            <span v-if="product.stock > 10" class="text-green-600 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
              </svg>
              In Stock ({{ product.stock }} available)
            </span>
            <span v-else-if="product.stock > 0" class="text-yellow-600 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l.278.919a.564.564 0 00.475.345l1.622.23c1.277.18 1.762 1.715.782 2.474l-.665.514a.563.563 0 00.182.529l1.409 1.16c.708.582.564 1.653-.232 1.724l-1.275.116c-.938.085-1.664.916-1.173 1.738l.317.529a.564.564 0 01-.048.631l-.665.514c-.979.757-.004 2.294.987 2.294h1.734c.99 0 1.966-1.537.988-2.294l-.317-.529a.564.564 0 01.048-.631l.317-.529c.491-.822-.235-1.653-1.173-1.738l-1.275-.116a.563.563 0 00-.182.529l.665.514c.796.614.94 2.144.232 1.724l-1.409 1.16a.563.563 0 00-.182.529l.665.514c.98.757 1.495 2.294.782 2.474l-1.622.23a.564.564 0 00-.475.345l-.278.919c-.765 1.36-2.722 1.36-3.486 0l-.278-.919a.564.564 0 00-.475-.345l-1.622-.23c-1.277-.18-1.762-1.715-.782-2.474l.665-.514a.563.563 0 00-.182-.529l-1.409-1.16c-.708-.582-.564-1.653.232-1.724l1.275-.116c.938-.085 1.664-.916 1.173-1.738l-.317-.529a.564.564 0 01.048-.631l.665-.514c.979-.757.004-2.294-.987-2.294H8.257c-.99 0-1.966 1.537-.988 2.294l.317.529a.564.564 0 01.048.631l-.317.529c-.491.822.235 1.653 1.173 1.738l1.275.116c.938.085 1.664.916 1.173 1.738l-.665.514a.563.563 0 00.182.529l1.409 1.16c.708.582.564 1.653-.232 1.724l-1.275.116c-.938.085-1.664.916-1.173 1.738l.317.529c.283.473.283 1.056 0 1.529z" clip-rule="evenodd" />
              </svg>
              Only {{ product.stock }} left
            </span>
          </p>
        </div>

        <div class="mb-6">
          <label class="block text-sm font-semibold text-on_surface_variant mb-3">Quantity</label>
          <div class="inline-flex items-center border border-outline-variant/30 rounded-lg bg-surface-container-low overflow-hidden">
            <button
              @click="decrementQuantity"
              class="px-4 py-3 text-on_surface_variant hover:bg-surface-container hover:text-primary transition-all duration-200 hover:scale-110 active:scale-95"
              :disabled="quantity <= 1"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
              </svg>
            </button>
            <span class="px-6 py-3 font-semibold min-w-[3rem] text-center text-on_surface">{{ quantity }}</span>
            <button
              @click="incrementQuantity"
              class="px-4 py-3 text-on_surface_variant hover:bg-surface-container hover:text-primary transition-all duration-200 hover:scale-110 active:scale-95"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </button>
          </div>
        </div>

        <div class="flex gap-4 mb-8">
          <button
            @click="addToCart"
            class="flex-1 py-4 px-6 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg font-semibold flex items-center justify-center gap-2 hover:opacity-90 transition-all shadow-lg shadow-primary/25 hover:scale-[1.02] active:scale-[0.98]"
            :class="{ 'bg-green-600 hover:bg-green-700': added }"
          >
            <svg v-if="!added" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            {{ added ? 'Added to Cart!' : 'Add to Cart' }}
          </button>
        </div>

        <!-- Reviews Section -->
        <div class="border-t border-outline-variant/20 pt-8">
          <h3 class="text-xl font-bold text-on_surface font-display mb-6">Customer Reviews</h3>
          
          <div class="space-y-4">
            <div 
              v-for="review in reviews" 
              :key="review.id"
              class="bg-surface-container-low p-4 rounded-lg"
            >
              <div class="flex items-center justify-between mb-2">
                <div class="flex items-center gap-2">
                  <div class="w-8 h-8 rounded-full bg-gradient-to-r from-primary to-primary-container flex items-center justify-center text-white text-sm font-semibold">
                    {{ review.name.charAt(0) }}
                  </div>
                  <span class="font-semibold text-on_surface">{{ review.name }}</span>
                </div>
                <div class="flex items-center gap-1">
                  <svg v-for="i in 5" :key="i" class="w-4 h-4" :class="i <= review.rating ? 'text-yellow-400' : 'text-outline'" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                </div>
              </div>
              <p class="text-on_surface_variant">{{ review.comment }}</p>
              <p class="text-xs text-outline mt-2">{{ review.date }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
