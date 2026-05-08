<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProducts } from '@/composables/useProducts'
import ProductCard from '@/components/product/ProductCard.vue'
import FilterCard from '@/components/ui/FilterCard.vue'
import PaginationControls from '@/components/ui/PaginationControls.vue'
import SkeletonCard from '@/components/ui/SkeletonCard.vue'
import { AlertCircle, Inbox } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const productsStore = useProducts()

const filters = ref({
  category: route.query.category || '',
  minPrice: route.query.minPrice || '',
  maxPrice: route.query.maxPrice || '',
  search: route.query.search || ''
})

const appliedFilterCount = computed(() => {
  return Object.values(filters.value).filter(v => v && v.toString().trim()).length
})

const categories = computed(() => productsStore.categories)

const syncToRoute = (page = null) => {
  const query = {}
  if (filters.value.search) query.search = filters.value.search
  if (filters.value.category) query.category = filters.value.category
  if (filters.value.minPrice) query.minPrice = filters.value.minPrice
  if (filters.value.maxPrice) query.maxPrice = filters.value.maxPrice
  if (page && page > 1) query.page = page
  router.replace({ query })
}

const applyFilters = () => {
  syncToRoute(1)
}

const clearFilters = () => {
  filters.value = { category: '', minPrice: '', maxPrice: '', search: '' }
  syncToRoute()
}

const handlePageChange = (page) => {
  syncToRoute(page)
}

const handleProductClick = (productId) => {
  router.push({ name: 'ProductDetail', params: { id: productId } })
}

watch(() => route.query, (query) => {
  filters.value = {
    search: query.search || '',
    category: query.category || '',
    minPrice: query.minPrice || '',
    maxPrice: query.maxPrice || ''
  }
  productsStore.setFilters({
    search: query.search || null,
    category: query.category || null,
    minPrice: query.minPrice || null,
    maxPrice: query.maxPrice || null
  }, true)
  const page = query.page ? Number(query.page) : 1
  productsStore.pagination.page = page
  productsStore.fetchProducts()
}, { immediate: true })

onMounted(async () => {
  await productsStore.fetchCategories()
})
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
              :categories="categories.value"
              :search-value="filters.search"
              :min-price-value="filters.minPrice"
              :max-price-value="filters.maxPrice"
              @update-search="filters.search = $event"
              @update-category="filters.category = $event"
              @update-min-price="filters.minPrice = $event"
              @update-max-price="filters.maxPrice = $event"
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
          <div v-if="productsStore.loading && productsStore.products.length === 0" class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
            <SkeletonCard v-for="n in 9" :key="n" />
          </div>

          <!-- Error State -->
          <div v-else-if="productsStore.error" class="bg-gradient-to-r from-red-50 to-red-100/50 dark:from-red-900/20 dark:to-red-800/20 border border-red-200/50 dark:border-red-700/50 text-red-700 dark:text-red-400 px-6 py-5 rounded-2xl text-center backdrop-blur-sm">
            <div class="flex items-center justify-center gap-3 mb-2">
              <AlertCircle class="w-6 h-6" />
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
          <div v-else-if="!productsStore.products || productsStore.products.length === 0" class="text-center py-20">
            <div class="inline-flex items-center justify-center w-24 h-24 rounded-full bg-gradient-to-br from-gray-100 to-gray-50 dark:from-brand-700 dark:to-brand-800 mb-6">
              <Inbox class="w-12 h-12 text-gray-400 dark:text-gray-500" />
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
              :key="product.productId"
              :product="product"
              @click="handleProductClick(product.productId)"
            />
          </div>

          <!-- Pagination Controls -->
          <PaginationControls
            v-if="productsStore.pagination.totalPages > 1"
            :current-page="productsStore.pagination.page"
            :total-pages="productsStore.pagination.totalPages"
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