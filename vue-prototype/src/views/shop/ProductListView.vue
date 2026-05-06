<script setup>
import { ref, computed, onMounted } from 'vue'
import { useProductsStore } from '@/stores/products'
import { useRouter } from 'vue-router'
import ProductCard from '@/components/product/ProductCard.vue'
import FilterCard from '@/components/ui/FilterCard.vue'
import PaginationControls from '@/components/ui/PaginationControls.vue'
import SkeletonCard from '@/components/ui/SkeletonCard.vue'

const productsStore = useProductsStore()
const router = useRouter()

// Filter state
const filters = ref({
  category: '',
  minPrice: '',
  maxPrice: '',
  search: ''
})

const appliedFilterCount = computed(() => {
  return Object.values(filters.value).filter(v => v && v.toString().trim()).length
})

// Get categories from store
const categories = computed(() => productsStore.categories)

const applyFilters = async () => {
  await productsStore.setFilters({
    ...filters.value,
    minPrice: filters.value.minPrice || null,
    maxPrice: filters.value.maxPrice || null,
    search: filters.value.search || null
  })
}

const clearFilters = () => {
  filters.value = {
    category: '',
    minPrice: '',
    maxPrice: '',
    search: ''
  }
  productsStore.resetFilters()
}

onMounted(async () => {
  await productsStore.fetchCategories()
  await productsStore.fetchProducts()
})

const handleProductClick = (productId) => {
  router.push({ name: 'ProductDetail', params: { id: productId } })
}

const handlePageChange = (page) => {
  productsStore.setPage(page)
}
</script>

<template>
  <div class="product-list-page min-h-screen bg-gradient-to-br from-gray-50/80 via-white to-gray-50/80 dark:from-brand-900/80 dark:via-brand-800 dark:to-brand-900/80">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold bg-gradient-to-r from-gray-900 via-gray-800 to-gray-900 dark:from-gray-100 dark:via-gray-200 dark:to-gray-100 bg-clip-text text-transparent mb-2">
          Our Products
        </h1>
        <p class="text-gray-600 dark:text-gray-400 text-lg">
          Discover amazing items from our curated collection
        </p>
      </div>

      <div class="flex flex-col lg:flex-row gap-8">
        <!-- Filter Card - Sticky on desktop -->
        <div class="lg:w-80 flex-shrink-0">
          <div class="sticky top-8">
            <FilterCard
              :search-value="filters.search"
              :category-value="filters.category"
              :min-price-value="filters.minPrice"
              :max-price-value="filters.maxPrice"
              :categories="categories"
              @update:search="filters.search = $event"
              @update:category="filters.category = $event"
              @update:min-price="filters.minPrice = $event"
              @update:max-price="filters.maxPrice = $event"
              @apply="applyFilters"
              @clear="clearFilters"
            />
          </div>
        </div>

        <!-- Main Content Area -->
        <div class="flex-1 min-w-0">
          <!-- Results Summary -->
          <div class="flex items-center justify-between mb-6 pb-4 border-b border-gray-200/60 dark:border-brand-700/60">
            <p class="text-gray-600 dark:text-gray-400">
              <span class="font-semibold text-gray-900 dark:text-gray-100">{{ productsStore.pagination.total || productsStore.products.length }}</span>
              products found
            </p>
            <div class="flex items-center gap-4">
              <!-- Results count indicator -->
              <div class="flex items-center gap-2">
                <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
                <span class="text-sm text-gray-500 dark:text-gray-400">Live inventory</span>
              </div>
            </div>
          </div>

          <!-- Loading State -->
          <div v-if="productsStore.loading" class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
            <SkeletonCard v-for="n in 9" :key="n" />
          </div>

          <!-- Error State -->
          <div v-else-if="productsStore.error" class="bg-gradient-to-r from-red-50 to-red-100/50 dark:from-red-900/20 dark:to-red-800/20 border border-red-200/50 dark:border-red-700/50 text-red-700 dark:text-red-400 px-6 py-5 rounded-2xl text-center backdrop-blur-sm">
            <div class="flex items-center justify-center gap-3 mb-2">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <span class="font-semibold">Oops!</span>
            </div>
            <p>{{ productsStore.error }}</p>
            <button
              @click="productsStore.fetchProducts()"
              class="mt-3 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors text-sm font-medium"
            >
              Try Again
            </button>
          </div>

          <!-- Empty State -->
          <div v-else-if="productsStore.products.length === 0" class="text-center py-20">
            <div class="inline-flex items-center justify-center w-24 h-24 rounded-full bg-gradient-to-br from-gray-100 to-gray-50 dark:from-brand-700 dark:to-brand-800 mb-6">
              <svg class="w-12 h-12 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
              </svg>
            </div>
            <p class="text-2xl font-semibold text-gray-700 dark:text-gray-300 mb-2">No products found</p>
            <p class="text-gray-500 dark:text-gray-400 mb-6 max-w-md mx-auto">Try adjusting your filters or search terms to find what you're looking for.</p>
            <button
              @click="clearFilters"
              class="px-6 py-3 bg-gradient-to-r from-indigo-600 to-indigo-700 text-white font-semibold rounded-xl hover:from-indigo-700 hover:to-indigo-800 transition-all duration-300 shadow-lg shadow-indigo-500/30 hover:shadow-xl"
            >
              Clear All Filters
            </button>
          </div>

          <!-- Product Grid -->
          <div v-else class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
            <ProductCard
              v-for="product in productsStore.products"
              :key="product.product_id"
              :product="product"
              @click="handleProductClick(product.product_id)"
            />
          </div>

          <!-- Pagination Controls -->
          <PaginationControls
            v-if="productsStore.pagination.total_pages > 1"
            :current-page="productsStore.pagination.page"
            :total-pages="productsStore.pagination.total_pages"
            @page-change="handlePageChange"
          />
        </div>
      </div>
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