<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/lib/productService'
import { useErrorHandler } from '@/composables/useErrorHandler'
import { PAGINATION } from '@/constants'
import { snakeToCamelCase } from '@/lib/mapper'

const { showError, showSuccess } = useErrorHandler()

const inventory = ref([])
const loading = ref(true)

const fetchInventory = async () => {
  loading.value = true
  try {
    const [productsRes, inventoryRes] = await Promise.all([
      productService.getProducts({ limit: PAGINATION.INVENTORY_LIMIT }),
      productService.getInventory({ limit: PAGINATION.INVENTORY_LIMIT }).catch(err => {
        showError(err, 'Failed to fetch inventory data')
        return { data: [] }
      })
    ])

    const productsData = snakeToCamelCase(productsRes.data?.data || productsRes.data || [])
    let inventoryData = []
    try {
      const rawData = Array.isArray(inventoryRes.data)
        ? inventoryRes.data
        : (inventoryRes.data?.data || [])
      inventoryData = snakeToCamelCase(rawData)
    } catch (err) {
      showError(err, 'Failed to parse inventory data')
    }

    const inventoryByProduct = {}
    inventoryData.forEach(item => {
      if (!inventoryByProduct[item.productId]) {
        inventoryByProduct[item.productId] = []
      }
      inventoryByProduct[item.productId].push(item)
    })

    inventory.value = productsData.map(product => {
      const inv = inventoryByProduct[product.productId]
      if (inv && inv.length > 0) {
        return {
          productName: product.name,
          productId: product.productId,
          warehouseId: inv[0].warehouseId,
          quantity: inv.reduce((sum, i) => sum + i.quantity, 0),
          lastUpdated: inv[0].lastUpdated
        }
      }
      return {
        productName: product.name,
        productId: product.productId,
        warehouseId: 'N/A',
        quantity: 0,
        lastUpdated: null
      }
    })
  } catch (err) {
    showError(err, 'Failed to fetch inventory')
  } finally {
    loading.value = false
  }
}

const updateStock = async (item) => {
  try {
    await productService.updateInventory(item.productId, { quantity: item.quantity })
    showSuccess('Stock updated')
  } catch (err) {
    showError(err, 'Failed to update stock')
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
          <tr v-for="item in inventory" :key="item.productId" class="border-b border-gray-100 dark:border-brand-700 hover:bg-gray-50 dark:hover:bg-brand-700/50 transition-colors">
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">{{ item.productName || '(no product data)' }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ item.productId }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ item.warehouseId || 'N/A' }}</td>
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
