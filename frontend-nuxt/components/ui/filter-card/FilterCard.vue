<script setup>
import { X, Search } from 'lucide-vue-next'
import { useProductsStore } from '~/stores/products'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  },
  categories: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue', 'apply'])

const productsStore = useProductsStore()

const localFilters = ref({
  search: props.modelValue.search || '',
  category: props.modelValue.category || '',
  minPrice: props.modelValue.minPrice || '',
  maxPrice: props.modelValue.maxPrice || ''
})

const applyFilters = () => {
  emit('update:modelValue', { ...localFilters.value })
  emit('apply', { ...localFilters.value })
}

const clearFilters = () => {
  localFilters.value = {
    search: '',
    category: '',
    minPrice: '',
    maxPrice: ''
  }
  emit('update:modelValue', { ...localFilters.value })
  emit('apply', { ...localFilters.value })
}
</script>

<template>
  <div class="bg-surface-container-lowest rounded-xl p-4 space-y-4">
    <div class="relative">
      <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-outline" />
      <input
        v-model="localFilters.search"
        type="text"
        placeholder="Search products..."
        class="w-full pl-10 pr-4 py-2 bg-surface border border-outline-variant rounded-lg text-on_surface placeholder:text-outline text-sm"
      />
    </div>

    <select
      v-model="localFilters.category"
      class="w-full px-3 py-2 bg-surface border border-outline-variant rounded-lg text-on_surface text-sm"
    >
      <option value="">All Categories</option>
      <option v-for="cat in categories" :key="cat.category_id" :value="cat.category_id">
        {{ cat.name }}
      </option>
    </select>

    <div class="grid grid-cols-2 gap-3">
      <input
        v-model="localFilters.minPrice"
        type="number"
        placeholder="Min"
        class="px-3 py-2 bg-surface border border-outline-variant rounded-lg text-on_surface text-sm"
      />
      <input
        v-model="localFilters.maxPrice"
        type="number"
        placeholder="Max"
        class="px-3 py-2 bg-surface border border-outline-variant rounded-lg text-on_surface text-sm"
      />
    </div>

    <div class="flex gap-2">
      <button
        @click="applyFilters"
        class="flex-1 px-4 py-2 bg-primary text-white rounded-lg text-sm font-medium hover:bg-primary/90"
      >
        Apply
      </button>
      <button
        @click="clearFilters"
        class="px-4 py-2 bg-surface-container text-on_surface rounded-lg text-sm font-medium hover:bg-surface-container-high"
      >
        Clear
      </button>
    </div>
  </div>
</template>