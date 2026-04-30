<script setup>
import { ref, computed, onMounted } from 'vue'
import { useProductsStore } from '@/stores/products'

const productsStore = useProductsStore()

const localFilters = ref({
  search: productsStore.filters.search,
  category: productsStore.filters.category,
  minPrice: productsStore.filters.minPrice,
  maxPrice: productsStore.filters.maxPrice
})

onMounted(async () => {
  if (productsStore.categories.length === 0) {
    await productsStore.fetchCategories()
  }
})

const activeFilterCount = computed(() => {
  let count = 0
  if (localFilters.value.search) count++
  if (localFilters.value.category) count++
  if (localFilters.value.minPrice) count++
  if (localFilters.value.maxPrice) count++
  return count
})

const applyFilters = () => {
  productsStore.setFilters(localFilters.value)
}

const clearFilters = () => {
  localFilters.value = { search: '', category: '', minPrice: null, maxPrice: null }
  productsStore.resetFilters()
}

const removeFilter = (filterName) => {
  if (filterName === 'search') localFilters.value.search = ''
  if (filterName === 'category') localFilters.value.category = ''
  if (filterName === 'minPrice') localFilters.value.minPrice = null
  if (filterName === 'maxPrice') localFilters.value.maxPrice = null
  applyFilters()
}
</script>

<template>
  <aside class="w-full md:w-64 flex-shrink-0">
    <div class="bg-white p-4 rounded-xl border border-gray-200 space-y-4">
      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Search</h3>
        <input
          v-model="localFilters.search"
          type="text"
          placeholder="Search products..."
          class="input"
        />
      </div>

      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Category</h3>
        <select v-model="localFilters.category" class="input">
          <option value="">All Categories</option>
          <option v-for="cat in productsStore.categories" :key="cat.category_id" :value="cat.category_id">
            {{ cat.name }}
          </option>
        </select>
      </div>

      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Price Range</h3>
        <div class="flex gap-2">
          <input
            v-model.number="localFilters.minPrice"
            type="number"
            placeholder="Min"
            class="input"
          />
          <input
            v-model.number="localFilters.maxPrice"
            type="number"
            placeholder="Max"
            class="input"
          />
        </div>
      </div>

      <div class="flex gap-2">
        <button @click="applyFilters" class="btn-primary flex-1">
          Apply{{ activeFilterCount ? ` (${activeFilterCount})` : '' }}
        </button>
        <button @click="clearFilters" class="btn-secondary flex-1">
          Clear
        </button>
      </div>

      <div v-if="activeFilterCount > 0" class="pt-2 border-t border-gray-100">
        <p class="text-xs font-medium text-gray-700 mb-2">Active Filters:</p>
        <div class="flex flex-wrap gap-1">
          <span v-if="localFilters.search" class="inline-flex items-center gap-1 px-2 py-1 bg-indigo-50 text-indigo-700 rounded text-xs">
            "{{ localFilters.search }}"
            <button @click="removeFilter('search')" class="hover:text-indigo-900">&times;</button>
          </span>
          <span v-if="localFilters.category" class="inline-flex items-center gap-1 px-2 py-1 bg-indigo-50 text-indigo-700 rounded text-xs">
            {{ productsStore.categories.find(c => c.category_id === localFilters.category)?.name || localFilters.category }}
            <button @click="removeFilter('category')" class="hover:text-indigo-900">&times;</button>
          </span>
          <span v-if="localFilters.minPrice" class="inline-flex items-center gap-1 px-2 py-1 bg-indigo-50 text-indigo-700 rounded text-xs">
            Min: ₹{{ localFilters.minPrice }}
            <button @click="removeFilter('minPrice')" class="hover:text-indigo-900">&times;</button>
          </span>
          <span v-if="localFilters.maxPrice" class="inline-flex items-center gap-1 px-2 py-1 bg-indigo-50 text-indigo-700 rounded text-xs">
            Max: ₹{{ localFilters.maxPrice }}
            <button @click="removeFilter('maxPrice')" class="hover:text-indigo-900">&times;</button>
          </span>
        </div>
      </div>
    </div>
  </aside>
</template>
