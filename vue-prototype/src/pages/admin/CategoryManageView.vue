<script setup>
import { ref, onMounted, computed } from 'vue'
import categoryService from '@/lib/categoryService'
import { snakeToCamelCase } from '@/lib/mapper'
import { Folder, ChevronRight } from 'lucide-vue-next'

const categories = ref([])
const loading = ref(true)

const fetchCategories = async () => {
  loading.value = true
  try {
    const response = await categoryService.getCategories()
    const responseData = response.data || response
    categories.value = snakeToCamelCase(responseData.data || [])
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
    categoryMap[cat.categoryId] = { ...cat, children: [] }
  })

  // Build tree
  categories.value.forEach(cat => {
    if (cat.parentCategoryId && categoryMap[cat.parentCategoryId]) {
      categoryMap[cat.parentCategoryId].children.push(categoryMap[cat.categoryId])
    } else {
      tree.push(categoryMap[cat.categoryId])
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
        <div v-for="category in categoryTree" :key="category.categoryId" class="mb-4">
          <div class="flex items-center gap-3 p-3 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
            <Folder class="w-5 h-5 text-gray-700 dark:text-gray-300" />
            <span class="font-medium text-gray-900 dark:text-gray-100">{{ category.name }}</span>
            <span class="text-xs text-gray-500 dark:text-gray-400 ml-2">{{ category.categoryId }}</span>
          </div>
          <!-- Children -->
          <div v-if="category.children?.length" class="ml-8 mt-2 space-y-2">
            <div v-for="child in category.children" :key="child.categoryId" class="flex items-center gap-3 p-2 border border-gray-200 dark:border-brand-600 rounded-lg">
              <ChevronRight class="w-4 h-4 text-gray-400 dark:text-gray-500" />
              <span class="text-gray-900 dark:text-gray-100">{{ child.name }}</span>
              <span class="text-xs text-gray-500 dark:text-gray-400 ml-auto">{{ child.categoryId }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>