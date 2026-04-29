<script setup>
import { useProductsStore } from '@/stores/products'

const productsStore = useProductsStore()

const applyFilters = () => {
  productsStore.setFilters(productsStore.filters)
}
</script>

<template>
  <aside class="w-full md:w-64 flex-shrink-0">
    <div class="bg-white p-4 rounded-xl border border-gray-200 space-y-4">
      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Search</h3>
        <input
          v-model="productsStore.filters.search"
          @input="applyFilters"
          type="text"
          placeholder="Search products..."
          class="input"
        />
      </div>

      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Category</h3>
        <select v-model="productsStore.filters.category" @change="applyFilters" class="input">
          <option value="">All Categories</option>
          <option value="electronics">Electronics</option>
          <option value="clothing">Clothing</option>
          <option value="home">Home & Living</option>
        </select>
      </div>

      <div>
        <h3 class="text-sm font-semibold text-gray-900 mb-2">Price Range</h3>
        <div class="flex gap-2">
          <input
            v-model.number="productsStore.filters.minPrice"
            @input="applyFilters"
            type="number"
            placeholder="Min"
            class="input"
          />
          <input
            v-model.number="productsStore.filters.maxPrice"
            @input="applyFilters"
            type="number"
            placeholder="Max"
            class="input"
          />
        </div>
      </div>

      <button @click="productsStore.resetFilters()" class="text-sm text-gray-600 hover:text-gray-900 w-full text-center">
        Reset Filters
      </button>
    </div>
  </aside>
</template>
