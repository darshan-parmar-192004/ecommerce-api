<script setup>
import { ref, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

const products = ref([])
const loading = ref(true)
const showModal = ref(false)
const editingProduct = ref(null)

const form = ref({
  name: '',
  price: 0,
  category_id: '',
  description: '',
  stock: 0
})

const fetchProducts = async () => {
  loading.value = true
  try {
    const { data } = await useApi('/products')
    products.value = data || []
  } catch (err) {
    console.error('Failed to fetch products', err)
  } finally {
    loading.value = false
  }
}

const editProduct = (product) => {
  editingProduct.value = product
  form.value = { ...product }
  showModal.value = true
}

const saveProduct = async () => {
  try {
    if (editingProduct.value) {
      await useApi(`/products/${editingProduct.value.product_id}`, {
        method: 'PUT',
        body: form.value
      })
    } else {
      await useApi('/products', {
        method: 'POST',
        body: form.value
      })
    }
    showModal.value = false
    await fetchProducts()
  } catch (err) {
    console.error('Failed to save product', err)
  }
}

onMounted(fetchProducts)
</script>

<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">Manage Products</h1>
      <button @click="showModal = true; editingProduct.value = null; form = { name: '', price: 0, category_id: '', description: '', stock: 0 }" class="btn-primary">
        Add Product
      </button>
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-gray-100 dark:bg-brand-700 rounded-lg animate-pulse" />
    </div>

    <div v-else class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50 dark:bg-brand-700 border-b border-gray-200 dark:border-brand-600">
          <tr>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Name</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Price</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Stock</th>
            <th class="text-right p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in products" :key="product.product_id" class="border-b border-gray-100 dark:border-brand-700 hover:bg-gray-50 dark:hover:bg-brand-700/50">
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">{{ product.name }}</td>
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">${{ product.price?.toFixed(2) }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ product.stock ?? 0 }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editProduct(product)" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100">Edit</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal placeholder -->
    <BaseModal v-if="showModal" :show="showModal" :title="editingProduct ? 'Edit Product' : 'Add Product'" @close="showModal = false">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Product Name</label>
          <input v-model="form.name" placeholder="Product Name" class="input" required />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Price</label>
          <input v-model.number="form.price" type="number" step="0.01" placeholder="Price" class="input" required />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Stock</label>
          <input v-model.number="form.stock" type="number" placeholder="Stock" class="input" required />
        </div>
      </div>
    </BaseModal>
  </div>
</template>