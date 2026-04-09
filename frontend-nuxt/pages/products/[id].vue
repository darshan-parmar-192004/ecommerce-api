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
    setTimeout(() => {
      added.value = false
    }, 2000)
  }
}
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
      <div class="bg-surface-container aspect-square rounded-lg flex items-center justify-center">
        <svg class="w-32 h-32 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
        </svg>
      </div>

      <div>
        <h1 class="text-3xl font-bold text-on_surface font-display mb-2">
          {{ product.name }}
        </h1>

        <p class="text-3xl font-semibold text-primary mb-6">
          ₹ {{ Number(product.price).toFixed(2) }}
        </p>

        <div v-if="product.description" class="prose prose-gray mb-6">
          <p class="text-on_surface_variant">{{ product.description }}</p>
        </div>

        <!-- Quantity Selector -->
        <div class="mb-6">
          <label class="block text-sm font-semibold text-on_surface_variant mb-3">Quantity</label>
          <div class="inline-flex items-center border border-outline-variant/30 rounded-lg bg-surface-container-low overflow-hidden">
            <button
              @click="decrementQuantity"
              class="px-4 py-3 text-on_surface_variant hover:bg-surface-container hover:text-primary transition-colors"
              :disabled="quantity <= 1"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
              </svg>
            </button>
            <span class="px-6 py-3 font-semibold min-w-[3rem] text-center text-on_surface">{{ quantity }}</span>
            <button
              @click="incrementQuantity"
              class="px-4 py-3 text-on_surface_variant hover:bg-surface-container hover:text-primary transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </button>
          </div>
        </div>

        <div class="flex gap-4">
          <button
            @click="addToCart"
            class="flex-1 py-4 px-6 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg font-semibold flex items-center justify-center gap-2 hover:opacity-90 transition-all shadow-lg shadow-primary/25"
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

          <NuxtLink to="/products" class="py-4 px-6 bg-surface-container text-on_surface_variant rounded-lg font-medium hover:bg-surface-container-high transition-colors">
            Back to Products
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
