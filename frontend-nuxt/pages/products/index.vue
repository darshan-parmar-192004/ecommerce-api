<script setup>
const route = useRoute()
const router = useRouter()
const page = computed(() => Number(route.query.page) || 1)
const limit = 12

const search = ref(route.query.search || '')
const minPrice = ref(route.query.min_price || '')
const maxPrice = ref(route.query.max_price || '')
const selectedCategory = ref(route.query.category || '')

const queryParams = computed(() => {
  const params = { 
    page: page.value, 
    limit: limit 
  }
  if (selectedCategory.value) params.category = selectedCategory.value
  if (search.value) params.search = search.value
  if (minPrice.value) params.min_price = minPrice.value
  if (maxPrice.value) params.max_price = maxPrice.value
  return params
})

const { data: productsData, pending: productsLoading, error: productsError, refresh: refreshProducts } = await useFetch('/api/products', {
  query: queryParams,
  key: () => `products-${page.value}-${limit}-${selectedCategory.value}-${search.value}-${minPrice.value}-${maxPrice.value}`,
  watch: [() => page.value, () => limit, () => selectedCategory.value, () => search.value, () => minPrice.value, () => maxPrice.value]
})

const { data: categoriesData } = await useFetch('/api/categories')

const products = computed(() => productsData.value?.data || [])
const pagination = computed(() => productsData.value?.pagination || { page: 1, total_pages: 1, total_items: 0 })
const categories = computed(() => categoriesData.value?.data || [])

const applyFilters = () => {
  router.push({
    query: {
      page: 1,
      ...(search.value && { search: search.value }),
      ...(minPrice.value && { min_price: minPrice.value }),
      ...(maxPrice.value && { max_price: maxPrice.value }),
      ...(selectedCategory.value && { category: selectedCategory.value })
    }
  })
}

const clearFilters = () => {
  search.value = ''
  minPrice.value = ''
  maxPrice.value = ''
  selectedCategory.value = ''
  router.push({ query: { page: 1 } })
}

useSeoMeta({
  title: 'Products - E-Commerce Store',
  description: 'Browse our collection of products'
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900">Our Products</h1>
        <p class="text-gray-600 mt-2">Browse our collection of high-quality products</p>
      </div>

      <div class="flex flex-col lg:flex-row gap-8">
        <aside class="w-full lg:w-72 flex-shrink-0">
          <div class="bg-white rounded-xl shadow-sm p-6 sticky top-24 transition-all duration-300 hover:shadow-md">
            <div class="flex items-center justify-between mb-6">
              <h3 class="font-semibold text-gray-900">Filters</h3>
              <button @click="clearFilters" class="text-sm text-primary-600 hover:text-primary-700">
                Clear all
              </button>
            </div>
            
            <div class="space-y-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Search</label>
                <div class="relative">
                  <input 
                    type="text"
                    v-model="search"
                    placeholder="Search products..."
                    class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                    @keyup.enter="applyFilters"
                  />
                  <svg class="absolute left-3 top-2.5 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Category</label>
                <select 
                  v-model="selectedCategory"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                >
                  <option value="">All Categories</option>
                  <option v-for="cat in categories" :key="cat.category_id" :value="cat.category_id">
                    {{ cat.name }}
                  </option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Price Range</label>
                <div class="flex items-center gap-2">
                  <input 
                    type="number"
                    v-model="minPrice"
                    placeholder="Min"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 text-center"
                  />
                  <span class="text-gray-400">-</span>
                  <input 
                    type="number"
                    v-model="maxPrice"
                    placeholder="Max"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 text-center"
                  />
                </div>
              </div>

              <button @click="applyFilters" class="w-full py-2.5 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors font-medium">
                Apply Filters
              </button>
            </div>
          </div>
        </aside>

        <main class="flex-1">
          <div v-if="productsLoading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="i in 6" :key="i" class="bg-white rounded-xl p-4 animate-pulse">
              <div class="bg-gray-200 aspect-square rounded-lg mb-4" />
              <div class="h-4 bg-gray-200 rounded w-3/4 mb-2" />
              <div class="h-4 bg-gray-200 rounded w-1/2" />
            </div>
          </div>

          <div v-else-if="productsError" class="bg-white rounded-xl shadow-sm p-8 text-center">
            <svg class="mx-auto h-16 w-16 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h3 class="mt-4 text-lg font-medium text-gray-900">Failed to load products</h3>
            <p class="mt-2 text-gray-500">Something went wrong. Please try again.</p>
            <button @click="refreshProducts" class="mt-4 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700">
              Try Again
            </button>
          </div>

          <template v-else>
            <div v-if="products.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center">
              <svg class="mx-auto h-24 w-24 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
              <h2 class="mt-4 text-xl font-medium text-gray-900">No products found</h2>
              <p class="mt-2 text-gray-500">Try adjusting your filters or search terms.</p>
              <button @click="clearFilters" class="mt-6 px-6 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700">
                Clear Filters
              </button>
            </div>

            <div v-else>
              <p class="text-gray-600 mb-4">{{ pagination.total_items }} products found</p>
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <ProductCard 
                  v-for="(product, index) in products" 
                  :key="product.product_id" 
                  :product="product"
                  :class="['animate-fade-in-up']"
                  :style="{ animationDelay: `${index * 50}ms` }"
                />
              </div>

              <div v-if="pagination.total_pages > 1" class="mt-12 flex justify-center gap-2">
                <button 
                  :disabled="page <= 1"
                  @click="navigateTo({ query: { ...route.query, page: page - 1 } })"
                  class="px-4 py-2 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                >
                  Previous
                </button>
                
                <div class="flex items-center gap-2">
                  <template v-for="p in pagination.total_pages" :key="p">
                    <button 
                      v-if="p === 1 || p === pagination.total_pages || (p >= page - 1 && p <= page + 1)"
                      @click="navigateTo({ query: { ...route.query, page: p } })"
                      :class="[
                        'w-10 h-10 rounded-lg transition-colors',
                        p === page 
                          ? 'bg-primary-600 text-white' 
                          : 'bg-white border border-gray-300 hover:bg-gray-50'
                      ]"
                    >
                      {{ p }}
                    </button>
                    <span v-else-if="p === page - 2 || p === page + 2" class="text-gray-400">...</span>
                  </template>
                </div>
                
                <button 
                  :disabled="page >= pagination.total_pages"
                  @click="navigateTo({ query: { ...route.query, page: page + 1 } })"
                  class="px-4 py-2 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                >
                  Next
                </button>
              </div>
            </div>
          </template>
        </main>
      </div>
    </div>
  </div>
</template>