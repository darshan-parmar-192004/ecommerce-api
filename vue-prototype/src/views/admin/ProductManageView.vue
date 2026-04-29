<script setup>
import { ref, onMounted } from 'vue'
import { useMotion } from '@vueuse/motion'
import productService from '@/services/productService'

const products = ref([])
const loading = ref(true)
const showModal = ref(false)
const editingProduct = ref(null)

const form = ref({
  name: '',
  price: 0,
  category: '',
  description: '',
  stock: 0
})

const fetchProducts = async () => {
  loading.value = true
  try {
    const { data } = await productService.getProducts({ limit: 50 })
    products.value = data.products
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

const deleteProduct = async (id) => {
  if (confirm('Are you sure you want to delete this product?')) {
    try {
      await productService.deleteProduct(id)
      await fetchProducts()
    } catch (err) {
      console.error('Delete failed', err)
    }
  }
}

const saveProduct = async () => {
  try {
    if (editingProduct.value) {
      await productService.updateProduct(editingProduct.value.id, form.value)
    } else {
      await productService.createProduct(form.value)
    }
    showModal.value = false
    editingProduct.value = null
    form.value = { name: '', price: 0, category: '', description: '', stock: 0 }
    await fetchProducts()
  } catch (err) {
    console.error('Save failed', err)
  }
}

onMounted(fetchProducts)
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Manage Products</h1>
      <button @click="showModal = true; editingProduct.value = null; form.value = { name: '', price: 0, category: '', description: '', stock: 0 }" class="btn-primary">
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
          <tr v-for="product in products" :key="product.id" class="border-b border-gray-100 hover:bg-gray-50 transition-colors">
            <td class="p-4 text-sm text-gray-900">{{ product.name }}</td>
            <td class="p-4 text-sm text-gray-600">{{ product.category }}</td>
            <td class="p-4 text-sm text-gray-900">${{ product.price?.toFixed(2) }}</td>
            <td class="p-4 text-sm text-gray-600">{{ product.stock }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editProduct(product)" class="text-sm text-gray-600 hover:text-gray-900">Edit</button>
              <button @click="deleteProduct(product.id)" class="text-sm text-red-600 hover:text-red-700">Delete</button>
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
            <input v-model="form.category" placeholder="Category" class="input" required />
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
