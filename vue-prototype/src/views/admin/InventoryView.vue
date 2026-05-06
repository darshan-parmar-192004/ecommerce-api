<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/services/productService'
import { useToastStore } from '@/stores/toast'
import { PAGINATION } from '@/constants'

const toastStore = useToastStore()

const inventory = ref([])
const loading = ref(true)

const fetchInventory = async () => {
  loading.value = true
  try {
    const productsRes = await productService.getProducts({ limit: PAGINATION.INVENTORY_LIMIT })
    const productsData = productsRes.data?.data || productsRes.data || []
    
    let inventoryData = []
    try {
      const inventoryRes = await productService.getInventory({ limit: PAGINATION.INVENTORY_LIMIT })
      inventoryData = Array.isArray(inventoryRes.data)
        ? inventoryRes.data
        : (inventoryRes.data?.data || [])
    } catch {
      // Inventory endpoint may have issues; continue with just products
    }
    
    const inventoryByProduct = {}
    inventoryData.forEach(item => {
      if (!inventoryByProduct[item.product_id]) {
        inventoryByProduct[item.product_id] = []
      }
      inventoryByProduct[item.product_id].push(item)
    })
    
    inventory.value = productsData.map(product => {
      const inv = inventoryByProduct[product.product_id]
      if (inv && inv.length > 0) {
        return {
          product_name: product.name,
          product_id: product.product_id,
          warehouse_id: inv[0].warehouse_id,
          quantity: inv.reduce((sum, i) => sum + i.quantity, 0),
          last_updated: inv[0].last_updated
        }
      }
      return {
        product_name: product.name,
        product_id: product.product_id,
        warehouse_id: 'N/A',
        quantity: 0,
        last_updated: null
      }
    })
  } catch (err) {
    toastStore.error('Failed to fetch inventory')
  } finally {
    loading.value = false
  }
}

const updateStock = async (item) => {
  try {
    await productService.updateInventory(item.product_id, { quantity: item.quantity })
    toastStore.success('Stock updated')
  } catch (err) {
    toastStore.error('Failed to update stock')
  }
}

onMounted(fetchInventory)
</script>

<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Inventory Management</h1>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-gray-100 dark:bg-brand-700 rounded-xl animate-pulse" />
    </div>

    <div v-else class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50 dark:bg-brand-700 border-b border-gray-200 dark:border-brand-600">
          <tr>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Product Name</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Product ID</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Warehouse</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Current Stock</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Status</th>
            <th class="text-right p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Update</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in inventory" :key="item.product_id" class="border-b border-gray-100 dark:border-brand-700 hover:bg-gray-50 dark:hover:bg-brand-700/50 transition-colors">
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">{{ item.product_name || '(no product data)' }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ item.product_id }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ item.warehouse_id || 'N/A' }}</td>
            <td class="p-4">
              <input
                v-model.number="item.quantity"
                type="number"
                class="w-24 px-3 py-1 border border-gray-300 dark:border-brand-600 dark:bg-brand-700 dark:text-gray-100 rounded-lg text-sm"
              />
            </td>
            <td class="p-4">
              <span
                class="px-2 py-1 text-xs rounded-full"
                :class="item.quantity > 10 ? 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400' : item.quantity > 0 ? 'bg-yellow-100 dark:bg-yellow-900/30 text-yellow-700 dark:text-yellow-400' : 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-400'"
              >
                {{ item.quantity > 10 ? 'In Stock' : item.quantity > 0 ? 'Low Stock' : 'Out of Stock' }}
              </span>
            </td>
            <td class="p-4 text-right">
              <button
                @click="updateStock(item)"
                class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
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