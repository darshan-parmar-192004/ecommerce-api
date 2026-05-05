<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/services/productService'
import categoryService from '@/services/categoryService'
import { useToastStore } from '@/stores/toast'
import BaseModal from '@/components/ui/BaseModal.vue'

const toastStore = useToastStore()

const products = ref([])
const categories = ref([])
const loading = ref(true)
const showModal = ref(false)
const showDeleteModal = ref(false)
const editingProduct = ref(null)
const productToDelete = ref(null)

const form = ref({
  product_id: '',
  name: '',
  price: 0,
  category_id: '',
  description: '',
  stock: 0
})

const fetchProducts = async () => {
  loading.value = true
  try {
    const response = await productService.getProducts({ limit: 50 })
    const responseData = response.data || response
    products.value = responseData.data || []
  } catch (err) {
    console.error('Failed to fetch products', err)
    toastStore.error('Failed to fetch products')
  } finally {
    loading.value = false
  }
}

const fetchCategories = async () => {
  try {
    const response = await categoryService.getCategories()
    const responseData = response.data || response
    categories.value = responseData.data || []
  } catch (err) {
    console.error('Failed to fetch categories', err)
  }
}

const editProduct = (product) => {
  editingProduct.value = product
  form.value = {
    product_id: product.product_id,
    name: product.name,
    price: product.price,
    category_id: product.category_id,
    description: product.description || '',
    stock: product.stock || 0
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
    toastStore.success('Product deleted successfully')
    showDeleteModal.value = false
    productToDelete.value = null
  } catch (err) {
    console.error('Delete failed', err)
    toastStore.error('Failed to delete product')
  }
}

const saveProduct = async () => {
  if (!form.value.category_id) {
    toastStore.error('Please select a category')
    return
  }
  if (!form.value.name || form.value.price <= 0) {
    toastStore.error('Product name and valid price are required')
    return
  }
  try {
    const payload = {
      name: form.value.name,
      price: form.value.price,
      category_id: form.value.category_id,
      description: form.value.description || null,
      stock: form.value.stock || 0
    }

    if (editingProduct.value) {
      await productService.updateProduct(editingProduct.value.product_id, payload)
      toastStore.success('Product updated successfully')
    } else {
      await productService.createProduct({
        ...payload,
        product_id: `PRD-${Date.now()}`
      })
      toastStore.success('Product created successfully')
    }
    showModal.value = false
    editingProduct.value = null
    form.value = { product_id: '', name: '', price: 0, category_id: '', description: '', stock: 0 }
    await fetchProducts()
  } catch (err) {
    const backendErrors = err.response?.data?.errors
    if (backendErrors) {
      const messages = Object.values(backendErrors).flat()
      toastStore.error(messages.join(', '))
    } else {
      toastStore.error('Failed to save product')
    }
    console.error('Save failed', err)
  }
}

const getCategoryName = (categoryId) => {
  const category = categories.value.find(c => c.category_id === categoryId)
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
      <button @click="showModal = true; editingProduct.value = null; form.value = { product_id: '', name: '', price: 0, category_id: '', description: '', stock: 0 }" class="btn-primary">
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
          <tr v-for="product in products" :key="product.product_id" class="border-b border-gray-100 dark:border-brand-700 hover:bg-gray-50 dark:hover:bg-brand-700/50 transition-colors">
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">{{ product.name }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ getCategoryName(product.category_id) }}</td>
            <td class="p-4 text-sm text-gray-900 dark:text-gray-100">₹{{ product.price?.toFixed(2) }}</td>
            <td class="p-4 text-sm text-gray-600 dark:text-gray-400">{{ product.stock ?? 0 }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editProduct(product)" class="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100">Edit</button>
              <button @click="confirmDelete(product.product_id)" class="text-sm text-red-600 hover:text-red-700" :aria-label="`Delete ${product.name}`">
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
          <select v-model="form.category_id" class="input" required>
            <option value="">Select Category</option>
            <option v-for="category in categories" :key="category.category_id" :value="category.category_id">
              {{ category.name }}
            </option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Description</label>
          <textarea v-model="form.description" placeholder="Description" rows="3" class="input" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Stock</label>
          <input v-model.number="form.stock" type="number" placeholder="Stock" class="input" required />
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
