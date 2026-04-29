<script setup>
import { ref, onMounted } from 'vue'
import { useMotion } from '@vueuse/motion'
import productService from '@/services/productService'

const inventory = ref([])
const loading = ref(true)

const fetchInventory = async () => {
  loading.value = true
  try {
    const { data } = await productService.getInventory({ limit: 50 })
    inventory.value = data.items || []
  } catch (err) {
    console.error('Failed to fetch inventory', err)
  } finally {
    loading.value = false
  }
}

const updateStock = async (item) => {
  try {
    await productService.updateInventory(item.id, { stock: item.stock })
  } catch (err) {
    console.error('Update failed', err)
  }
}

onMounted(fetchInventory)
</script>

<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Inventory Management</h1>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-gray-100 rounded-lg animate-pulse" />
    </div>

    <div v-else class="bg-white rounded-xl border border-gray-200 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Product</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">SKU</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Current Stock</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Status</th>
            <th class="text-right p-4 text-sm font-medium text-gray-700">Update</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in inventory" :key="item.id" class="border-b border-gray-100 hover:bg-gray-50 transition-colors">
            <td class="p-4">
              <div class="flex items-center gap-3">
                <img :src="item.image || '/placeholder.jpg'" class="w-10 h-10 rounded object-cover" />
                <span class="text-sm text-gray-900">{{ item.name }}</span>
              </div>
            </td>
            <td class="p-4 text-sm text-gray-600">{{ item.sku || 'N/A' }}</td>
            <td class="p-4">
              <input
                v-model.number="item.stock"
                type="number"
                class="input w-24"
                @blur="updateStock(item)"
              />
            </td>
            <td class="p-4">
              <span
                class="px-2 py-1 text-xs rounded-full"
                :class="item.stock > 10 ? 'bg-green-100 text-green-700' : item.stock > 0 ? 'bg-yellow-100 text-yellow-700' : 'bg-red-100 text-red-700'"
              >
                {{ item.stock > 10 ? 'In Stock' : item.stock > 0 ? 'Low Stock' : 'Out of Stock' }}
              </span>
            </td>
            <td class="p-4 text-right">
              <button
                @click="updateStock(item)"
                class="text-sm text-gray-600 hover:text-gray-900"
              >
                Save
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
