<script setup>
const route = useRoute()
const cartStore = useCartStore()

const { data: product, pending, error } = await useFetch(`/api/products/${route.params.id}`)

useSeoMeta({
  title: () => product.value ? `${product.value.name} - E-Commerce Store` : 'Product',
  description: () => product.value?.description || 'Product details'
})

const added = ref(false)

const addToCart = () => {
  if (product.value) {
    cartStore.addItem(product.value)
    added.value = true
    setTimeout(() => {
      added.value = false
    }, 2000)
  }
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <NuxtLink to="/products" class="inline-flex items-center text-sm text-gray-500 hover:text-indigo-600 mb-6 transition-colors">
      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Products
    </NuxtLink>

    <div v-if="pending" class="animate-pulse">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="bg-gray-200 aspect-square rounded-lg" />
        <div>
          <div class="h-8 bg-gray-200 rounded w-3/4 mb-4" />
          <div class="h-6 bg-gray-200 rounded w-1/4 mb-6" />
          <div class="h-4 bg-gray-200 rounded w-full mb-2" />
          <div class="h-4 bg-gray-200 rounded w-2/3" />
        </div>
      </div>
    </div>

    <div v-else-if="error" class="bg-white rounded-lg shadow-sm p-8 text-center">
      <svg class="mx-auto h-16 w-16 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <h1 class="mt-4 text-2xl font-bold text-gray-900 mb-2">Product Not Found</h1>
      <p class="text-gray-500 mb-6">The product you're looking for doesn't exist.</p>
      <NuxtLink to="/products" class="btn-primary">
        Browse Products
      </NuxtLink>
    </div>

    <div v-else-if="product" class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="bg-gray-100 aspect-square rounded-lg flex items-center justify-center">
        <svg class="w-32 h-32 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
        </svg>
      </div>

      <div>
        <h1 class="text-3xl font-bold text-gray-900 mb-2">
          {{ product.name }}
        </h1>
        
        <p class="text-3xl font-semibold text-indigo-600 mb-6">
          ${{ Number(product.price).toFixed(2) }}
        </p>

        <div v-if="product.description" class="prose prose-gray mb-6">
          <p class="text-gray-600">{{ product.description }}</p>
        </div>

        <div class="flex gap-4">
          <button 
            @click="addToCart"
            class="flex-1 btn-primary flex items-center justify-center gap-2"
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
          
          <NuxtLink to="/products" class="btn-secondary">
            Back to Products
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
