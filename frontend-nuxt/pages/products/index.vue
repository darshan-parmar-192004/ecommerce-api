<script setup>
import ProductCard from '~/components/products/ProductCard.vue'

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

const { products: productsApi, categories: categoriesApi } = useApi()

const fetchSuggestions = async (query) => {
  if (!query || query.length < 2) {
    searchSuggestions.value = []
    return
  }
  try {
    const data = await productsApi.list({ search: query, limit: 5 })
    searchSuggestions.value = data.data || []
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
    const response = await productsApi.list({
      page: currentPage.value > 1 ? currentPage.value : undefined,
      limit,
      category: selectedCategory.value,
      search: debouncedSearch.value,
      min_price: minPrice.value,
      max_price: maxPrice.value
    })
    productsData.value = response
  } catch (e) {
    productsError.value = e
  } finally {
    productsLoading.value = false
  }
}

const refreshProducts = () => {
  loadProducts()
}

onMounted(() => {
  loadProducts()
})

watch([selectedCategory, minPrice, maxPrice, debouncedSearch, currentPage], () => {
  loadProducts()
})

const { data: categoriesData } = useAsyncData('categories-all', () => categoriesApi.list())

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
              <Button @click="clearFilters" label="Clear all" text size="small" class="!text-sm" />
            </div>
            
            <div class="space-y-6">
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Search</label>
                <div class="relative">
                  <IconField>
                    <InputIcon class="pi pi-search" />
                    <InputText
                      v-model="search"
                      placeholder="Search products..."
                      class="w-full"
                      @input="handleSearchInput"
                      @focus="showSuggestions = true"
                      @blur="setTimeout(() => showSuggestions = false, 200)"
                      @keyup.enter="applyFilters"
                    />
                  </IconField>
                  
                  <Transition name="fade">
                    <div 
                      v-if="showSuggestions && (searchSuggestions.length > 0 || recentSearches.length > 0)" 
                      class="absolute z-50 w-full mt-2 bg-surface-container-lowest rounded-lg shadow-xl border border-outline-variant/20 overflow-hidden"
                    >
                      <div v-if="searchSuggestions.length > 0" class="py-2">
                        <p class="px-3 py-1 text-xs font-semibold text-on_surface_variant uppercase">Products</p>
                        <div
                          v-for="product in searchSuggestions"
                          :key="product.product_id"
                          @mousedown="selectSuggestion(product)"
                          class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors cursor-pointer"
                        >
                          <i class="pi pi-box text-outline flex-shrink-0" />
                          <div class="flex-1 min-w-0">
                            <p class="text-on_surface truncate">{{ product.name }}</p>
                            <p class="text-xs text-primary font-medium">₹{{ Number(product.price).toFixed(2) }}</p>
                          </div>
                        </div>
                      </div>
                      
                      <div v-if="recentSearches.length > 0" class="border-t border-outline-variant/20 py-2">
                        <div class="flex items-center justify-between px-3 py-1">
                          <p class="text-xs font-semibold text-on_surface_variant uppercase">Recent Searches</p>
                          <Button @mousedown="clearRecentSearches" label="Clear" text size="small" class="!text-xs" />
                        </div>
                        <div
                          v-for="term in recentSearches"
                          :key="term"
                          @mousedown="selectRecentSearch(term)"
                          class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors cursor-pointer"
                        >
                          <i class="pi pi-clock text-outline text-xs" />
                          <span class="text-on_surface">{{ term }}</span>
                        </div>
                      </div>
                    </div>
                  </Transition>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Category</label>
                <Select
                  v-model="selectedCategory"
                  :options="categories"
                  optionLabel="name"
                  optionValue="category_id"
                  placeholder="All Categories"
                  class="w-full"
                  @change="applyFilters"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Price Range</label>
                <div class="flex items-center gap-2">
                  <InputNumber v-model="minPrice" placeholder="Min" class="w-full" :min="0" />
                  <span class="text-outline">-</span>
                  <InputNumber v-model="maxPrice" placeholder="Max" class="w-full" :min="0" />
                </div>
              </div>

              <Button @click="applyFilters" label="Apply Filters" class="w-full" />
            </div>
          </div>
        </aside>

        <main class="flex-1">
          <div v-if="productsLoading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="i in 6" :key="i" class="bg-surface-container-lowest rounded-xl p-4 animate-pulse">
              <div class="aspect-[4/3] bg-surface-container rounded-lg mb-4" />
              <div class="h-4 bg-surface-container rounded w-3/4 mb-2" />
              <div class="h-4 bg-surface-container rounded w-1/2" />
            </div>
          </div>

          <div v-else-if="productsError" class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center">
            <i class="pi pi-exclamation-triangle text-5xl text-error mb-4" />
            <h3 class="mt-4 text-lg font-medium text-on_surface">Failed to load products</h3>
            <p class="mt-2 text-on_surface_variant">Something went wrong. Please try again.</p>
            <Button @click="refreshProducts" label="Try Again" class="mt-4" />
          </div>

          <template v-else>
            <div v-if="categories.length > 0" class="mb-8">
              <div class="flex flex-wrap gap-3">
                <Button
                  @click="selectedCategory = ''; applyFilters()"
                  :label="'All'"
                  :class="[
                    selectedCategory === '' ? 'bg-gradient-to-r from-primary to-primary-container text-white shadow-lg shadow-primary/25' : 'bg-surface-container text-on_surface_variant hover:bg-surface-container-high'
                  ]"
                  size="small"
                />
                <Button
                  v-for="cat in categories"
                  :key="cat.category_id"
                  @click="selectedCategory = cat.category_id; applyFilters()"
                  :label="cat.name"
                  :class="[
                    selectedCategory === cat.category_id ? 'bg-gradient-to-r from-primary to-primary-container text-white shadow-lg shadow-primary/25' : 'bg-surface-container text-on_surface_variant hover:bg-surface-container-high'
                  ]"
                  size="small"
                />
              </div>
            </div>

            <div v-if="products.length === 0" class="bg-surface-container-lowest rounded-xl shadow-ambient p-12 text-center">
              <i class="pi pi-box text-6xl text-outline mb-4" />
              <h2 class="mt-4 text-xl font-medium text-on_surface">No products found</h2>
              <p class="mt-2 text-on_surface_variant">Try adjusting your filters or search terms.</p>
              <Button @click="clearFilters" label="Clear Filters" class="mt-6" />
            </div>

            <div v-else>
              <p class="text-on_surface_variant mb-4">{{ pagination.total_items }} products found</p>
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <ProductCard 
                  v-for="(product, index) in products" 
                  :key="product.product_id" 
                  :product="product"
                  class="animate-fade-in-up"
                  :style="{ animationDelay: `${index * 50}ms` }"
                />
              </div>

              <div v-if="pagination.total_pages > 1" class="mt-12 flex justify-center gap-2">
                <Paginator
                  :first="(currentPage - 1) * limit"
                  :rows="limit"
                  :totalRecords="pagination.total_items"
                  @page="goToPage($event.page + 1)"
                  class="mt-6"
                />
              </div>
            </div>
          </template>
        </main>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.animate-fade-in-up {
  animation: fadeInUp 0.4s ease-out forwards;
  opacity: 0;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
