<script setup>
import { computed } from 'vue'

const props = defineProps({
  searchValue: { type: String, default: '' },
  categoryValue: { type: String, default: '' },
  minPriceValue: { type: String, default: '' },
  maxPriceValue: { type: String, default: '' },
  categories: { type: Array, default: () => [] },
})

const emit = defineEmits(['update:search', 'update:category', 'update:minPrice', 'update:maxPrice', 'apply', 'clear'])

const categoryOptions = computed(() => [
  { value: '', label: 'All Categories' },
  ...props.categories.map(cat => ({ value: cat.category_id, label: cat.name }))
])
</script>

<template>
  <aside 
    class="glass-card rounded-2xl p-6 mb-8 transition-all duration-500 ease-out hover:shadow-3xl hover:bg-white/25 hover:backdrop-blur-3xl"
    aria-label="Product filters"
  >
    <form @submit.prevent="$emit('apply')" class="space-y-5">
      <!-- Search Input -->
      <div class="relative group">
        <label for="search-input" class="block text-sm font-semibold text-gray-800 mb-2">Search Products</label>
        <div class="relative">
          <div class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 group-focus-within:text-gray-900 transition-colors duration-300">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
          <input
            id="search-input"
            type="text"
            :value="searchValue"
            @input="$emit('update:search', $event.target.value)"
            placeholder="Search products..."
            class="w-full pl-12 pr-4 py-3.5 rounded-xl border-2 border-white/30 bg-white/15 backdrop-blur-xl text-gray-900 placeholder:text-gray-500/70 transition-all duration-300 hover:bg-white/20 hover:border-white/40 focus:bg-white/25 focus:border-gray-400/60 focus:ring-4 focus:ring-gray-900/20 focus:outline-none"
          />
        </div>
      </div>

      <!-- Category Dropdown -->
      <div class="relative">
        <label for="category-select" class="block text-sm font-semibold text-gray-800 mb-2">Category</label>
        <div class="relative">
          <select
            id="category-select"
            :value="categoryValue"
            @change="$emit('update:category', $event.target.value)"
            class="w-full px-4 py-3.5 rounded-xl border-2 border-white/30 bg-white/15 backdrop-blur-xl text-gray-900 transition-all duration-300 hover:bg-white/20 hover:border-white/40 focus:bg-white/25 focus:border-gray-400/60 focus:ring-4 focus:ring-gray-900/20 focus:outline-none appearance-none cursor-pointer pr-10"
          >
            <option v-for="option in categoryOptions" :key="option.value" :value="option.value" class="bg-white/95 backdrop-blur-xl">
              {{ option.label }}
            </option>
          </select>
          <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M19 9l-7 7-7-7"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Price Range -->
      <div class="grid grid-cols-2 gap-4">
        <div class="relative">
          <label for="min-price" class="block text-sm font-semibold text-gray-800 mb-2">Min Price</label>
          <div class="relative">
            <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 font-medium">₹</span>
            <input
              id="min-price"
              type="number"
              :value="minPriceValue"
              @input="$emit('update:minPrice', $event.target.value)"
              placeholder="0"
              min="0"
              step="0.01"
              class="w-full pl-8 pr-4 py-3.5 rounded-xl border-2 border-white/30 bg-white/15 backdrop-blur-xl text-gray-900 placeholder:text-gray-500/70 transition-all duration-300 hover:bg-white/20 hover:border-white/40 focus:bg-white/25 focus:border-gray-400/60 focus:ring-4 focus:ring-gray-900/20 focus:outline-none"
            />
          </div>
        </div>
        <div class="relative">
          <label for="max-price" class="block text-sm font-semibold text-gray-800 mb-2">Max Price</label>
          <div class="relative">
            <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 font-medium">₹</span>
            <input
              id="max-price"
              type="number"
              :value="maxPriceValue"
              @input="$emit('update:maxPrice', $event.target.value)"
              placeholder="9999"
              min="0"
              step="0.01"
              class="w-full pl-8 pr-4 py-3.5 rounded-xl border-2 border-white/30 bg-white/15 backdrop-blur-xl text-gray-900 placeholder:text-gray-500/70 transition-all duration-300 hover:bg-white/20 hover:border-white/40 focus:bg-white/25 focus:border-gray-400/60 focus:ring-4 focus:ring-gray-900/20 focus:outline-none"
            />
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex flex-col sm:flex-row gap-3 pt-3">
        <button
          type="submit"
          class="flex-1 px-6 py-3.5 bg-gradient-to-r from-gray-900 to-gray-800 text-white font-semibold rounded-xl hover:from-gray-800 hover:to-gray-700 active:scale-[0.98] transition-all duration-300 transform hover:shadow-lg hover:shadow-gray-900/30 focus:ring-4 focus:ring-gray-900/30 focus:outline-none flex items-center justify-center gap-2 border border-gray-700/20"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
          Apply Filters
        </button>
        <button
          type="button"
          @click="$emit('clear')"
          class="px-6 py-3.5 border-2 border-gray-300/60 text-gray-700 font-semibold rounded-xl hover:bg-white/30 hover:border-white/50 active:scale-[0.98] transition-all duration-300 focus:ring-4 focus:ring-gray-300/50 focus:outline-none backdrop-blur-sm"
        >
          Clear All
        </button>
      </div>
    </form>
  </aside>
</template>

<style scoped>
.glass-card {
  background: linear-gradient(135deg, 
    rgba(255, 255, 255, 0.25) 0%, 
    rgba(255, 255, 255, 0.15) 50%, 
    rgba(255, 255, 255, 0.05) 100%
  );
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid rgba(255, 255, 255, 0.35);
  box-shadow: 
    0 4px 6px -1px rgba(0, 0, 0, 0.05), 
    0 10px 20px -5px rgba(99, 102, 241, 0.15),
    0 0 0 1px rgba(255, 255, 255, 0.1) inset;
}

/* Custom scrollbar for select dropdown */
select {
  scrollbar-width: thin;
  scrollbar-color: rgba(99, 102, 241, 0.5) transparent;
}

select::-webkit-scrollbar {
  width: 6px;
}

select::-webkit-scrollbar-track {
  background: transparent;
}

select::-webkit-scrollbar-thumb {
  background: rgba(99, 102, 241, 0.4);
  border-radius: 3px;
}

/* Input number spinners hide */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}
</style>