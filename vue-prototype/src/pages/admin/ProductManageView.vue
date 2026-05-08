<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/lib/productService'
import categoryService from '@/lib/categoryService'
import { useErrorHandler } from '@/composables/useErrorHandler'
import BaseModal from '@/components/ui/BaseModal.vue'
import { PAGINATION } from '@/constants'
import { snakeToCamelCase, camelToSnakeCase } from '@/lib/mapper'

const { showError, showSuccess } = useErrorHandler()

const products = ref([])
const categories = ref([])
const loading = ref(true)
const showModal = ref(false)
const showDeleteModal = ref(false)
const editingProduct = ref(null)
const productToDelete = ref(null)

const form = ref({
  productId: '',
  name: '',
  price: 0,
  categoryId: '',
  description: ''
})

const fetchProducts = async () => {
  loading.value = true
  try {
    const response = await productService.getProducts({ limit: PAGINATION.PRODUCT_MANAGE_LIMIT })
    const responseData = response.data || response
    products.value = snakeToCamelCase(responseData.data || [])
  } catch (err) {
    showError(err, 'Failed to fetch products')
  } finally {
    loading.value = false
  }
}

const fetchCategories = async () => {
  try {
    const response = await categoryService.getCategories()
    const responseData = response.data || response
    categories.value = snakeToCamelCase(responseData.data || [])
  } catch (err) {
    showError(err, 'Failed to fetch categories')
  }
}

const editProduct = (product) => {
  editingProduct.value = product
  form.value = {
    productId: product.productId,
    name: product.name,
    price: product.price,
    categoryId: product.categoryId,
    description: product.description || ''
  }
  showModal.value = true
}

const confirmDelete = (id) => {
  productToDelete.value = id
  showDeleteModal.value = true
}

const deleteProduct = async () => {
  if (!productToDelete.value) return

  try {
    await productService.deleteProduct(productToDelete.value)
    await fetchProducts()
    showSuccess('Product deleted successfully')
    showDeleteModal.value = false
    productToDelete.value = null
  } catch (err) {
    showError(err, 'Failed to delete product')
  }
}

const saveProduct = async () => {
  if (!form.value.categoryId) {
    showError(new Error('Please select a category'))
    return
  }
  if (!form.value.name || form.value.price <= 0) {
    showError(new Error('Product name and valid price are required'))
    return
  }
  try {
    const payload = camelToSnakeCase({
      name: form.value.name,
      price: form.value.price,
      categoryId: form.value.categoryId,
      description: form.value.description || null
    })

    if (editingProduct.value) {
      await productService.updateProduct(editingProduct.value.productId, payload)
      showSuccess('Product updated successfully')
    } else {
      await productService.createProduct(payload)
      showSuccess('Product created successfully')
    }
    showModal.value = false
    editingProduct.value = null
    form.value = { productId: '', name: '', price: 0, categoryId: '', description: '' }
    await fetchProducts()
    } catch (err) {
    const backendErrors = err.response?.data?.errors
    if (backendErrors) {
      const messages = Object.values(backendErrors).flat()
      showError(new Error(messages.join(', ')))
    } else {
      showError(err, 'Failed to save product')
    }
  }
}

const getCategoryName = (categoryId) => {
  const category = categories.value.find(c => c.categoryId === categoryId)
  return category?.name || 'Unknown'
}

onMounted(() => {
  fetchCategories()
  fetchProducts()
})
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">Manage Products</h1>
      <button @click="showModal = true; editingProduct.value = null; form.value = { productId: '', name: '', price: 0, categoryId: '', description: '' }" class="btn-primary">
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
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Category</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Price</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Stock</th>
            <th class="text-right p-4 text-sm font-medium text-gray-700 dark:text-gray-300">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in products" :key="product.productId" class="border-b border-gray-100 dark:border-brand-700 hover:bg-gray-50 dark:hover:bg-brand-700/50 transition-colors">
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">{{ product.name }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ getCategoryName(product.categoryId) }}</td>
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">₹{{ product.price?.toFixed(2) }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ product.stockQuantity ?? 0 }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editProduct(product)" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100">Edit</button>
              <button @click="confirmDelete(product.productId)" class="text-sm text-red-600 hover:text-red-700" :aria-label="`Delete ${product.name}`">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Add/Edit Product Modal -->
    <BaseModal
      :show="showModal"
      :title="editingProduct ? 'Edit Product' : 'Add Product'"
      @close="showModal = false"
      @confirm="saveProduct"
    >
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
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Category</label>
          <select v-model="form.categoryId" class="input" required>
            <option value="">Select Category</option>
            <option v-for="category in categories" :key="category.categoryId" :value="category.categoryId">
              {{ category.name }}
            </option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Description</label>
          <textarea v-model="form.description" placeholder="Description" rows="3" class="input" />
        </div>

      </div>

      <template #footer>
        <button type="button" @click="showModal = false" class="btn-secondary">
          Cancel
        </button>
        <button type="button" @click="saveProduct" class="btn-primary">
          {{ editingProduct ? 'Update' : 'Create' }}
        </button>
      </template>
    </BaseModal>

    <!-- Delete Confirmation Modal -->
    <BaseModal
      :show="showDeleteModal"
      title="Confirm Delete"
      @close="showDeleteModal = false"
      @confirm="deleteProduct"
    >
      <p class="text-sm text-gray-600 dark:text-gray-400">
        Are you sure you want to delete this product? This action cannot be undone.
      </p>

      <template #footer>
        <button type="button" @click="showDeleteModal = false" class="btn-secondary">
          Cancel
        </button>
        <button type="button" @click="deleteProduct" class="btn-danger">
          Delete
        </button>
      </template>
    </BaseModal>
  </div>
</template>
