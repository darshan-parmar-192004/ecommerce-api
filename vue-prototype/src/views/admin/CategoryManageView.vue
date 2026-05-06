<script setup>
import { ref, onMounted, computed } from 'vue'
import categoryService from '@/services/categoryService'

const categories = ref([])
const loading = ref(true)

const fetchCategories = async () => {
  loading.value = true
  try {
    const response = await categoryService.getCategories()
    const responseData = response.data || response
    categories.value = responseData.data || []
  } catch (err) {
    console.error('Failed to fetch categories', err)
  } finally {
    loading.value = false
  }
}

// Build hierarchy tree
const categoryTree = computed(() => {
  const tree = []
  const categoryMap = {}
  
  // Initialize map
  categories.value.forEach(cat => {
    categoryMap[cat.category_id] = { ...cat, children: [] }
  })
  
  // Build tree
  categories.value.forEach(cat => {
    if (cat.parent_category_id && categoryMap[cat.parent_category_id]) {
      categoryMap[cat.parent_category_id].children.push(categoryMap[cat.category_id])
    } else {
      tree.push(categoryMap[cat.category_id])
    }
  })
  
  return tree
})

onMounted(fetchCategories)
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">Category Hierarchy</h1>
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-gray-100 dark:bg-brand-700 rounded-xl animate-pulse" />
    </div>

    <div v-else class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 p-6">
      <div v-if="categoryTree.length === 0" class="text-gray-500 dark:text-gray-400 text-center py-8">
        No categories found
      </div>
      <div v-else class="space-y-4">
        <div v-for="category in categoryTree" :key="category.category_id" class="mb-4">
          <div class="flex items-center gap-3 p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
            <svg class="w-5 h-5 text-gray-700 dark:text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
            </svg>
            <span class="font-medium text-gray-900 dark:text-gray-100">{{ category.name }}</span>
            <span class="text-xs text-gray-500 dark:text-gray-400 ml-2">{{ category.category_id }}</span>
          </div>
          <!-- Children -->
          <div v-if="category.children?.length" class="ml-8 mt-2 space-y-2">
            <div v-for="child in category.children" :key="child.category_id" class="flex items-center gap-3 p-2 border border-gray-200 dark:border-brand-600 rounded-lg">
              <svg class="w-4 h-4 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
              <span class="text-gray-900 dark:text-gray-100">{{ child.name }}</span>
              <span class="text-xs text-gray-500 dark:text-gray-400 ml-auto">{{ child.category_id }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>