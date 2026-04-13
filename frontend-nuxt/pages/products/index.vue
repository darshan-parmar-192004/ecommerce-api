<script setup>
const route = useRoute()
const router = useRouter()

const limit = 12

const search = ref(route.query.search || '')
const minPrice = ref(route.query.min_price || '')
const maxPrice = ref(route.query.max_price || '')
const selectedCategory = ref(route.query.category || '')
const currentPage = ref(Number(route.query.page) || 1)

const showSuggestions = ref(false)
const searchSuggestions = ref([])
const recentSearches = ref([])
const focusedSuggestion = ref(0)

const debouncedSearch = ref('')
let searchTimeout = null

const RECENT_SEARCHES_KEY = 'recent_searches'

const loadRecentSearches = () => {
  if (import.meta.client) {
    const stored = localStorage.getItem(RECENT_SEARCHES_KEY)
    if (stored) {
      try {
        recentSearches.value = JSON.parse(stored)
      } catch (e) {
        recentSearches.value = []
      }
    }
  }
}

const saveRecentSearch = (searchTerm) => {
  if (!searchTerm || searchTerm.length < 2) return
  const updated = [searchTerm, ...recentSearches.value.filter(s => s !== searchTerm)].slice(0, 5)
  recentSearches.value = updated
  if (import.meta.client) {
    localStorage.setItem(RECENT_SEARCHES_KEY, JSON.stringify(updated))
  }
}

const clearRecentSearches = () => {
  recentSearches.value = []
  if (import.meta.client) {
    localStorage.removeItem(RECENT_SEARCHES_KEY)
  }
}

onMounted(() => {
  loadRecentSearches()
})

const fetchSuggestions = async (query) => {
  if (!query || query.length < 2) {
    searchSuggestions.value = []
    return
  }
  try {
    const data = await $fetch(`/api/products?search=${encodeURIComponent(query)}&limit=5`)
    searchSuggestions.value = data.data || data || []
  } catch (e) {
    searchSuggestions.value = []
  }
}

const handleSearchInput = (e) => {
  const value = e.target.value
  search.value = value
  
  if (searchTimeout) clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    fetchSuggestions(value)
  }, 300)
}

const selectSuggestion = (product) => {
  saveRecentSearch(product.name)
  router.push(`/products/${product.product_id}`)
  showSuggestions.value = false
}

const selectRecentSearch = (term) => {
  search.value = term
  applyFilters()
  showSuggestions.value = false
}

watch(() => route.query, (newQuery) => {
  search.value = newQuery.search || ''
  minPrice.value = newQuery.min_price || ''
  maxPrice.value = newQuery.max_price || ''
  selectedCategory.value = newQuery.category || ''
  currentPage.value = Number(newQuery.page) || 1
}, { immediate: true, deep: true })

const productsData = ref(null)
const productsLoading = ref(false)
const productsError = ref(null)

const loadProducts = async () => {
  productsLoading.value = true
  productsError.value = null
  try {
    const q = new URLSearchParams()
    if (currentPage.value > 1) q.set('page', String(currentPage.value))
    q.set('limit', String(limit))
    if (selectedCategory.value) q.set('category', selectedCategory.value)
    if (debouncedSearch.value) q.set('search', debouncedSearch.value)
    if (minPrice.value) q.set('min_price', minPrice.value)
    if (maxPrice.value) q.set('max_price', maxPrice.value)
    
    productsData.value = await $fetch(`/api/products?${q.toString()}`)
  } catch (e) {
    productsError.value = e
  } finally {
    productsLoading.value = false
  }
}

onMounted(() => {
  loadProducts()
})

watch([selectedCategory, minPrice, maxPrice, debouncedSearch, currentPage], () => {
  loadProducts()
})

const { data: categoriesData } = await useFetch('/api/categories')

const products = computed(() => productsData.value?.data || productsData.value || [])
const pagination = computed(() => productsData.value?.pagination || { page: 1, total_pages: 1, total_items: 0 })
const categories = computed(() => categoriesData.value?.data || categoriesData.value || [])

const applyFilters = () => {
  if (search.value) {
    saveRecentSearch(search.value)
  }
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
  debouncedSearch.value = ''
  router.push({ query: { page: 1 } })
}

const goToPage = (newPage) => {
  if (newPage >= 1 && newPage <= pagination.value.total_pages) {
    router.push({
      query: {
        ...route.query,
        page: newPage
      }
    })
  }
}

watch(search, (newVal) => {
  if (searchTimeout) clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    debouncedSearch.value = newVal
    applyFilters()
  }, 500)
})

watch([minPrice, maxPrice], () => {
  applyFilters()
})

watch(selectedCategory, () => {
  applyFilters()
})

useSeoMeta({
  title: 'Products - E-Commerce Store',
  description: 'Browse our collection of products'
})
</script>

<template>
  <div class="min-h-screen bg-surface py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-on_surface font-display">Our Collection</h1>
        <p class="text-on_surface_variant mt-2">Browse our curated selection of products</p>
      </div>

      <div class="flex flex-col lg:flex-row gap-8">
        <aside class="w-full lg:w-72 flex-shrink-0">
          <div class="bg-surface-container-lowest/80 backdrop-blur-sm rounded-xl p-6 sticky top-24 shadow-ambient transition-all duration-300 hover:shadow-lg border border-white/5">
            <div class="flex items-center justify-between mb-6">
              <h3 class="font-semibold text-on_surface">Filters</h3>
              <button @click="clearFilters" class="text-sm text-primary hover:text-primary/80">
                Clear all
              </button>
            </div>
            
            <div class="space-y-6">
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Search</label>
                <div class="relative">
                  <input 
                    type="text"
                    v-model="search"
                    placeholder="Search products..."
                    class="w-full pl-10 pr-4 py-2.5 bg-surface-container-low border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    @input="handleSearchInput"
                    @focus="showSuggestions = true"
                    @blur="setTimeout(() => showSuggestions = false, 200)"
                    @keyup.enter="applyFilters"
                    @keydown.down.prevent="focusedSuggestion = Math.min(focusedSuggestion + 1, searchSuggestions.length + recentSearches.length)"
                    @keydown.up.prevent="focusedSuggestion = Math.max(focusedSuggestion - 1, 0)"
                    @keydown.enter.prevent="focusedSuggestion > 0 && (
                      focusedSuggestion <= searchSuggestions.length 
                        ? selectSuggestion(searchSuggestions[focusedSuggestion - 1])
                        : selectRecentSearch(recentSearches[focusedSuggestion - searchSuggestions.length - 1])
                    )"
                  />
                  <svg class="absolute left-3 top-3 w-5 h-5 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                  
                  <Transition name="fade">
                    <div 
                      v-if="showSuggestions && (searchSuggestions.length > 0 || recentSearches.length > 0)" 
                      class="absolute z-50 w-full mt-2 bg-surface-container-lowest rounded-lg shadow-xl border border-outline-variant/20 overflow-hidden"
                    >
                      <div v-if="searchSuggestions.length > 0" class="py-2">
                        <p class="px-3 py-1 text-xs font-semibold text-on_surface_variant uppercase">Products</p>
                        <button
                          v-for="(product, idx) in searchSuggestions"
                          :key="product.product_id"
                          @mousedown="selectSuggestion(product)"
                          class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors"
                        >
                          <svg class="w-5 h-5 text-outline flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                          </svg>
                          <div class="flex-1 min-w-0">
                            <p class="text-on_surface truncate">{{ product.name }}</p>
                            <p class="text-xs text-primary font-medium">₹{{ Number(product.price).toFixed(2) }}</p>
                          </div>
                        </button>
                      </div>
                      
                      <div v-if="recentSearches.length > 0" class="border-t border-outline-variant/20 py-2">
                        <div class="flex items-center justify-between px-3 py-1">
                          <p class="text-xs font-semibold text-on_surface_variant uppercase">Recent Searches</p>
                          <button @mousedown="clearRecentSearches" class="text-xs text-primary hover:text-primary/80">Clear</button>
                        </div>
                        <button
                          v-for="(term, idx) in recentSearches"
                          @mousedown="selectRecentSearch(term)"
                          class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors"
                        >
                          <svg class="w-4 h-4 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                          </svg>
                          <span class="text-on_surface">{{ term }}</span>
                        </button>
                      </div>
                    </div>
                  </Transition>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Category</label>
                <select 
                  v-model="selectedCategory"
                  class="w-full px-4 py-2.5 bg-surface-container-low border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface"
                >
                  <option value="">All Categories</option>
                  <option v-for="cat in categories" :key="cat.category_id" :value="cat.category_id">
                    {{ cat.name }}
                  </option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Price Range</label>
                <div class="flex items-center gap-2">
                  <input 
                    type="number"
                    v-model="minPrice"
                    placeholder="Min"
                    class="w-full px-3 py-2.5 bg-surface-container-low border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-center text-on_surface placeholder:text-outline"
                  />
                  <span class="text-outline">-</span>
                  <input 
                    type="number"
                    v-model="maxPrice"
                    placeholder="Max"
                    class="w-full px-3 py-2.5 bg-surface-container-low border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-center text-on_surface placeholder:text-outline"
                  />
                </div>
              </div>

              <button @click="applyFilters" class="w-full py-2.5 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 hover:shadow-lg hover:shadow-primary/20 transition-all font-medium">
                Apply Filters
              </button>
            </div>
          </div>
        </aside>

        <main class="flex-1">
          <div v-if="productsLoading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="i in 6" :key="i" class="bg-surface-container-lowest rounded-xl p-4 animate-pulse">
              <div class="bg-surface-container aspect-square rounded-lg mb-4" />
              <div class="h-4 bg-surface-container rounded w-3/4 mb-2" />
              <div class="h-4 bg-surface-container rounded w-1/2" />
            </div>
          </div>

          <div v-else-if="productsError" class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center">
            <svg class="mx-auto h-16 w-16 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h3 class="mt-4 text-lg font-medium text-on_surface">Failed to load products</h3>
            <p class="mt-2 text-on_surface_variant">Something went wrong. Please try again.</p>
            <button @click="refreshProducts" class="mt-4 px-4 py-2 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
              Try Again
            </button>
          </div>

          <template v-else>
            <!-- Category Filter Chips -->
            <div v-if="categories.length > 0" class="mb-8">
              <div class="flex flex-wrap gap-3">
                <button
                  @click="selectedCategory = ''; applyFilters()"
                  :class="[
                    'px-4 py-2 rounded-full text-sm font-medium transition-all duration-300',
                    selectedCategory === '' 
                      ? 'bg-gradient-to-r from-primary to-primary-container text-white shadow-lg shadow-primary/25' 
                      : 'bg-surface-container text-on_surface_variant hover:bg-surface-container-high hover:scale-105'
                  ]"
                >
                  All
                </button>
                <button
                  v-for="cat in categories"
                  :key="cat.category_id"
                  @click="selectedCategory = cat.category_id; applyFilters()"
                  :class="[
                    'px-4 py-2 rounded-full text-sm font-medium transition-all duration-300',
                    selectedCategory === cat.category_id 
                      ? 'bg-gradient-to-r from-primary to-primary-container text-white shadow-lg shadow-primary/25' 
                      : 'bg-surface-container text-on_surface_variant hover:bg-surface-container-high hover:scale-105'
                  ]"
                >
                  {{ cat.name }}
                </button>
              </div>
            </div>

            <div v-if="products.length === 0" class="bg-surface-container-lowest rounded-xl shadow-ambient p-12 text-center">
              <svg class="mx-auto h-24 w-24 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
              <h2 class="mt-4 text-xl font-medium text-on_surface">No products found</h2>
              <p class="mt-2 text-on_surface_variant">Try adjusting your filters or search terms.</p>
              <button @click="clearFilters" class="mt-6 px-6 py-2 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
                Clear Filters
              </button>
            </div>

            <div v-else>
              <p class="text-on_surface_variant mb-4">{{ pagination.total_items }} products found</p>
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <ProductCard 
                  v-for="(product, index) in products" 
                  :key="product.product_id" 
                  :product="product"
                  :class="['animate-fade-in-up']"
                  :style="{ animationDelay: `${index * 50}ms` }"
                  loading="lazy"
                />
              </div>

              <div v-if="pagination.total_pages > 1" class="mt-12 flex justify-center gap-2">
                <button 
                  :disabled="currentPage <= 1"
                  @click="goToPage(currentPage - 1)"
                  class="px-4 py-2 bg-surface-container-lowest border border-outline-variant rounded-lg hover:bg-surface-container disabled:opacity-50 disabled:cursor-not-allowed transition-colors text-on_surface"
                >
                  Previous
                </button>
                
                <div class="flex items-center gap-2">
                  <template v-for="p in pagination.total_pages" :key="p">
                    <button 
                      v-if="p === 1 || p === pagination.total_pages || (p >= currentPage - 1 && p <= currentPage + 1)"
                      @click="goToPage(p)"
                      :class="[
                        'w-10 h-10 rounded-lg transition-colors',
                        p === currentPage 
                          ? 'bg-gradient-to-r from-primary to-primary-container text-white' 
                          : 'bg-surface-container-lowest border border-outline-variant hover:bg-surface-container text-on_surface'
                      ]"
                    >
                      {{ p }}
                    </button>
                    <span v-else-if="p === currentPage - 2 || p === currentPage + 2" class="text-outline">...</span>
                  </template>
                </div>
                
                <button 
                  :disabled="currentPage >= pagination.total_pages"
                  @click="goToPage(currentPage + 1)"
                  class="px-4 py-2 bg-surface-container-lowest border border-outline-variant rounded-lg hover:bg-surface-container disabled:opacity-50 disabled:cursor-not-allowed transition-colors text-on_surface"
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