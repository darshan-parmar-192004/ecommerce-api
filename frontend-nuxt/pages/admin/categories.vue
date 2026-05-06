<script setup>
import { ref, onMounted, computed } from 'vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

const categories = ref([])
const loading = ref(true)

const categoryTree = computed(() => {
  const tree = []
  const categoryMap = {}
  
  categories.value.forEach(cat => {
    categoryMap[cat.category_id] = { ...cat, children: [] }
  })
  
  categories.value.forEach(cat => {
    if (cat.parent_category_id && categoryMap[cat.parent_category_id]) {
      categoryMap[cat.parent_category_id].children.push(categoryMap[cat.category_id])
    } else {
      tree.push(categoryMap[cat.category_id])
    }
  })
  
  return tree
})

const fetchCategories = async () => {
  loading.value = true
  try {
    const { data } = await useApi('/categories')
    categories.value = data || []
  } catch (err) {
    console.error('Failed to fetch categories', err)
  } finally {
    loading.value = false
  }
}

onMounted(fetchCategories)
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Category Hierarchy</h1>

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
            <span class="font-medium text-gray-900 dark:text-gray-100">{{ category.name }}</span>
            <span class="text-xs text-gray-500 dark:text-gray-400 ml-2">{{ category.category_id }}</span>
          </div>
          <div v-if="category.children?.length" class="ml-8 mt-2 space-y-2">
            <div v-for="child in category.children" :key="child.category_id" class="flex items-center gap-3 p-2 border border-gray-200 dark:border-brand-600 rounded-lg">
              <span class="text-gray-900 dark:text-gray-100">{{ child.name }}</span>
              <span class="text-xs text-gray-500 dark:text-gray-400 ml-auto">{{ child.category_id }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>