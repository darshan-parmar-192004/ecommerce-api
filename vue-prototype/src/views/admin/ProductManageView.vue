<script setup>
import { ref, onMounted } from 'vue'
import productService from '@/services/productService'
import categoryService from '@/services/categoryService'
import { useToastStore } from '@/stores/toast'

const toastStore = useToastStore()

const products = ref([])
const categories = ref([])
const loading = ref(true)
const showModal = ref(false)
const editingProduct = ref(null)

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

const deleteProduct = async (id) => {
  if (confirm('Are you sure you want to delete this product?')) {
    try {
      await productService.deleteProduct(id)
      await fetchProducts()
      toastStore.success('Product deleted successfully')
    } catch (err) {
      console.error('Delete failed', err)
      toastStore.error('Failed to delete product')
    }
  }
}

const saveProduct = async () => {
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
        product_id: `PRD-${Date.now()}` // Generate product ID
      })
      toastStore.success('Product created successfully')
    }
    showModal.value = false
    editingProduct.value = null
    form.value = { product_id: '', name: '', price: 0, category_id: '', description: '', stock: 0 }
    await fetchProducts()
  } catch (err) {
    console.error('Save failed', err)
    toastStore.error('Failed to save product')
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
      <h1 class="text-3xl font-bold text-gray-900">Manage Products</h1>
      <button @click="showModal = true; editingProduct.value = null; form.value = { product_id: '', name: '', price: 0, category_id: '', description: '', stock: 0 }" class="btn-primary">
        Add Product
      </button>
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-gray-100 rounded-lg animate-pulse" />
    </div>

    <div v-else class="bg-white rounded-xl border border-gray-200 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Name</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Category</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Price</th>
            <th class="text-left p-4 text-sm font-medium text-gray-700">Stock</th>
            <th class="text-right p-4 text-sm font-medium text-gray-700">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in products" :key="product.product_id" class="border-b border-gray-100 hover:bg-gray-50 transition-colors">
            <td class="p-4 text-sm text-gray-900">{{ product.name }}</td>
            <td class="p-4 text-sm text-gray-600">{{ getCategoryName(product.category_id) }}</td>
            <td class="p-4 text-sm text-gray-900">₹{{ product.price?.toFixed(2) }}</td>
            <td class="p-4 text-sm text-gray-600">{{ product.stock ?? 0 }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editProduct(product)" class="text-sm text-gray-600 hover:text-gray-900">Edit</button>
              <button @click="deleteProduct(product.product_id)" class="text-sm text-red-600 hover:text-red-700">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <Transition name="modal">
      <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50" @click="showModal = false" />
        <div class="relative bg-white rounded-xl p-6 w-full max-w-md">
          <h2 class="text-xl font-bold text-gray-900 mb-4">
            {{ editingProduct ? 'Edit Product' : 'Add Product' }}
          </h2>
          <form @submit.prevent="saveProduct" class="space-y-4">
            <input v-model="form.name" placeholder="Product Name" class="input" required />
            <input v-model.number="form.price" type="number" step="0.01" placeholder="Price" class="input" required />
            <select v-model="form.category_id" class="input" required>
              <option value="">Select Category</option>
              <option v-for="category in categories" :key="category.category_id" :value="category.category_id">
                {{ category.name }}
              </option>
            </select>
            <textarea v-model="form.description" placeholder="Description" rows="3" class="input" />
            <input v-model.number="form.stock" type="number" placeholder="Stock" class="input" required />
            <div class="flex gap-3">
              <button type="button" @click="showModal = false" class="btn-secondary flex-1">Cancel</button>
              <button type="submit" class="btn-primary flex-1">Save</button>
            </div>
          </form>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-from .relative,
.modal-leave-to .relative {
  transform: scale(0.95);
}
.modal-enter-active .relative,
.modal-leave-active .relative {
  transition: transform 0.3s ease;
}
</style>