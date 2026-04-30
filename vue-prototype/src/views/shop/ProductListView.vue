<script setup>
import { onMounted } from 'vue'
import { useProductsStore } from '@/stores/products'
import { useRouter } from 'vue-router'
import ProductCard from '@/components/product/ProductCard.vue'

const productsStore = useProductsStore()
const router = useRouter()

onMounted(async () => {
  await productsStore.fetchProducts()
})

const handleProductClick = (productId) => {
  router.push({ name: 'ProductDetail', params: { id: productId } })
}
</script>

<template>
  <div class="product-list-page">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Products</h1>

    <div v-if="productsStore.loading" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <div v-for="n in 12" :key="n" class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden animate-pulse">
        <div class="bg-gray-200 w-full h-48"></div>
        <div class="p-4 space-y-3">
          <div class="h-4 bg-gray-200 rounded w-1/3"></div>
          <div class="h-6 bg-gray-200 rounded w-3/4"></div>
          <div class="h-8 bg-gray-200 rounded w-1/4"></div>
        </div>
      </div>
    </div>

    <div v-else-if="productsStore.error" class="bg-red-50 border border-red-200 text-red-700 px-5 py-4 rounded-xl text-center">
      {{ productsStore.error }}
    </div>

    <div v-else-if="productsStore.products.length === 0" class="text-center py-16">
      <svg class="w-20 h-20 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
      </svg>
      <p class="text-xl font-medium text-gray-500 mb-2">No products found</p>
      <p class="text-gray-400">Try adjusting your filters or search terms</p>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="product in productsStore.products"
        :key="product.product_id"
        :product="product"
      />
    </div>

    <div v-if="productsStore.pagination.total > productsStore.products.length" class="flex justify-center mt-10">
      <button
        @click="productsStore.pagination.page++; productsStore.fetchProducts()"
        class="px-6 py-3 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 transition-colors"
      >
        Load More
      </button>
    </div>
  </div>
</template>

<style scoped>
.product-list-page {
  padding: 1rem;
  max-width: 1400px;
  margin: 0 auto;
}
</style>
